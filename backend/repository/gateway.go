package repository

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// GatewayRepository ...
type GatewayRepository interface {
	GetCustomerIDByGatewayUUID(gwUUID string) (*deremsmodels.Gateway, error)
}

type defaultGatewayRepository struct {
	db *sql.DB
}

// NewGatewayRepository ...
func NewGatewayRepository(db *sql.DB) GatewayRepository {
	return &defaultGatewayRepository{db}
}

// GetCustomerIDByGatewayUUID ...
func (repo defaultGatewayRepository) GetCustomerIDByGatewayUUID(gwUUID string) (*deremsmodels.Gateway, error) {
	return deremsmodels.Gateways(
		qm.Where("uuid = ?", gwUUID)).One(repo.db)
}
