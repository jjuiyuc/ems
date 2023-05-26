package apps

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
	"der-ems/services"
)

// WeatherWorker godoc
type WeatherWorker struct {
	kafka.SimpleConsumer
}

type weatherConsumerHandler struct {
	cfg  *viper.Viper
	repo *repository.Repository
	weather services.WeatherService
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
	weather services.WeatherService,
	name string,
) (w *WeatherWorker) {
	topics := []string{
		kafka.ReceiveWeatherData,
	}
	handler := weatherConsumerHandler{
		cfg:  cfg,
		repo: repo,
		weather: weather,
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
	utils.PrintFunctionName()
	lat, lng, err := h.saveWeatherData(msg)
	if err != nil {
		return
	}
	data, gatewayUUIDs, err := h.weather.GenerateWeatherInfo(lat, lng)
	if err != nil {
		return
	}
	kafka.SendDataToGateways(h.cfg, kafka.SendWeatherDataToLocalGW, data, gatewayUUIDs)
}

func (h weatherConsumerHandler) saveWeatherData(msg []byte) (lat, lng float32, err error) {
	var latestWeather services.LatestWeather
	if err = json.Unmarshal(msg, &latestWeather); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}

	for i, data := range latestWeather.Values {
		const validDate = "validDate"

		dt, ok := data[validDate]
		if !ok {
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
			Alt:       null.Float32From(latestWeather.Alt),
			ValidDate: dt.(time.Time),
			Data:      null.JSONFrom(dataJSON),
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
