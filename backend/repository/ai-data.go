package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// AIDataRepository godoc
type AIDataRepository interface {
	UpsertAIData(aiData *deremsmodels.AiDatum) (err error)
	GetAIDataCount() (int64, error)
}

type defaultAIDataRepository struct {
	db *sql.DB
}

// NewAIDataRepository godoc
func NewAIDataRepository(db *sql.DB) AIDataRepository {
	return &defaultAIDataRepository{db}
}

// UpsertAIData godoc
func (repo defaultAIDataRepository) UpsertAIData(aiData *deremsmodels.AiDatum) (err error) {
	var aiDataReturn *deremsmodels.AiDatum
	aiDataReturn, err = deremsmodels.AiData(
		qm.Where("gw_uuid = ?", aiData.GWUUID),
		qm.Where("log_date = ?", aiData.LogDate)).One(repo.db)
	now := time.Now().UTC()
	aiData.UpdatedAt = null.NewTime(now, true)
	if err != nil {
		aiData.CreatedAt = now
		err = aiData.Insert(repo.db, boil.Infer())
	} else {
		aiData.ID = aiDataReturn.ID
		_, err = aiData.Update(repo.db, boil.Infer())
	}
	return
}

// GetAIDataCount godoc
func (repo defaultAIDataRepository) GetAIDataCount() (int64, error) {
	return deremsmodels.AiData().Count(repo.db)
}
