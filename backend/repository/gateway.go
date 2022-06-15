package repository

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// GatewayRepository ...
type GatewayRepository interface {
	GetCustomerIDByGatewayUUID(gwUUID string) (*deremsmodels.Gateway, error)
	GetGatewaysByLocation(lat, lng float32) ([]*deremsmodels.Gateway, error)
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

// GetGatewaysByLocation ...
func (repo defaultGatewayRepository) GetGatewaysByLocation(lat, lng float32) ([]*deremsmodels.Gateway, error) {
	return deremsmodels.Gateways(
		qm.InnerJoin("customer AS c ON gateway.customer_id = c.id"),
		qm.Where("(c.weather_lat = ? AND c.weather_lng = ?)", lat, lng)).All(repo.db)
}
