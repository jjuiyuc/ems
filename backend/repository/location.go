package repository

import (
	"database/sql"

	deremsmodels "der-ems/models/der-ems"
)

// LocationRepository godoc
type LocationRepository interface {
	GetLocationByLocationID(locationID int64) (*deremsmodels.Location, error)
}

type defaultLocationRepository struct {
	db *sql.DB
}

// NewLocationRepository godoc
func NewLocationRepository(db *sql.DB) LocationRepository {
	return &defaultLocationRepository{db}
}

// GetLocationByLocationID godoc
func (repo defaultLocationRepository) GetLocationByLocationID(locationID int64) (*deremsmodels.Location, error) {
	return deremsmodels.FindLocation(repo.db, locationID)
}
