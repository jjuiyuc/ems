package repository

import (
	"time"

	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"der-ems/infra"
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

func GetWeatherForecastByLocation(lat, lng float32, startValidDate, endValidDate time.Time) (weatherForecast []*deremsmodels.WeatherForecast, err error) {
	weatherForecast, err = deremsmodels.WeatherForecasts(
		qm.Where("lat = ?", lat),
		qm.Where("lng = ?", lng),
		qm.Where("(valid_date > ? and valid_date <= ?)", startValidDate, endValidDate)).All(models.GetDB())
	return
}

func GetGatewaysByLocation(lat, lng float32) (gatewayUUIDs []string, err error) {
	sql := `
		SELECT gw.uuid
		FROM gateway gw
		JOIN customer_info c ON c.weather_lat = ? AND c.weather_lng = ?
		WHERE gw.customer_id = c.id
	`
	var gateways []*deremsmodels.Gateway
	err = queries.Raw(sql, lat, lng).Bind(infra.GetGracefulShutdownCtx(), models.GetDB(), &gateways)
	for _, gateway := range gateways {
		gatewayUUIDs = append(gatewayUUIDs, gateway.UUID)
	}
	return
}
