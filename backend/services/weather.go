package services

import (
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
)

func UpsertWeatherForecast(weatherForecast *deremsmodels.WeatherForecast) (err error) {
	mods := make([]qm.QueryMod, 0)
	mods = append(mods, qm.Where("lat = ?", weatherForecast.Lat))
	mods = append(mods, qm.Where("lng = ?", weatherForecast.LNG))
	mods = append(mods, qm.Where("valid_date = ?", weatherForecast.ValidDate))

	var weatherForecastReturn *deremsmodels.WeatherForecast
	weatherForecastReturn, err = deremsmodels.WeatherForecasts(mods...).One(models.GetDeremsDB())
	if err != nil {
		err = weatherForecast.Insert(models.GetDeremsDB(), boil.Infer())
	} else {
		weatherForecast.ID = weatherForecastReturn.ID
		weatherForecast.UpdatedAt = time.Now()
		_, err = weatherForecast.Update(models.GetDeremsDB(), boil.Infer())
	}
	return
}
