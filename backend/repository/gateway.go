package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
)

// GatewayLocationWrap godoc
type GatewayLocationWrap struct {
	GatewayID    string  `json:"gatewayID"`
	LocationName string  `json:"locationName"`
	Address      string  `json:"address"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	PowerCompany string  `json:"powerCompany"`
	VoltageType  string  `json:"voltageType"`
	TOUType      string  `json:"touType"`
	Enable       bool    `json:"enable"`
}

// GPSLocationWrap godoc
type GPSLocationWrap struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

// DeviceWrap godoc
type DeviceWrap struct {
	ModelType     string    `json:"modelType"`
	ModelID       int64     `json:"modelID"`
	ModbusID      int64     `json:"modbusID"`
	UUEID         string    `json:"uueID"`
	PowerCapacity float32   `json:"powerCapacity"`
	ExtraInfo     null.JSON `json:"extraInfo"`
}

// GatewayRepository godoc
type GatewayRepository interface {
	GetGatewayByGatewayUUID(gwUUID string) (*deremsmodels.Gateway, error)
	GetGatewaysByLocation(lat, lng float32) ([]*deremsmodels.Gateway, error)
	GetGatewaysByUserID(userID int64) ([]*deremsmodels.Gateway, error)
	GetGatewayByGatewayID(gwID int64) (*deremsmodels.Gateway, error)
	GetGateways() ([]*deremsmodels.Gateway, error)
	GetGatewayLocationByGatewayID(gwID int64) (gatewayLocation GatewayLocationWrap, err error)
	GetGPSLocations() (locations []*GPSLocationWrap, err error)
	GetGatewayGroupsForUserID(executedUserID, gwID int64) ([]*deremsmodels.Group, error)
	IsGatewayExistedForUserID(executedUserID int64, gwUUID string) bool	
	GetDeviceModels() ([]*deremsmodels.DeviceModel, error)
	GetDeviceMappingByGatewayID(gwID int64) (devices []*DeviceWrap, err error)
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

func (repo defaultGatewayRepository) GetGatewayByGatewayID(gwID int64) (*deremsmodels.Gateway, error) {
	return deremsmodels.FindGateway(repo.db, gwID)
}

func (repo defaultGatewayRepository) GetGatewayLocationByGatewayID(gwID int64) (gatewayLocation GatewayLocationWrap, err error) {
	err = deremsmodels.NewQuery(
		qm.Select(
			"g.uuid AS gateway_id",
			"l.name AS location_name",
			"l.address AS address",
			"l.lat AS lat",
			"l.lng AS lng",
			"tl.power_company AS power_company",
			"l.voltage_type AS voltage_type",
			"l.tou_type AS tou_type",
			"g.enable AS enable",
		),
		qm.From("gateway AS g"),
		qm.InnerJoin("location AS l on g.location_id = l.id"),
		qm.InnerJoin("tou_location AS tl on l.tou_location_id = tl.id"),
		qm.Where("g.deleted_at IS NULL AND g.id = ?", gwID),
	).Bind(context.Background(), models.GetDB(), &gatewayLocation)
	return
}

func (repo defaultGatewayRepository) GetGPSLocations() (locations []*GPSLocationWrap, err error) {
	locations = make([]*GPSLocationWrap, 0)
	err = deremsmodels.NewQuery(
		qm.Select(
			"l.weather_lat AS lat",
			"l.weather_lng AS lng",
		),
		qm.From("location AS l"),
		qm.InnerJoin("gateway AS g ON l.id = g.location_id"),
		qm.Where("g.deleted_at IS NULL"),
		qm.GroupBy("l.weather_lat, l.weather_lng"),
	).Bind(context.Background(), models.GetDB(), &locations)
	return
}
// GetGateways godoc
func (repo defaultGatewayRepository) GetGateways() ([]*deremsmodels.Gateway, error) {
	return deremsmodels.Gateways().All(repo.db)
}

func (repo defaultGatewayRepository) GetGatewayGroupsForUserID(executedUserID, gwID int64) ([]*deremsmodels.Group, error) {
	return deremsmodels.Groups(
		qm.SQL(fmt.Sprintf(`
		WITH RECURSIVE gateway_groups AS
		(
		SELECT *
			FROM %s
			WHERE id = (
				SELECT group_id
				FROM user
				WHERE id = ?
			)
		UNION ALL
		SELECT g.*
			FROM gateway_groups AS gg JOIN %s AS g
			ON gg.id = g.parent_id
			AND g.deleted_at IS NULL
		),
		user_groups AS
		(
		SELECT %s.*
			FROM %s INNER JOIN group_gateway_right AS gr
			ON gr.gw_id = ?
			AND gr.group_id = group.id
			WHERE deleted_at IS NULL
		)
		SELECT gateway_groups.*
			FROM gateway_groups JOIN user_groups
			ON user_groups.id = gateway_groups.id;`, "`group`", "`group`", "`group`", "`group`"), executedUserID, gwID)).All(repo.db)
}

func (repo defaultGatewayRepository) IsGatewayExistedForUserID(executedUserID int64, gwUUID string) (exist bool) {
	exist, _ = deremsmodels.Gateways(
		qm.InnerJoin("group_gateway_right AS gr ON gateway.id = gr.gw_id"),
		qm.InnerJoin("user AS u ON gr.group_id = u.group_id"),
		qm.Where("uuid = ? AND u.id = ?", gwUUID, executedUserID)).Exists(repo.db)
	return
}

func (repo defaultGatewayRepository) GetDeviceModels() ([]*deremsmodels.DeviceModel, error) {
	return deremsmodels.DeviceModels().All(repo.db)
}

func (repo defaultGatewayRepository) GetDeviceMappingByGatewayID(gwID int64) (devices []*DeviceWrap, err error) {
	devices = make([]*DeviceWrap, 0)
	err = deremsmodels.NewQuery(
		qm.Select(
			"dm2.type AS model_type",
			"d.model_id AS model_id",
			"d.modbusid AS modbus_id",
			"dm.uueid AS uue_id",
			"d.power_capacity AS power_capacity",
			"d.extra_info AS extra_info",
		),
		qm.From("device AS d"),
		qm.InnerJoin("device_module AS dm ON d.module_id = dm.id"),
		qm.InnerJoin("device_model AS dm2 ON d.model_id = dm2.id"),
		qm.Where("d.deleted_at IS NULL AND d.gw_id = ?", gwID),
	).Bind(context.Background(), models.GetDB(), &devices)
	return
}
