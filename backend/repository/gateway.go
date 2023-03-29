package repository

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// GatewayRepository godoc
type GatewayRepository interface {
	GetGatewayByGatewayUUID(gwUUID string) (*deremsmodels.Gateway, error)
	GetGatewaysByLocation(lat, lng float32) ([]*deremsmodels.Gateway, error)
	GetGatewaysByUserID(userID int64) ([]*deremsmodels.Gateway, error)
	GetGateways() ([]*deremsmodels.Gateway, error)
}

type defaultGatewayRepository struct {
	db *sql.DB
}

// NewGatewayRepository godoc
func NewGatewayRepository(db *sql.DB) GatewayRepository {
	return &defaultGatewayRepository{db}
}

// GetGatewayByGatewayUUID godoc
func (repo defaultGatewayRepository) GetGatewayByGatewayUUID(gwUUID string) (*deremsmodels.Gateway, error) {
	return deremsmodels.Gateways(
		qm.Where("uuid = ?", gwUUID)).One(repo.db)
}

// GetGatewaysByLocation godoc
func (repo defaultGatewayRepository) GetGatewaysByLocation(lat, lng float32) ([]*deremsmodels.Gateway, error) {
	return deremsmodels.Gateways(
		qm.InnerJoin("location AS l ON gateway.location_id = l.id"),
		qm.Where("(l.weather_lat = ? AND l.weather_lng = ?)", lat, lng)).All(repo.db)
}

// GetGatewaysByUserID godoc
func (repo defaultGatewayRepository) GetGatewaysByUserID(userID int64) ([]*deremsmodels.Gateway, error) {
	return deremsmodels.Gateways(
		qm.InnerJoin("group_gateway_right AS gr ON gateway.id = gr.gw_id"),
		qm.InnerJoin("user AS u ON gr.group_id = u.group_id"),
		qm.Where("(u.id = ?)", userID)).All(repo.db)
}

// GetGateways godoc
func (repo defaultGatewayRepository) GetGateways() ([]*deremsmodels.Gateway, error) {
	return deremsmodels.Gateways().All(repo.db)
}
