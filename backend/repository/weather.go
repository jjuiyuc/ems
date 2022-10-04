package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// WeatherRepository godoc
type WeatherRepository interface {
	UpsertWeatherForecast(weatherForecast *deremsmodels.WeatherForecast) (err error)
	GetWeatherForecastByLocation(lat, lng float32, startValidDate, endValidDate time.Time) ([]*deremsmodels.WeatherForecast, error)
	GetWeatherForecastCount() (int64, error)
}

type defaultWeatherRepository struct {
	db *sql.DB
}

// NewWeatherRepository godoc
func NewWeatherRepository(db *sql.DB) WeatherRepository {
	return &defaultWeatherRepository{db}
}

// UpsertWeatherForecast godoc
func (repo defaultWeatherRepository) UpsertWeatherForecast(weatherForecast *deremsmodels.WeatherForecast) (err error) {
	var weatherForecastReturn *deremsmodels.WeatherForecast
	weatherForecastReturn, err = deremsmodels.WeatherForecasts(
		qm.Where("lat = ?", weatherForecast.Lat),
		qm.Where("lng = ?", weatherForecast.Lng),
		qm.Where("valid_date = ?", weatherForecast.ValidDate)).One(repo.db)
	now := time.Now().UTC()
	weatherForecast.UpdatedAt = now
	if err != nil {
		weatherForecast.CreatedAt = now
		err = weatherForecast.Insert(repo.db, boil.Infer())
	} else {
		weatherForecast.ID = weatherForecastReturn.ID
		_, err = weatherForecast.Update(repo.db, boil.Infer())
	}
	return
}

// GetWeatherForecastByLocation godoc
func (repo defaultWeatherRepository) GetWeatherForecastByLocation(lat, lng float32, startValidDate, endValidDate time.Time) ([]*deremsmodels.WeatherForecast, error) {
	return deremsmodels.WeatherForecasts(
		qm.Where("lat = ?", lat),
		qm.Where("lng = ?", lng),
		qm.Where("(valid_date > ? and valid_date <= ?)", startValidDate, endValidDate)).All(repo.db)
}

// GetWeatherForecastCount godoc
func (repo defaultWeatherRepository) GetWeatherForecastCount() (int64, error) {
	return deremsmodels.WeatherForecasts().Count(repo.db)
}
