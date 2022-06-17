package repository

import "database/sql"

// Repository godoc
type Repository struct {
	User    UserRepository
	Weather WeatherRepository
	Gateway GatewayRepository
	CCData  CCDataRepository
}

// NewRepository godoc
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Weather: NewWeatherRepository(db),
		Gateway: NewGatewayRepository(db),
		CCData:  NewCCDataRepository(db),
	}
}
