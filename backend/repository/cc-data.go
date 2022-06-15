package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// CCDataRepository ...
type CCDataRepository interface {
	UpsertCCData(ccData *deremsmodels.CCDatum) (err error)
}

type defaultCCDataRepository struct {
	db *sql.DB
}

// NewCCDataRepository ...
func NewCCDataRepository(db *sql.DB) CCDataRepository {
	return &defaultCCDataRepository{db}
}

// UpsertCCData ...
func (repo defaultCCDataRepository) UpsertCCData(ccData *deremsmodels.CCDatum) (err error) {
	var ccDataReturn *deremsmodels.CCDatum
	ccDataReturn, err = deremsmodels.CCData(
		qm.Where("gw_uuid = ?", ccData.GWUUID),
		qm.Where("log_date = ?", ccData.LogDate)).One(repo.db)
	now := time.Now()
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
