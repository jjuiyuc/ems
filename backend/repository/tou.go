package repository

import (
	"database/sql"
	deremsmodels "der-ems/models/der-ems"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// TOURepository godoc
type TOURepository interface {
	GetTOULocationByTOULocationID(touLocationID int) (*deremsmodels.TouLocation, error)
	GetBillingsByTOUInfo(touLocationID int, voltageType string, touType string, periodType string, isSummer bool, day string) ([]*deremsmodels.Tou, error)
	GetHolidayByDay(touLocationID int, year string, day string) (int64, error)
}

type defaultTOURepository struct {
	db *sql.DB
}

// NewTOURepository godoc
func NewTOURepository(db *sql.DB) TOURepository {
	return &defaultTOURepository{db}
}

// GetTOULocationByTOULocationID godoc
func (repo defaultTOURepository) GetTOULocationByTOULocationID(touLocationID int) (*deremsmodels.TouLocation, error) {
	return deremsmodels.FindTouLocation(repo.db, touLocationID)
}

// GetBillingsByTOUInfo godoc
func (repo defaultTOURepository) GetBillingsByTOUInfo(touLocationID int, voltageType string, touType string, periodType string, isSummer bool, day string) ([]*deremsmodels.Tou, error) {
	return deremsmodels.Tous(
		qm.Where("enable_at <= ?", day),
		qm.Where("disable_at >= ?", day),
		qm.Where("is_summer = ?", isSummer),
		qm.Where("period_type = ?", periodType),
		qm.Where("tou_type = ?", touType),
		qm.Where("voltage_type = ?", voltageType),
		qm.Where("tou_location_id = ?", touLocationID)).All(repo.db)
}

// GetHolidayByDay godoc
func (repo defaultTOURepository) GetHolidayByDay(touLocationID int, year string, day string) (int64, error) {
	return deremsmodels.Tous(
		qm.Where("day = ?", day),
		qm.Where("year = ?", year),
		qm.Where("tou_location_id = ?", touLocationID)).Count(repo.db)
}
