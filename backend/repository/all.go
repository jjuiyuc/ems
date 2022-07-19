package repository

import "database/sql"

// Repository godoc
type Repository struct {
	User     UserRepository
	Weather  WeatherRepository
	Gateway  GatewayRepository
	CCData   CCDataRepository
	Customer CustomerRepository
	TOU      TOURepository
}

// NewRepository godoc
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:     NewUserRepository(db),
		Weather:  NewWeatherRepository(db),
		Gateway:  NewGatewayRepository(db),
		CCData:   NewCCDataRepository(db),
		Customer: NewCustomerRepository(db),
		TOU:      NewTOURepository(db),
	}
}
