package repository

import (
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
)

func UpsertWeatherForecast(weatherForecast *deremsmodels.WeatherForecast) (err error) {
	mods := []qm.QueryMod{
		qm.Where("lat = ?", weatherForecast.Lat),
		qm.Where("lng = ?", weatherForecast.Lng),
		qm.Where("valid_date = ?", weatherForecast.ValidDate),
	}

	var weatherForecastReturn *deremsmodels.WeatherForecast
	weatherForecastReturn, err = deremsmodels.WeatherForecasts(mods...).One(models.GetDB())
	if err != nil {
		err = weatherForecast.Insert(models.GetDB(), boil.Infer())
	} else {
		weatherForecast.ID = weatherForecastReturn.ID
		weatherForecast.UpdatedAt = null.NewTime(time.Now(), true)
		_, err = weatherForecast.Update(models.GetDB(), boil.Infer())
	}
	return
}

func GetWeatherForecastByLocation(lat, lng float32, startValidDate, endValidDate time.Time) (weatherForecast []*deremsmodels.WeatherForecast, err error) {
	weatherForecast, err = deremsmodels.WeatherForecasts(
		qm.Where("lat = ?", lat),
		qm.Where("lng = ?", lng),
		qm.Where("(valid_date > ? and valid_date <= ?)", startValidDate, endValidDate)).All(models.GetDB())
	return
}

func GetGatewaysByLocation(lat, lng float32) (gateways []*deremsmodels.Gateway, err error) {
	gateways, err = deremsmodels.Gateways(
		qm.InnerJoin("customer AS c ON gateway.customer_id = c.id"),
		qm.Where("(c.weather_lat = ? AND c.weather_lng = ?)", lat, lng)).All(models.GetDB())
	return
}
