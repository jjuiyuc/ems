package repository

import (
	"time"

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
		weatherForecast.UpdatedAt = time.Now()
		_, err = weatherForecast.Update(models.GetDB(), boil.Infer())
	}
	return
}

func GetWeatherForecastByLocation(lat, lng float32) (weatherForecastList []*deremsmodels.WeatherForecast, err error) {
	mods := []qm.QueryMod{
		qm.Where("lat = ?", lat),
		qm.Where("lng = ?", lng),
		qm.Where("(valid_date > ? and valid_date <= ?)", time.Now().UTC(), time.Now().UTC().Add(30*time.Hour)),
	}

	weatherForecastList, err = deremsmodels.WeatherForecasts(mods...).All(models.GetDB())
	if err != nil {
		return
	}

	return
}

func GetCustomerInfoListByLocation(lat, lng float32) (customerInfos []*deremsmodels.CustomerInfo, err error) {
	mods := []qm.QueryMod{
		qm.Where("weather_lat = ?", lat),
		qm.Where("weather_lng = ?", lng),
	}

	customerInfos, err = deremsmodels.CustomerInfos(mods...).All(models.GetDB())
	if err != nil {
		return
	}

	return
}

func GetGatewayListByCustomerID(id int) (gateways []*deremsmodels.Gateway, err error) {
	mods := []qm.QueryMod{
		qm.Where("customer_id = ?", id),
	}

	gateways, err = deremsmodels.Gateways(mods...).All(models.GetDB())
	if err != nil {
		return
	}

	return
}
