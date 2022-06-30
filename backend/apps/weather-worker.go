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

// WeatherWorker godoc
type WeatherWorker struct {
	kafka.SimpleConsumer
}

// LatestWeather godoc
type LatestWeather struct {
	Lat    float32                  `json:"lat"`
	Lng    float32                  `json:"lng"`
	Alt    float32                  `json:"alt"`
	Values []map[string]interface{} `json:"values"`
}

type weatherConsumerHandler struct {
	cfg  *viper.Viper
	repo *repository.Repository
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
			h.processWeatherData(msg.Value)
		}

		// Mark a message as consumed
		sess.MarkMessage(msg, "")
	}
	return nil
}

// NewWeatherWorker godoc
func NewWeatherWorker(
	ctx context.Context,
	cfg *viper.Viper,
	repo *repository.Repository,
	name string,
) (w *WeatherWorker) {
	topics := []string{
		kafka.ReceiveWeatherData,
	}
	handler := weatherConsumerHandler{
		cfg:  cfg,
		repo: repo,
	}

	simpleConsumer, err := kafka.NewSimpleConsumer(ctx, cfg, name, handler, topics)
	if err != nil {
		return
	}

	w = &WeatherWorker{
		SimpleConsumer: *simpleConsumer,
	}

	return
}

// MainLoop godoc
func (w *WeatherWorker) MainLoop() {
	w.SimpleConsumer.MainLoop()
}

func (h weatherConsumerHandler) processWeatherData(msg []byte) {
	log.Debug("processWeatherData")
	lat, lng, err := h.saveWeatherData(msg)
	if err != nil {
		return
	}
	latestWeatherJSON, gatewayUUIDs, err := h.generateWeatherSendingInfo(lat, lng)
	if err != nil {
		return
	}
	h.sendWeatherDataToGateway(latestWeatherJSON, gatewayUUIDs)
}

func (h weatherConsumerHandler) saveWeatherData(msg []byte) (lat, lng float32, err error) {
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
			err = e.ErrNewKeyNotExist(validDate)
			log.WithFields(log.Fields{
				"caused-by": validDate,
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
		dataJSON, _ := json.Marshal(data)

		weatherForecast := &deremsmodels.WeatherForecast{
			Lat:       latestWeather.Lat,
			Lng:       latestWeather.Lng,
			Alt:       null.NewFloat32(latestWeather.Alt, true),
			ValidDate: dt.(time.Time),
			Data:      null.NewJSON(dataJSON, true),
		}

		log.WithFields(log.Fields{
			"i":                                           i,
			deremsmodels.WeatherForecastColumns.Lat:       weatherForecast.Lat,
			deremsmodels.WeatherForecastColumns.Lng:       weatherForecast.Lng,
			deremsmodels.WeatherForecastColumns.Alt:       weatherForecast.Alt,
			deremsmodels.WeatherForecastColumns.ValidDate: weatherForecast.ValidDate,
			deremsmodels.WeatherForecastColumns.Data:      string(weatherForecast.Data.JSON),
		}).Debug("upsert weatherForecast data")
		err = h.repo.Weather.UpsertWeatherForecast(weatherForecast)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "h.repo.Weather.UpsertWeatherForecast",
				"err":       err,
			}).Error()
			return
		}
	}

	lat = latestWeather.Lat
	lng = latestWeather.Lng
	return
}

func (h weatherConsumerHandler) generateWeatherSendingInfo(lat, lng float32) (latestWeatherJSON []byte, gatewayUUIDs []string, err error) {
	latestWeatherJSON, err = h.getWeatherDataByLocation(lat, lng)
	if err != nil {
		return
	}
	gatewayUUIDs, err = h.getGatewayUUIDsByLocation(lat, lng)
	return
}

func (h weatherConsumerHandler) getWeatherDataByLocation(lat, lng float32) (latestWeatherJSON []byte, err error) {
	now := time.Now().UTC()
	startValidDate := now
	endValidDate := now.Add(30 * time.Hour)
	weatherForecastList, err := h.repo.Weather.GetWeatherForecastByLocation(lat, lng, startValidDate, endValidDate)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.Weather.GetWeatherForecastByLocation",
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
	latestWeatherJSON, err = json.Marshal(latestWeather)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
	}
	return
}

func (h weatherConsumerHandler) getGatewayUUIDsByLocation(lat, lng float32) (gatewayUUIDs []string, err error) {
	gateways, err := h.repo.Gateway.GetGatewaysByLocation(lat, lng)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.Gateway.GetGatewaysByLocation",
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

func (h weatherConsumerHandler) sendWeatherDataToGateway(latestWeatherJSON []byte, gatewayUUIDs []string) {
	for _, uuid := range gatewayUUIDs {
		sendWeatherDataToLocalGW := strings.Replace(kafka.SendWeatherDataToLocalGW, "{gw-id}", uuid, 1)
		log.Debug("sendWeatherDataToLocalGW: ", sendWeatherDataToLocalGW)
		kafka.Produce(h.cfg, sendWeatherDataToLocalGW, string(latestWeatherJSON))
	}
}
