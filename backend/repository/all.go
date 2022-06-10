package repository

import "database/sql"

// Repository ...
type Repository struct {
	User    UserRepository
	Weather WeatherRepository
}

// NewRepository ...
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Weather: NewWeatherRepository(db),
	}
}
