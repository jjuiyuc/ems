package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"der-ems/internal/e"
	deremsmodels "der-ems/models/der-ems"
)

// CCDataRepository godoc
type CCDataRepository interface {
	// CC data
	UpsertCCData(ccData *deremsmodels.CCDatum) (err error)
	GetCCDataCount() (int64, error)
	// CC data log
	UpsertCCDataLog(ccDataLog *deremsmodels.CCDataLog) (err error)
	GetLatestLog(gwUUID string, startTime, endTime time.Time) (*deremsmodels.CCDataLog, error)
	GetFirstLog(gwUUID string, startTime, endTime time.Time) (*deremsmodels.CCDataLog, error)
	GetLogs(gwUUID string, startTime, endTime time.Time) ([]*deremsmodels.CCDataLog, error)
	GetCCDataLogCount() (int64, error)
	// CC data calculated log
	GetLatestCalculatedLog(gwUUID, resolution string, startTime, endTime time.Time) (interface{}, error)
	GetCalculatedLogs(gwUUID, resolution string, startTime, endTime time.Time) (interface{}, error)
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
	ccData.UpdatedAt = now
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
	ccDataLog.UpdatedAt = now
	if err != nil {
		ccDataLog.CreatedAt = now
		err = ccDataLog.Insert(repo.db, boil.Infer())
	} else {
		ccDataLog.ID = ccDataLogReturn.ID
		_, err = ccDataLog.Update(repo.db, boil.Infer())
	}
	return
}

// GetLatestLog godoc
func (repo defaultCCDataRepository) GetLatestLog(gwUUID string, startTime, endTime time.Time) (*deremsmodels.CCDataLog, error) {
	if startTime.IsZero() || endTime.IsZero() {
		return deremsmodels.CCDataLogs(
			qm.Where("gw_uuid = ?", gwUUID),
			qm.OrderBy("log_date DESC")).One(repo.db)
	}

	return deremsmodels.CCDataLogs(
		qm.Where("(gw_uuid = ? and log_date >= ? and log_date < ?)", gwUUID, startTime, endTime),
		qm.OrderBy("log_date DESC")).One(repo.db)
}

// GetFirstLog godoc
func (repo defaultCCDataRepository) GetFirstLog(gwUUID string, startTime, endTime time.Time) (*deremsmodels.CCDataLog, error) {
	if endTime.IsZero() {
		return deremsmodels.CCDataLogs(
			qm.Where("(gw_uuid = ? and log_date >= ?)", gwUUID, startTime),
			qm.OrderBy("log_date ASC")).One(repo.db)
	}

	return deremsmodels.CCDataLogs(
		qm.Where("(gw_uuid = ? and log_date >= ? and log_date < ?)", gwUUID, startTime, endTime),
		qm.OrderBy("log_date ASC")).One(repo.db)
}

// GetLogs
func (repo defaultCCDataRepository) GetLogs(gwUUID string, startTime, endTime time.Time) ([]*deremsmodels.CCDataLog, error) {
	return deremsmodels.CCDataLogs(
		qm.Where("(gw_uuid = ? and log_date >= ? and log_date < ?)", gwUUID, startTime, endTime),
		qm.OrderBy("log_date ASC")).All(repo.db)
}

// GetCCDataLogCount godoc
func (repo defaultCCDataRepository) GetCCDataLogCount() (int64, error) {
	return deremsmodels.CCDataLogs().Count(repo.db)
}

func (repo defaultCCDataRepository) GetLatestCalculatedLog(gwUUID, resolution string, startTime, endTime time.Time) (interface{}, error) {
	switch resolution {
	case "day":
		return deremsmodels.CCDataLogCalculatedDailies(
			qm.Where("(gw_uuid = ? and latest_log_date >= ? and latest_log_date < ?)", gwUUID, startTime, endTime),
			qm.OrderBy("latest_log_date DESC")).One(repo.db)
	case "month":
		return deremsmodels.CCDataLogCalculatedMonthlies(
			qm.Where("(gw_uuid = ? and latest_log_date >= ? and latest_log_date < ?)", gwUUID, startTime, endTime),
			qm.OrderBy("latest_log_date DESC")).One(repo.db)
	default:
		return nil, e.ErrNewUnexpectedResolution
	}
}

func (repo defaultCCDataRepository) GetCalculatedLogs(gwUUID, resolution string, startTime, endTime time.Time) (interface{}, error) {
	switch resolution {
	case "day":
		return deremsmodels.CCDataLogCalculatedDailies(
			qm.Where("(gw_uuid = ? and latest_log_date >= ? and latest_log_date < ?)", gwUUID, startTime, endTime),
			qm.OrderBy("latest_log_date ASC")).All(repo.db)
	case "month":
		return deremsmodels.CCDataLogCalculatedMonthlies(
			qm.Where("(gw_uuid = ? and latest_log_date >= ? and latest_log_date < ?)", gwUUID, startTime, endTime),
			qm.OrderBy("latest_log_date ASC")).All(repo.db)
	default:
		return nil, e.ErrNewUnexpectedResolution
	}
}
