package repository

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"

	deremsmodels "der-ems/models/der-ems"
)

// LocationRepository godoc
type LocationRepository interface {
	GetLocationByLocationID(locationID int64) (*deremsmodels.Location, error)
	CreateLocation(tx *sql.Tx, location *deremsmodels.Location) error
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

func (repo defaultLocationRepository) CreateLocation(tx *sql.Tx, location *deremsmodels.Location) error {
	return location.Insert(repo.getExecutor(tx), boil.Infer())
}

func (repo defaultLocationRepository) getExecutor(tx *sql.Tx) boil.Executor {
	if tx == nil {
		return repo.db
	}
	return tx
}
