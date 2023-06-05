package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/e"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// FieldManagementService godoc
type FieldManagementService interface {
	GetFields(userID int64) (getFields *GetFieldsResponse, err error)
	GetDeviceModels() (getDeviceModels *GetDeviceModelsResponse, err error)
	GetField(executedUserID int64, gwUUID string) (getField *GetFieldResponse, err error)
	EnableField(executedUserID int64, gwUUID string, enable bool) (err error)
	GenerateDeviceSettings(executedUserID int64, gwUUID string) (deviceSettings *DeviceSettingsData, err error)
	GenerateDLDeviceMappingInfo(gwID int64) (data []byte, err error)
}

// GetFieldsResponse godoc
type GetFieldsResponse struct {
	Gateways []GroupGatewayInfo `json:"gateways"`
}

// GetDeviceModelsResponse godoc
type GetDeviceModelsResponse struct {
	Models []DeviceModelInfo `json:"models"`
}

// DeviceModelInfo godoc
type DeviceModelInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// GetFieldResponse godoc
type GetFieldResponse struct {
	*repository.GatewayLocationWrap
	Devices []DeviceInfo     `json:"devices"`
	Groups  []FieldGroupInfo `json:"groups"`
}

// DeviceInfo godoc
type DeviceInfo struct {
	ModelID       int64           `json:"modelID"`
	ModbusID      int64           `json:"modbusID"`
	UUEID         string          `json:"uueID"`
	PowerCapacity float32         `json:"powerCapacity"`
	ExtraInfo     null.JSON       `json:"extraInfo"`
	SubDevices    []SubDeviceInfo `json:"subDevices"`
}

// SubDeviceInfo godoc
type SubDeviceInfo struct {
	ModelID       int64     `json:"modelID"`
	PowerCapacity float32   `json:"powerCapacity"`
	ExtraInfo     null.JSON `json:"extraInfo"`
}

// FieldGroupInfo godoc
type FieldGroupInfo struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	ParentID null.Int64 `json:"parentID"`
}

// DLDeviceMappingInfo godoc
type DLDeviceMappingInfo struct {
	Values []*repository.DLDeviceWrap `json:"values"`
}

// DeviceSettingsData godoc
type DeviceSettingsData struct {
	GWUUID            string
	WeatherData       []byte
	BillingData       []byte
	DeviceMappingData []byte
	LocationData      []byte
}

type defaultFieldManagementService struct {
	repo              *repository.Repository
	accountManagement AccountManagementService
	Weather           WeatherService
	Billing           BillingService
}

// NewFieldManagementService godoc
func NewFieldManagementService(repo *repository.Repository, accountManagement AccountManagementService, weather WeatherService, billing BillingService) FieldManagementService {
	return &defaultFieldManagementService{repo, accountManagement, weather, billing}
}

func (s defaultFieldManagementService) GetFields(userID int64) (getFields *GetFieldsResponse, err error) {
	user, err := s.repo.User.GetUserByUserID(nil, userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}
	gateways := s.accountManagement.GetGroupGateways(user.GroupID)
	getFields = &GetFieldsResponse{
		Gateways: gateways,
	}
	return
}

func (s defaultFieldManagementService) GetDeviceModels() (getDeviceModels *GetDeviceModelsResponse, err error) {
	models, err := s.repo.Gateway.GetDeviceModels()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetDeviceModels",
			"err":       err,
		}).Error()
		return
	}
	var getDeviceModelInfos []DeviceModelInfo
	for _, model := range models {
		deviceModelInfo := DeviceModelInfo{
			ID:   model.ID,
			Name: model.Name,
			Type: model.Type,
		}
		getDeviceModelInfos = append(getDeviceModelInfos, deviceModelInfo)
	}
	getDeviceModels = &GetDeviceModelsResponse{
		Models: getDeviceModelInfos,
	}
	return
}

func (s defaultFieldManagementService) GetField(executedUserID int64, gwUUID string) (getField *GetFieldResponse, err error) {
	if !s.authorizeGatewayUUID(executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(gwUUID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}
	getField, err = s.getFieldResponse(executedUserID, gateway.ID)
	return
}

func (s defaultFieldManagementService) authorizeGatewayUUID(executedUserID int64, gwUUID string) bool {
	return s.repo.Gateway.IsGatewayExistedForUserID(executedUserID, gwUUID)
}

func (s defaultFieldManagementService) getFieldResponse(executedUserID, gwID int64) (getField *GetFieldResponse, err error) {
	gatewayLocation, err := s.getGatewayLocation(gwID)
	if err != nil {
		return
	}
	fieldGroups, err := s.getFieldGroups(executedUserID, gwID)
	if err != nil {
		return
	}
	fieldDevices, err := s.getFieldDevices(gwID)
	if err != nil {
		return
	}
	getField = &GetFieldResponse{
		GatewayLocationWrap: &gatewayLocation,
		Groups:              fieldGroups,
		Devices:             fieldDevices,
	}
	return
}

func (s defaultFieldManagementService) getGatewayLocation(gwID int64) (gatewayLocation repository.GatewayLocationWrap, err error) {
	gatewayLocation, err = s.repo.Gateway.GetGatewayLocationByGatewayID(gwID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayLocationByGatewayID",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) getFieldGroups(executedUserID, gwID int64) (fieldGroups []FieldGroupInfo, err error) {
	groups, err := s.repo.Gateway.GetGatewayGroupsForUserID(executedUserID, gwID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayGroupsForUserID",
			"err":       err,
		}).Error()
		return
	}

	for _, group := range groups {
		fieldGroup := FieldGroupInfo{
			ID:       group.ID,
			Name:     group.Name,
			ParentID: group.ParentID,
		}
		fieldGroups = append(fieldGroups, fieldGroup)
	}
	return
}

