package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// CCDataRepository godoc
type CCDataRepository interface {
	// CC data
	UpsertCCData(ccData *deremsmodels.CCDatum) (err error)
	GetCCDataCount() (int64, error)
	// CC data log
	UpsertCCDataLog(ccDataLog *deremsmodels.CCDataLog) (err error)
	GetLatestLogByGatewayUUIDAndPeriod(gwUUID string, startTime, endTime time.Time) (*deremsmodels.CCDataLog, error)
	GetFirstLogByGatewayUUIDAndPeriod(gwUUID string, startTime time.Time, endTime time.Time) (*deremsmodels.CCDataLog, error)
	GetCCDataLogCount() (int64, error)
}

type defaultCCDataRepository struct {
	db *sql.DB
}

// NewCCDataRepository godoc
func NewCCDataRepository(db *sql.DB) CCDataRepository {
	return &defaultCCDataRepository{db}
}

// UpsertCCData godoc
func (repo defaultCCDataRepository) UpsertCCData(ccData *deremsmodels.CCDatum) (err error) {
	var ccDataReturn *deremsmodels.CCDatum
	ccDataReturn, err = deremsmodels.CCData(
		qm.Where("gw_uuid = ?", ccData.GWUUID),
		qm.Where("log_date = ?", ccData.LogDate)).One(repo.db)
	now := time.Now().UTC()
	ccData.UpdatedAt = null.NewTime(now, true)
	if err != nil {
		ccData.CreatedAt = now
		err = ccData.Insert(repo.db, boil.Infer())
	} else {
		ccData.ID = ccDataReturn.ID
		_, err = ccData.Update(repo.db, boil.Infer())
	}
	return
}

// GetCCDataCount godoc
func (repo defaultCCDataRepository) GetCCDataCount() (int64, error) {
	return deremsmodels.CCData().Count(repo.db)
}

// UpsertCCDataLog godoc
func (repo defaultCCDataRepository) UpsertCCDataLog(ccDataLog *deremsmodels.CCDataLog) (err error) {
	var ccDataLogReturn *deremsmodels.CCDataLog
	ccDataLogReturn, err = deremsmodels.CCDataLogs(
		qm.Where("gw_uuid = ?", ccDataLog.GWUUID),
		qm.Where("log_date = ?", ccDataLog.LogDate)).One(repo.db)
	now := time.Now().UTC()
	ccDataLog.UpdatedAt = null.NewTime(now, true)
	if err != nil {
		ccDataLog.CreatedAt = now
		err = ccDataLog.Insert(repo.db, boil.Infer())
	} else {
		ccDataLog.ID = ccDataLogReturn.ID
		_, err = ccDataLog.Update(repo.db, boil.Infer())
	}
	return
}

// GetLatestLogByGatewayUUIDAndPeriod godoc
func (repo defaultCCDataRepository) GetLatestLogByGatewayUUIDAndPeriod(gwUUID string, startTime, endTime time.Time) (*deremsmodels.CCDataLog, error) {
	if startTime.IsZero() || endTime.IsZero() {
		return deremsmodels.CCDataLogs(
			qm.Where("gw_uuid = ?", gwUUID),
			qm.OrderBy("log_date DESC")).One(repo.db)
	}

	return deremsmodels.CCDataLogs(
		qm.Where("(gw_uuid = ? and log_date > ? and log_date <= ?)", gwUUID, startTime, endTime),
		qm.OrderBy("log_date DESC")).One(repo.db)
}

// GetFirstLogByGatewayUUIDAndPeriod godoc
func (repo defaultCCDataRepository) GetFirstLogByGatewayUUIDAndPeriod(gwUUID string, startTime time.Time, endTime time.Time) (*deremsmodels.CCDataLog, error) {
	if endTime.IsZero() {
		return deremsmodels.CCDataLogs(
			qm.Where("(gw_uuid = ? and log_date > ?)", gwUUID, startTime),
			qm.OrderBy("log_date ASC")).One(repo.db)
	}

	return deremsmodels.CCDataLogs(
		qm.Where("(gw_uuid = ? and log_date > ? and log_date <= ?)", gwUUID, startTime, endTime),
		qm.OrderBy("log_date ASC")).One(repo.db)
}

// GetCCDataLogCount godoc
func (repo defaultCCDataRepository) GetCCDataLogCount() (int64, error) {
	return deremsmodels.CCDataLogs().Count(repo.db)
}
