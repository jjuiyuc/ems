package repository

import (
	"database/sql"
	deremsmodels "der-ems/models/der-ems"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// TOURepository godoc
type TOURepository interface {
	GetTOULocationByTOULocationID(touLocationID int64) (*deremsmodels.TouLocation, error)
	GetTOUsByTOUInfo(touLocationID int64, voltageType, touType, periodType string, isSummer bool, day string) ([]*deremsmodels.Tou, error)
	CountHolidayByDay(touLocationID int64, year, day string) (int64, error)
	GetTOULocationByPowerCompany(tx *sql.Tx, powerCompany string) (*deremsmodels.TouLocation, error)
}

type defaultTOURepository struct {
	db *sql.DB
}

// NewTOURepository godoc
func NewTOURepository(db *sql.DB) TOURepository {
	return &defaultTOURepository{db}
}

// GetTOULocationByTOULocationID godoc
func (repo defaultTOURepository) GetTOULocationByTOULocationID(touLocationID int64) (*deremsmodels.TouLocation, error) {
	return deremsmodels.FindTouLocation(repo.db, touLocationID)
}

// GetTOUsByTOUInfo godoc
func (repo defaultTOURepository) GetTOUsByTOUInfo(touLocationID int64, voltageType, touType, periodType string, isSummer bool, day string) ([]*deremsmodels.Tou, error) {
	return deremsmodels.Tous(
		qm.Where("enable_at <= ?", day),
		qm.Where("disable_at >= ?", day),
		qm.Where("is_summer = ?", isSummer),
		qm.Where("period_type = ?", periodType),
		qm.Where("tou_type = ?", touType),
		qm.Where("voltage_type = ?", voltageType),
		qm.Where("tou_location_id = ?", touLocationID)).All(repo.db)
}

// CountHolidayByDay godoc
func (repo defaultTOURepository) CountHolidayByDay(touLocationID int64, year, day string) (int64, error) {
	return deremsmodels.TouHolidays(
		qm.Where("day = ?", day),
		qm.Where("year = ?", year),
		qm.Where("tou_location_id = ?", touLocationID)).Count(repo.db)
}

func (repo defaultTOURepository) GetTOULocationByPowerCompany(tx *sql.Tx, powerCompany string) (*deremsmodels.TouLocation, error) {
	return deremsmodels.TouLocations(
		qm.Where("power_company = ?", powerCompany)).One(repo.getExecutor(tx))
}

func (repo defaultTOURepository) getExecutor(tx *sql.Tx) boil.Executor {
	if tx == nil {
		return repo.db
	}
	return tx
}
