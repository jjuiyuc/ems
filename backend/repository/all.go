package repository

import "database/sql"

// Repository ...
type Repository struct {
	User    UserRepository
	Weather WeatherRepository
	Gateway GatewayRepository
	CCData  CCDataRepository
}

// NewRepository ...
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Weather: NewWeatherRepository(db),
		Gateway: NewGatewayRepository(db),
		CCData:  NewCCDataRepository(db),
	}
}
