package apps

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/e"
	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

type WeatherWorker struct {
	kafka.SimpleConsumer
}

type LatestWeather struct {
	Lat    float32                  `json:"lat"`
	Lng    float32                  `json:"lng"`
	Alt    float32                  `json:"alt"`
	Values []map[string]interface{} `json:"values"`
}

type weatherConsumerHandler struct {
}

func (weatherConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (weatherConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h weatherConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.WithFields(log.Fields{
			"topic":     msg.Topic,
			"partition": msg.Partition,
			"offset":    msg.Offset,
			"value":     string(msg.Value),
		}).Info("consuming")

		if msg.Topic == kafka.ReceiveWeatherData {
			ProcessWeatherData(msg.Value)
		}

		// Mark a message as consumed
		sess.MarkMessage(msg, "")
	}
	return nil
}

func NewWeatherWorker(
	ctx context.Context,
	cfg *viper.Viper,
	name string,
) (w *WeatherWorker) {
	topics := []string{
		kafka.ReceiveWeatherData,
	}
	handler := weatherConsumerHandler{}

	simpleConsumer, err := kafka.NewSimpleConsumer(ctx, cfg, name, handler, topics)
	if err != nil {
		return
	}

	w = &WeatherWorker{
		SimpleConsumer: *simpleConsumer,
	}

	return
}

func (w *WeatherWorker) MainLoop() {
	w.SimpleConsumer.MainLoop()
}

func ProcessWeatherData(msg []byte) {
	log.Debug("processWeatherData")
	lat, lng, err := UpsertWeatherData(msg)
	if err == nil {
		publishWeatherDataToLocalGW(lat, lng)
	}
}

func UpsertWeatherData(msg []byte) (lat, lng float32, err error) {
	var latestWeather LatestWeather
	err = json.Unmarshal(msg, &latestWeather)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}

	for i, data := range latestWeather.Values {
		const validDate = "validDate"

		dt := data[validDate]
		if dt == nil {
			err = e.ErrUlNoValidDate
			log.WithFields(log.Fields{
				"caused-by": "data[validDate]",
				"err":       err,
			}).Error()
			return
		}
		if dt, err = time.Parse(time.RFC3339, fmt.Sprintf("%v", dt)); err != nil {
			log.WithFields(log.Fields{
				"caused-by": "time.Parse",
				"err":       err,
			}).Error()
			return
		}

		delete(data, validDate)
		dataJson, _ := json.Marshal(data)

		weatherForecast := &deremsmodels.WeatherForecast{
			Lat:       latestWeather.Lat,
			Lng:       latestWeather.Lng,
			Alt:       null.NewFloat32(latestWeather.Alt, true),
			ValidDate: dt.(time.Time),
			Data:      null.NewJSON(dataJson, true),
		}

		log.WithFields(log.Fields{
			"i":            i,
			"Lat":          weatherForecast.Lat,
			"Lng":          weatherForecast.Lng,
			"Alt":          weatherForecast.Alt,
			"ValidDate":    weatherForecast.ValidDate,
			"string(Data)": string(weatherForecast.Data.JSON),
		}).Debug("upsert weatherForecast data")
		err = repository.UpsertWeatherForecast(weatherForecast)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "repository.UpsertWeatherForecast",
				"err":       err,
			}).Error()
			return
		}
	}

	lat = latestWeather.Lat
	lng = latestWeather.Lng
	return
}

func publishWeatherDataToLocalGW(lat, lng float32) {
	latestWeatherJson, err := GetWeatherDataByLocation(lat, lng)
	if err != nil {
		return
	}
	gatewayUUIDs, err := GetGateWayUUIDsByLocation(lat, lng)
	if err != nil {
		return
	}
	publish(latestWeatherJson, gatewayUUIDs)
}

func GetWeatherDataByLocation(lat, lng float32) (latestWeatherJson []byte, err error) {
	startValidDate := time.Now().UTC()
	endValidDate := time.Now().UTC().Add(30 * time.Hour)
	weatherForecastList, err := repository.GetWeatherForecastByLocation(lat, lng, startValidDate, endValidDate)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "repository.GetWeatherForecastByLocation",
			"err":       err,
		}).Error()
		return
	}
	var latestWeather LatestWeather
	for i, weatherForecast := range weatherForecastList {
		const validDate = "validDate"
		var value map[string]interface{}
		if i == 0 {
			latestWeather.Lat = weatherForecast.Lat
			latestWeather.Lng = weatherForecast.Lng
			latestWeather.Alt = weatherForecast.Alt.Float32
		}

		err = json.Unmarshal(weatherForecast.Data.JSON, &value)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "json.Unmarshal",
				"err":       err,
			}).Error()
			return
		}
		value[validDate] = weatherForecast.ValidDate.Format(time.RFC3339)
		latestWeather.Values = append(latestWeather.Values, value)
	}
	latestWeatherJson, err = json.Marshal(latestWeather)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
	}
	return
}

func GetGateWayUUIDsByLocation(lat, lng float32) (gatewayUUIDs []string, err error) {
	gateways, err := repository.GetGatewaysByLocation(lat, lng)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "repository.GetGatewaysByLocation",
			"err":       err,
		}).Error()
		return
	}
	for _, gateway := range gateways {
		gatewayUUIDs = append(gatewayUUIDs, gateway.UUID)
	}
	log.Debug("gatewayUUIDs: ", gatewayUUIDs)
	return
}

func publish(latestWeatherJson []byte, gatewayUUIDs []string) {
	for _, uuid := range gatewayUUIDs {
		sendWeatherDatatoLocalGW := strings.Replace(kafka.SendWeatherDatatoLocalGW, "{gw-id}", uuid, 1)
		log.Debug("sendWeatherDatatoLocalGW: ", sendWeatherDatatoLocalGW)
		kafka.Produce(sendWeatherDatatoLocalGW, string(latestWeatherJson))
	}
}