func (s defaultFieldManagementService) getFieldDevices(gwID int64) (deviceInfos []DeviceInfo, err error) {
	devices, err := s.repo.Gateway.GetDeviceMappingByGatewayID(gwID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetDeviceMappingByGatewayID",
			"err":       err,
		}).Error()
		return
	}

	for _, device := range devices {
		if s.isFakeModbusID(device.ModbusID) {
			continue
		}
		deviceInfo := DeviceInfo{
			ModelID:       device.ModelID,
			ModbusID:      device.ModbusID,
			UUEID:         device.UUEID,
			PowerCapacity: device.PowerCapacity,
			ExtraInfo:     device.ExtraInfo,
		}

		if device.ModelType != "Hybrid-Inverter" && device.ModelType != "Inverter" {
			deviceInfos = append(deviceInfos, deviceInfo)
			continue
		}

		var subDeviceInfos []SubDeviceInfo
		for _, subDevice := range devices {
			if subDevice.UUEID == device.UUEID && s.isFakeModbusID(subDevice.ModbusID) {
				subDeviceInfo := SubDeviceInfo{
					ModelID:       subDevice.ModelID,
					PowerCapacity: subDevice.PowerCapacity,
					ExtraInfo:     subDevice.ExtraInfo,
				}
				subDeviceInfos = append(subDeviceInfos, subDeviceInfo)
			}
		}
		deviceInfo.SubDevices = subDeviceInfos
		deviceInfos = append(deviceInfos, deviceInfo)
	}
	return
}

func (s defaultFieldManagementService) isFakeModbusID(modbusID int64) bool {
	// XXX: Fake modbus id decrements from 255
	return modbusID > 200
}

func (s defaultFieldManagementService) EnableField(executedUserID int64, gwUUID string, enable bool) (err error) {
	if !s.authorizeGatewayUUID(executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(gwUUID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}
	if gateway.Enable.Bool == enable {
		logrus.Warn("value-is-the-same-ignored")
		tx.Rollback()
		return
	}
	if err = s.updateGatewayForEnableField(tx, executedUserID, enable, gateway); err != nil {
		tx.Rollback()
		return
	}
	if err = s.updateGatewayLog(tx, gateway); err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultFieldManagementService) updateGatewayForEnableField(tx *sql.Tx, executedUserID int64, enable bool, gateway *deremsmodels.Gateway) (err error) {
	gateway.Enable.Bool = enable
	gateway.UpdatedAt = time.Now().UTC()
	gateway.UpdatedBy = null.Int64From(executedUserID)
	if err = s.repo.Gateway.UpdateGateway(tx, gateway); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.UpdateGateway",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) updateGatewayLog(tx *sql.Tx, gateway *deremsmodels.Gateway) (err error) {
	gatewayLog := &deremsmodels.GatewayLog{
		GWID:        null.Int64From(gateway.ID),
		UUID:        null.StringFrom(gateway.UUID),
		LocationID:  gateway.LocationID,
		Enable:      gateway.Enable,
		Remark:      gateway.Remark,
		GWUpdatedAt: null.TimeFrom(gateway.UpdatedAt),
		GWUpdatedBy: gateway.UpdatedBy,
		GWDeletedAt: gateway.DeletedAt,
		GWDeletedBy: gateway.DeletedBy,
	}
	if err = s.repo.Gateway.InsertGatewayLog(tx, gatewayLog); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.InsertGatewayLog",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) GenerateDeviceSettings(executedUserID int64, gwUUID string) (deviceSettings *DeviceSettingsData, err error) {
	if !s.authorizeGatewayUUID(executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(gwUUID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}

	if !s.repo.Gateway.MatchDownlinkRules(gateway) {
		logrus.WithField("gateway-uuid", gateway.UUID).Warning("not-match-downlink-rules")
		err = e.ErrNewFieldIsDisabled
		return
	}

	deviceSettings, err = s.getDeviceSettings(gateway)
	return
}

func (s defaultFieldManagementService) getDeviceSettings(gateway *deremsmodels.Gateway) (deviceSettings *DeviceSettingsData, err error) {
	location, err := s.repo.Location.GetLocationByLocationID(gateway.LocationID.Int64)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Location.GetLocationByLocationID",
			"err":       err,
		}).Error()
		return
	}

	weatherData, err := s.Weather.GetWeatherDataByLocation(location.WeatherLat.Float32, location.WeatherLng.Float32)
	if err != nil {
		return
	}
	billingData, err := s.Billing.GenerateBillingParams(gateway, true)
	if err != nil {
		return
	}
	deviceMappingData, err := s.GenerateDLDeviceMappingInfo(gateway.ID)
	if err != nil {
		return
	}
	locationData, err := s.Weather.GenerateGPSLocations()
	if err != nil {
		return
	}

	deviceSettings = &DeviceSettingsData{
		GWUUID:            gateway.UUID,
		WeatherData:       weatherData,
		BillingData:       billingData,
		DeviceMappingData: deviceMappingData,
		LocationData:      locationData,
	}
	return
}

func (s defaultFieldManagementService) GenerateDLDeviceMappingInfo(gwID int64) (data []byte, err error) {
	devices, err := s.repo.Gateway.GetDLDeviceMappingByGatewayID(gwID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetDLDeviceMappingByGatewayID",
			"err":       err,
		}).Error()
		return
	}
	deviceMappingInfo := DLDeviceMappingInfo{
		Values: devices,
	}
	data, err = json.Marshal(deviceMappingInfo)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
		return
	}
	logrus.Debug("deviceMappingInfoJSON: ", string(data))
	return
}
