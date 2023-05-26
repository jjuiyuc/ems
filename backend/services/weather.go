package services

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"

	"der-ems/repository"
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
	GenerateWeatherInfo(lat, lng float32) (data []byte, gatewayUUIDs []string, err error)
}

type defaultWeatherService struct {
	repo *repository.Repository
}

// NewWeatherService godoc
func NewWeatherService(repo *repository.Repository) WeatherService {
	return &defaultWeatherService{repo}
}

func (s defaultWeatherService) GenerateWeatherInfo(lat, lng float32) (data []byte, gatewayUUIDs []string, err error) {
	data, err = s.getWeatherDataByLocation(lat, lng)
	if err != nil {
		return
	}
	gatewayUUIDs, err = s.getGatewayUUIDsByLocation(lat, lng)
	return
}

func (s defaultWeatherService) getWeatherDataByLocation(lat, lng float32) (data []byte, err error) {
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
	data, err = json.Marshal(latestWeather)
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
