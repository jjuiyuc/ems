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
