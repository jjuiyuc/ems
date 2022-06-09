package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// WeatherRepository ...
type WeatherRepository interface {
	UpsertWeatherForecast(weatherForecast *deremsmodels.WeatherForecast) (err error)
	GetWeatherForecastByLocation(lat, lng float32, startValidDate, endValidDate time.Time) ([]*deremsmodels.WeatherForecast, error)
	GetWeatherForecastCount() (int64, error)
	GetGatewaysByLocation(lat, lng float32) ([]*deremsmodels.Gateway, error)
}

type defaultWeatherRepository struct {
	db *sql.DB
}

// NewWeatherRepository ...
func NewWeatherRepository(db *sql.DB) WeatherRepository {
	return &defaultWeatherRepository{db}
}

// UpsertWeatherForecast ...
func (repo defaultWeatherRepository) UpsertWeatherForecast(weatherForecast *deremsmodels.WeatherForecast) (err error) {
	var weatherForecastReturn *deremsmodels.WeatherForecast
	weatherForecastReturn, err = deremsmodels.WeatherForecasts(
		qm.Where("lat = ?", weatherForecast.Lat),
		qm.Where("lng = ?", weatherForecast.Lng),
		qm.Where("valid_date = ?", weatherForecast.ValidDate)).One(repo.db)
	if err != nil {
		err = weatherForecast.Insert(repo.db, boil.Infer())
	} else {
		weatherForecast.ID = weatherForecastReturn.ID
		weatherForecast.UpdatedAt = null.NewTime(time.Now(), true)
		_, err = weatherForecast.Update(repo.db, boil.Infer())
	}
	return
}

// GetWeatherForecastByLocation ...
func (repo defaultWeatherRepository) GetWeatherForecastByLocation(lat, lng float32, startValidDate, endValidDate time.Time) ([]*deremsmodels.WeatherForecast, error) {
	return deremsmodels.WeatherForecasts(
		qm.Where("lat = ?", lat),
		qm.Where("lng = ?", lng),
		qm.Where("(valid_date > ? and valid_date <= ?)", startValidDate, endValidDate)).All(repo.db)
}

// GetWeatherForecastCount ...
func (repo defaultWeatherRepository) GetWeatherForecastCount() (int64, error) {
	return deremsmodels.WeatherForecasts().Count(repo.db)
}

// GetGatewaysByLocation ...
func (repo defaultWeatherRepository) GetGatewaysByLocation(lat, lng float32) ([]*deremsmodels.Gateway, error) {
	return deremsmodels.Gateways(
		qm.InnerJoin("customer AS c ON gateway.customer_id = c.id"),
		qm.Where("(c.weather_lat = ? AND c.weather_lng = ?)", lat, lng)).All(repo.db)
}
