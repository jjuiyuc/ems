package repository

import "database/sql"

// Repository godoc
type Repository struct {
	User     UserRepository
	Weather  WeatherRepository
	Gateway  GatewayRepository
	CCData   CCDataRepository
	Location LocationRepository
	TOU      TOURepository
	AIData   AIDataRepository
}

// NewRepository godoc
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:     NewUserRepository(db),
		Weather:  NewWeatherRepository(db),
		Gateway:  NewGatewayRepository(db),
		CCData:   NewCCDataRepository(db),
		Location: NewLocationRepository(db),
		TOU:      NewTOURepository(db),
		AIData:   NewAIDataRepository(db),
	}
}
