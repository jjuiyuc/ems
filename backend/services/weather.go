package services

import (
	"der-ems/kafka"
	"der-ems/repository"
	"encoding/json"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LatestWeather godoc
type LatestWeather struct {
	Lat    float32                  `json:"lat"`
	Lng    float32                  `json:"lng"`
	Alt    float32                  `json:"alt"`
	Values []map[string]interface{} `json:"values"`
}

// WeatherService godoc
type WeatherService interface {
	GenerateWeatherSendingInfo(lat, lng float32) (latestWeatherJSON []byte, gatewayUUIDs []string, err error)
	SendWeatherDataToGateway(cfg *viper.Viper, latestWeatherJSON []byte, gatewayUUIDs []string)
}

type defaultWeatherService struct {
	repo *repository.Repository
}

// NewWeatherService godoc
func NewWeatherService(repo *repository.Repository) WeatherService {
	return &defaultWeatherService{repo}
}

func (s defaultWeatherService) GenerateWeatherSendingInfo(lat, lng float32) (latestWeatherJSON []byte, gatewayUUIDs []string, err error) {
	latestWeatherJSON, err = s.getWeatherDataByLocation(lat, lng)
	if err != nil {
		return
	}
	gatewayUUIDs, err = s.getGatewayUUIDsByLocation(lat, lng)
	return
}

func (s defaultWeatherService) getWeatherDataByLocation(lat, lng float32) (latestWeatherJSON []byte, err error) {
	now := time.Now().UTC()
	startValidDate := now
	endValidDate := now.Add(31 * time.Hour)
	weatherForecastList, err := s.repo.Weather.GetWeatherForecastByLocation(lat, lng, startValidDate, endValidDate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Weather.GetWeatherForecastByLocation",
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

		if err = json.Unmarshal(weatherForecast.Data.JSON, &value); err != nil {
			logrus.WithFields(logrus.Fields{
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
		logrus.WithFields(logrus.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultWeatherService) getGatewayUUIDsByLocation(lat, lng float32) (gatewayUUIDs []string, err error) {
	gateways, err := s.repo.Gateway.GetGatewaysByLocation(lat, lng)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewaysByLocation",
			"err":       err,
		}).Error()
		return
	}
	for _, gateway := range gateways {
		gatewayUUIDs = append(gatewayUUIDs, gateway.UUID)
	}
	logrus.Debug("gatewayUUIDs: ", gatewayUUIDs)
	return
}

func (s defaultWeatherService) SendWeatherDataToGateway(cfg *viper.Viper, latestWeatherJSON []byte, gatewayUUIDs []string) {
	for _, uuid := range gatewayUUIDs {
		sendWeatherDataToLocalGW := strings.Replace(kafka.SendWeatherDataToLocalGW, "{gw-id}", uuid, 1)
		logrus.Debug("sendWeatherDataToLocalGW: ", sendWeatherDataToLocalGW)
		kafka.Produce(cfg, sendWeatherDataToLocalGW, string(latestWeatherJSON))
	}
}
