package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"math"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"golang.org/x/exp/slices"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

type (
	// DeviceModelType godoc
	DeviceModelType string
)

// XXX: Type Hybrid-Inverter and Inverter would not be sub-device
const (
	// HybridInverter godoc
	HybridInverter DeviceModelType = "Hybrid-Inverter"
	// Inverter godoc
	Inverter DeviceModelType = "Inverter"
)

// FakeModbusIDMapping godoc
// XXX: For sub-device, modbusID is fake and mapping from modelID
var FakeModbusIDMapping = map[int64]int{
	8: 255,
	3: 253,
	4: 252,
	5: 251,
}

// FieldManagementService godoc
type FieldManagementService interface {
	GetFields(userID int64) (getFields *GetFieldsResponse, err error)
	GetDeviceModels() (getDeviceModels *GetDeviceModelsResponse, err error)
	GetField(executedUserID int64, gwUUID string) (getField *GetFieldResponse, err error)
	AuthorizeGatewayUUID(tx *sql.Tx, executedUserID int64, gwUUID string) bool
	EnableField(executedUserID int64, gwUUID string, enable bool) (err error)
	GenerateDeviceSettings(executedUserID int64, gwUUID string) (deviceSettings *DeviceSettingsData, err error)
	GenerateDLDeviceMappingInfo(gwID int64) (data []byte, err error)
	UpdateFieldGroups(executedUserID int64, gwUUID string, groups []app.FieldGroupInfo) (err error)
	GetSubDeviceModels() (getSubDeviceModels *GetSubDeviceModelsResponse, err error)
	ValidateGatewayUUID(gwUUID string) (err error)
	ValidateDeviceUUEID(deviceUUEID string) (err error)
	CreateField(executedUserID int64, body *app.CreateFieldBody) (err error)
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
	Devices []DeviceInfo                 `json:"devices"`
	Groups  []*repository.FieldGroupWrap `json:"groups"`
}

// DeviceInfo godoc
type DeviceInfo struct {
	ModelType     string          `json:"modelType"`
	ModelID       int64           `json:"modelID"`
	ModbusID      int64           `json:"modbusID"`
	UUEID         string          `json:"uueID"`
	PowerCapacity float32         `json:"powerCapacity"`
	ExtraInfo     null.JSON       `json:"extraInfo"`
	SubDevices    []SubDeviceInfo `json:"subDevices"`
}

// SubDeviceInfo godoc
type SubDeviceInfo struct {
	ModelType     string    `json:"modelType"`
	ModelID       int64     `json:"modelID"`
	PowerCapacity float32   `json:"powerCapacity"`
	ExtraInfo     null.JSON `json:"extraInfo"`
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

// GetSubDeviceModelsResponse godoc
type GetSubDeviceModelsResponse struct {
	SubDevices []SubDevicesInfo `json:"subDevices"`
}

// SubDevicesInfo godoc
type SubDevicesInfo struct {
	Type   string               `json:"type"`
	Models []SubDeviceModelInfo `json:"models"`
}

// SubDeviceModelInfo godoc
type SubDeviceModelInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type defaultFieldManagementService struct {
	repo              *repository.Repository
	accountManagement AccountManagementService
	weather           WeatherService
	billing           BillingService
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
	models, err := s.getDeviceModelsInfo()
	if err != nil {
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

func (s defaultFieldManagementService) getDeviceModelsInfo() (models []*deremsmodels.DeviceModel, err error) {
	models, err = s.repo.Gateway.GetDeviceModels()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetDeviceModels",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) GetField(executedUserID int64, gwUUID string) (getField *GetFieldResponse, err error) {
	if !s.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(nil, gwUUID)
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

func (s defaultFieldManagementService) AuthorizeGatewayUUID(tx *sql.Tx, executedUserID int64, gwUUID string) bool {
	return s.repo.Gateway.IsGatewayExistedForUserID(tx, executedUserID, gwUUID)
}

func (s defaultFieldManagementService) getFieldResponse(executedUserID, gwID int64) (getField *GetFieldResponse, err error) {
	gatewayLocation, err := s.getGatewayLocation(gwID)
	if err != nil {
		return
	}
	fieldGroups, err := s.repo.Gateway.GetGatewayGroupsForUserID(nil, executedUserID, gwID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayGroupsForUserID",
			"err":       err,
		}).Error()
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
			ModelType:     device.ModelType,
			ModelID:       device.ModelID,
			ModbusID:      device.ModbusID,
			UUEID:         device.UUEID,
			PowerCapacity: device.PowerCapacity,
			ExtraInfo:     device.ExtraInfo,
		}

		if s.isSubModel(device.ModelType) {
			deviceInfos = append(deviceInfos, deviceInfo)
			continue
		}

		var subDeviceInfos []SubDeviceInfo
		for _, subDevice := range devices {
			if subDevice.UUEID == device.UUEID && s.isFakeModbusID(subDevice.ModbusID) {
				subDeviceInfo := SubDeviceInfo{
					ModelType:     subDevice.ModelType,
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

func (s defaultFieldManagementService) isSubModel(deviceType string) bool {
	dt := DeviceModelType(deviceType)
	return dt != HybridInverter && dt != Inverter
}

func (s defaultFieldManagementService) isFakeModbusID(modbusID int64) bool {
	// XXX: Fake modbus id decrements from 255
	return modbusID > 200
}

func (s defaultFieldManagementService) EnableField(executedUserID int64, gwUUID string, enable bool) (err error) {
	if !s.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(tx, gwUUID)
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
	if !s.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(nil, gwUUID)
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

	weatherData, err := s.weather.GetWeatherDataByLocation(location.WeatherLat.Float32, location.WeatherLng.Float32)
	if err != nil {
		return
	}
	billingData, err := s.billing.GenerateBillingParams(gateway, true)
	if err != nil {
		return
	}
	deviceMappingData, err := s.GenerateDLDeviceMappingInfo(gateway.ID)
	if err != nil {
		return
	}
	locationData, err := s.weather.GenerateGPSLocations()
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

func (s defaultFieldManagementService) UpdateFieldGroups(executedUserID int64, gwUUID string, groups []app.FieldGroupInfo) (err error) {
	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return
	}

	if !s.AuthorizeGatewayUUID(tx, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		tx.Rollback()
		return
	}
	if !s.isOwnAccountGroupChecked(tx, executedUserID, groups) {
		err = e.ErrNewOwnAccountGroupModifiedNotAllow
		tx.Rollback()
		return
	}
	if err = s.authorizeGroupIDs(tx, executedUserID, groups); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.authorizeGroupIDs",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}

	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(tx, gwUUID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}
	addedGroupIDs, deletedGroupIDs := s.getUpdatedFieldGroupIDs(tx, executedUserID, gateway.ID, groups)
	logrus.Debug("addedGroupIDs: ", addedGroupIDs)
	logrus.Debug("deletedGroupIDs: ", deletedGroupIDs)
	if err = s.addFieldGroups(tx, executedUserID, addedGroupIDs, gateway); err != nil {
		tx.Rollback()
		return
	}
	if err = s.deleteFieldGroups(tx, executedUserID, deletedGroupIDs); err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultFieldManagementService) isOwnAccountGroupChecked(tx *sql.Tx, executedUserID int64, groups []app.FieldGroupInfo) (checked bool) {
	user, err := s.repo.User.GetUserByUserID(tx, executedUserID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}
	for _, group := range groups {
		if group.ID == user.GroupID {
			if *group.Check {
				checked = true
				break
			} else {
				break
			}
		}
	}
	return
}

func (s defaultFieldManagementService) authorizeGroupIDs(tx *sql.Tx, executedUserID int64, groups []app.FieldGroupInfo) (err error) {
	executedUserGroups, err := s.repo.User.GetGroupsByUserID(tx, executedUserID)
	if err != nil {
		return
	}

	var executedUserGroupIDs, updatedGroupIDs []int64
	for _, group := range executedUserGroups {
		executedUserGroupIDs = append(executedUserGroupIDs, group.ID)
	}
	for _, group := range groups {
		updatedGroupIDs = append(updatedGroupIDs, group.ID)
	}
	if !utils.UnorderedEqualTwoArrays(executedUserGroupIDs, updatedGroupIDs) {
		err = e.ErrNewAuthPermissionNotAllow
	}
	return
}

func (s defaultFieldManagementService) getUpdatedFieldGroupIDs(tx *sql.Tx, executedUserID, gatewayID int64, groups []app.FieldGroupInfo) (addedGroupIDs, deletedGroupIDs []int64) {
	fieldGroups, err := s.repo.Gateway.GetGatewayGroupsForUserID(tx, executedUserID, gatewayID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayGroupsForUserID",
			"err":       err,
		}).Error()
		return
	}

	for _, group := range groups {
		for _, fieldGroup := range fieldGroups {
			if group.ID == fieldGroup.ID {
				if *group.Check == fieldGroup.Check {
					break
				}
				if *group.Check {
					addedGroupIDs = append(addedGroupIDs, group.ID)
				} else {
					deletedGroupIDs = append(deletedGroupIDs, group.ID)
				}
				break
			}
		}
	}
	return
}

func (s defaultFieldManagementService) addFieldGroups(tx *sql.Tx, executedUserID int64, groupIDs []int64, gateway *deremsmodels.Gateway) (err error) {
	if len(groupIDs) == 0 {
		return
	}
	for _, groupID := range groupIDs {
		addedGatewaysPermission, err := s.insertGroupGatewayPermission(tx, executedUserID, groupID, gateway.ID, gateway.LocationID)
		if err != nil {
			break
		}
		if err = s.insertGroupGatewayPermissionLog(tx, addedGatewaysPermission); err != nil {
			break
		}
	}
	return
}

func (s defaultFieldManagementService) deleteFieldGroups(tx *sql.Tx, executedUserID int64, groupIDs []int64) (err error) {
	if len(groupIDs) == 0 {
		return
	}
	for _, groupID := range groupIDs {
		deletedGatewaysPermission, err := s.deletedGroupGatewayPermission(tx, executedUserID, groupID)
		if err != nil {
			break
		}
		if err = s.insertGroupGatewayPermissionLog(tx, deletedGatewaysPermission); err != nil {
			break
		}
	}
	return
}

func (s defaultFieldManagementService) insertGroupGatewayPermission(tx *sql.Tx, executedUserID, groupID, gatewayID int64, locationID null.Int64) (addedGatewaysPermission *deremsmodels.GroupGatewayRight, err error) {
	now := time.Now().UTC()
	addedGatewaysPermission = &deremsmodels.GroupGatewayRight{
		GroupID:    groupID,
		GWID:       gatewayID,
		LocationID: locationID,
		EnabledAt:  now,
		EnabledBy:  null.Int64From(executedUserID),
		CreatedAt:  now,
		CreatedBy:  null.Int64From(executedUserID),
		UpdatedAt:  now,
		UpdatedBy:  null.Int64From(executedUserID),
	}
	if err = s.repo.User.InsertGroupGatewayPermission(tx, addedGatewaysPermission); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.InsertGroupGatewayPermission",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) deletedGroupGatewayPermission(tx *sql.Tx, executedUserID, groupID int64) (deletedGatewaysPermission *deremsmodels.GroupGatewayRight, err error) {
	now := time.Now().UTC()
	gatewaysPermission, err := s.repo.User.GetGatewaysPermissionByGroupID(tx, groupID, false)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGatewaysPermissionByGroupID",
			"err":       err,
		}).Error()
		return
	}
	deletedGatewaysPermission = gatewaysPermission[0]
	deletedGatewaysPermission.DisabledAt = null.TimeFrom(now)
	deletedGatewaysPermission.DisabledBy = null.Int64From(executedUserID)
	deletedGatewaysPermission.UpdatedAt = now
	deletedGatewaysPermission.UpdatedBy = null.Int64From(executedUserID)
	if err = s.repo.User.UpdateGroupGatewayPermission(tx, deletedGatewaysPermission); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.UpdateGroupGatewayPermission",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) insertGroupGatewayPermissionLog(tx *sql.Tx, gatewaysPermission *deremsmodels.GroupGatewayRight) (err error) {
	groupGatewayRightLog := &deremsmodels.GroupGatewayRightLog{
		GroupGatewayRightID:        null.Int64From(gatewaysPermission.ID),
		GroupID:                    null.Int64From(gatewaysPermission.GroupID),
		GWID:                       null.Int64From(gatewaysPermission.GWID),
		LocationID:                 gatewaysPermission.LocationID,
		EnabledAt:                  null.TimeFrom(gatewaysPermission.EnabledAt),
		EnabledBy:                  gatewaysPermission.EnabledBy,
		DisabledAt:                 gatewaysPermission.DisabledAt,
		DisabledBy:                 gatewaysPermission.DisabledBy,
		GroupGatewayRightUpdatedAt: null.TimeFrom(gatewaysPermission.UpdatedAt),
		GroupGatewayRightUpdatedBy: gatewaysPermission.UpdatedBy,
	}
	if err := s.repo.User.InsertGroupGatewayPermissionLog(tx, groupGatewayRightLog); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.InsertGroupGatewayPermissionLog",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) GetSubDeviceModels() (getSubDeviceModels *GetSubDeviceModelsResponse, err error) {
	models, err := s.getDeviceModelsInfo()
	if err != nil {
		return
	}

	subModels := make(map[string][]SubDeviceModelInfo)
	for _, model := range models {
		if !s.isSubModel(model.Type) {
			continue
		}

		subModelInfo := SubDeviceModelInfo{
			ID:   model.ID,
			Name: model.Name,
		}
		subModels[model.Type] = append(subModels[model.Type], subModelInfo)
	}
	var getSubDevicesInfos []SubDevicesInfo
	for subModelType, subModelInfos := range subModels {
		subDevicesInfo := SubDevicesInfo{
			Type:   subModelType,
			Models: subModelInfos,
		}
		getSubDevicesInfos = append(getSubDevicesInfos, subDevicesInfo)
	}
	getSubDeviceModels = &GetSubDeviceModelsResponse{
		SubDevices: getSubDevicesInfos,
	}
	return
}

func (s defaultFieldManagementService) ValidateGatewayUUID(gwUUID string) (err error) {
	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(nil, gwUUID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		err = e.ErrNewGatewayIDIsInvalid
		return
	}
	if s.repo.Gateway.IsGatewayBoundField(gateway) {
		err = e.ErrNewGatewayIDIsUsed
	}
	return
}

func (s defaultFieldManagementService) ValidateDeviceUUEID(deviceUUEID string) (err error) {
	deviceModule, err := s.repo.Gateway.GetDeviceModuleByDeviceUUEID(deviceUUEID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetDeviceModuleByDeviceUUEID",
			"err":       err,
		}).Error()
		err = e.ErrNewDeviceUUEIDIsInvalid
		return
	}
	if s.repo.Gateway.IsDeviceBoundField(deviceModule.ID) {
		err = e.ErrNewDeviceUUEIDIsUsed
	}
	return
}

func (s defaultFieldManagementService) CreateField(executedUserID int64, body *app.CreateFieldBody) (err error) {
	if err = s.ValidateGatewayUUID(body.GatewayID); err != nil {
		return
	}
	for _, device := range body.Devices {
		if err = s.ValidateDeviceUUEID(device.UUEID); err != nil {
			return
		}
	}
	if err = s.validateDeviceModels(body); err != nil {
		return
	}
	err = s.processCreateField(executedUserID, body)
	return
}

func (s defaultFieldManagementService) validateDeviceModels(body *app.CreateFieldBody) (err error) {
	modelIDs, subModelIDs, err := s.getDeviceModelIDsAndSubModelIDs()
	if err != nil {
		return
	}

	for _, targetDevice := range body.Devices {
		if !slices.Contains(modelIDs, targetDevice.ModelID) {
			err = e.ErrNewDeviceModelIsInvalid
			return
		}

		for _, targetSubDevice := range targetDevice.SubDevices {
			if !slices.Contains(subModelIDs, targetSubDevice.ModelID) {
				err = e.ErrNewDeviceModelIsInvalid
				return
			}
		}
	}
	return
}

func (s defaultFieldManagementService) getDeviceModelIDsAndSubModelIDs() (modelIDs, subModelIDs []int64, err error) {
	models, err := s.getDeviceModelsInfo()
	if err != nil {
		return
	}

	for _, model := range models {
		modelIDs = append(modelIDs, model.ID)
		if s.isSubModel(model.Type) {
			subModelIDs = append(subModelIDs, model.ID)
		}
	}
	return
}

func (s defaultFieldManagementService) processCreateField(executedUserID int64, body *app.CreateFieldBody) (err error) {
	logrus.Debug("gateway uuid: ", body.GatewayID)
	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	locationID, err := s.insertFieldLocation(tx, body)
	if err != nil {
		tx.Rollback()
		return
	}
	logrus.Debug("created location id: ", locationID)

	gateway, err := s.updateFieldGateway(tx, executedUserID, body, locationID)
	if err != nil {
		tx.Rollback()
		return
	}
	if err = s.updateGatewayLog(tx, gateway); err != nil {
		tx.Rollback()
		return
	}

	if err = s.insertGroupGatewayPermissionAndLog(tx, executedUserID, gateway); err != nil {
		tx.Rollback()
		return
	}

	if err = s.insertDeviceAndLog(tx, executedUserID, body, gateway.ID); err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultFieldManagementService) insertFieldLocation(tx *sql.Tx, body *app.CreateFieldBody) (locationID int64, err error) {
	now := time.Now().UTC()
	touLocation, err := s.repo.TOU.GetTOULocationByPowerCompany(tx, body.PowerCompany)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.TOU.GetTOULocationByPowerCompany",
			"err":       err,
		}).Error()
		return
	}

	location := &deremsmodels.Location{
		Name:          body.LocationName,
		Address:       null.StringFrom(body.Address),
		Lat:           null.Float64From(body.Lat),
		Lng:           null.Float64From(body.Lng),
		WeatherLat:    null.Float32From(s.getWeatherLocation(body.Lat)),
		WeatherLng:    null.Float32From(s.getWeatherLocation(body.Lng)),
		TOULocationID: null.Int64From(touLocation.ID),
		VoltageType:   null.StringFrom(body.VoltageType),
		TOUType:       null.StringFrom(body.TOUType),
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	if err = s.repo.Location.CreateLocation(tx, location); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Location.CreateLocation",
			"err":       err,
		}).Error()
		return
	}
	locationID = location.ID
	return
}

func (s defaultFieldManagementService) getWeatherLocation(value float64) float32 {
	const scale = 0.25
	quotient := int(value / scale)
	if math.Abs(math.Mod(value, scale)) > scale/2 {
		if quotient > 0 {
			return (float32)(quotient+1) * scale
		}
		return (float32)(quotient-1) * scale
	}
	return (float32)(quotient) * scale
}

func (s defaultFieldManagementService) updateFieldGateway(tx *sql.Tx, executedUserID int64, body *app.CreateFieldBody, locationID int64) (gateway *deremsmodels.Gateway, err error) {
	now := time.Now().UTC()
	gateway, err = s.repo.Gateway.GetGatewayByGatewayUUID(tx, body.GatewayID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}
	gateway.LocationID = null.Int64From(locationID)
	gateway.Enable = null.BoolFrom(*body.Enable)
	gateway.UpdatedAt = now
	gateway.UpdatedBy = null.Int64From(executedUserID)
	if err = s.repo.Gateway.UpdateGateway(tx, gateway); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.UpdateGateway",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) insertGroupGatewayPermissionAndLog(tx *sql.Tx, executedUserID int64, gateway *deremsmodels.Gateway) (err error) {
	parentGroups, err := s.repo.User.GetParentGroupsByUserID(tx, executedUserID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetParentGroupsByUserID",
			"err":       err,
		}).Error()
		return
	}

	for _, group := range parentGroups {
		logrus.Debug("permitted group id: ", group.ID)
		var permission *deremsmodels.GroupGatewayRight
		permission, err = s.insertGroupGatewayPermission(tx, executedUserID, group.ID, gateway.ID, gateway.LocationID)
		if err != nil {
			return
		}
		if err = s.insertGroupGatewayPermissionLog(tx, permission); err != nil {
			return
		}
	}
	return
}

func (s defaultFieldManagementService) insertDeviceAndLog(tx *sql.Tx, executedUserID int64, body *app.CreateFieldBody, gwID int64) (err error) {
	for _, targetDevice := range body.Devices {
		var (
			deviceModule        *deremsmodels.DeviceModule
			device              *deremsmodels.Device
			deviceExtraInfoJSON []byte
		)
		deviceModule, err = s.repo.Gateway.GetDeviceModuleByDeviceUUEID(targetDevice.UUEID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"caused-by": "s.repo.Gateway.GetDeviceModuleByDeviceUUEID",
				"err":       err,
			}).Error()
			err = e.ErrNewDeviceUUEIDIsInvalid
			return
		}
		if targetDevice.ExtraInfo != nil {
			deviceExtraInfoJSON, err = json.Marshal(targetDevice.ExtraInfo)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"caused-by": "json.Marshal",
					"err":       err,
				}).Error()
				return
			}
		}

		device, err = s.insertDevice(tx, executedUserID, gwID, deviceModule.ID, targetDevice.ModbusID, targetDevice.ModelID, targetDevice.PowerCapacity, null.JSONFrom(deviceExtraInfoJSON))
		if err != nil {
			return
		}
		if err = s.insertDeviceLog(tx, device); err != nil {
			return
		}

		for _, targetSubDevice := range targetDevice.SubDevices {
			var subDeviceExtraInfoJSON []byte
			if targetSubDevice.ExtraInfo != nil {
				subDeviceExtraInfoJSON, err = json.Marshal(targetSubDevice.ExtraInfo)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"caused-by": "json.Marshal",
						"err":       err,
					}).Error()
					return
				}
			}

			fakeModbusID, ok := FakeModbusIDMapping[targetSubDevice.ModelID]
			if !ok {
				logrus.Error("No mapping between fake modbusID and modelID: ", targetSubDevice.ModelID)
				err = e.ErrNewDeviceModelIsInvalid
				return
			}
			device, err = s.insertDevice(tx, executedUserID, gwID, deviceModule.ID, fakeModbusID, targetSubDevice.ModelID, targetSubDevice.PowerCapacity, null.JSONFrom(subDeviceExtraInfoJSON))
			if err != nil {
				return
			}
			if err = s.insertDeviceLog(tx, device); err != nil {
				return
			}
		}
	}
	return
}

func (s defaultFieldManagementService) insertDevice(tx *sql.Tx, executedUserID, gwID, moduleID int64, modbusID int, modelID int64, powerCapacity float32, extraInfo null.JSON) (device *deremsmodels.Device, err error) {
	now := time.Now().UTC()
	device = &deremsmodels.Device{
		ModbusID:      modbusID,
		ModuleID:      moduleID,
		ModelID:       modelID,
		GWID:          null.Int64From(gwID),
		PowerCapacity: powerCapacity,
		ExtraInfo:     extraInfo,
		CreatedAt:     now,
		CreatedBy:     null.Int64From(executedUserID),
		UpdatedAt:     now,
		UpdatedBy:     null.Int64From(executedUserID),
	}
	if err = s.repo.Gateway.CreateDevice(tx, device); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.CreateDevice",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultFieldManagementService) insertDeviceLog(tx *sql.Tx, device *deremsmodels.Device) (err error) {
	deviceLog := &deremsmodels.DeviceLog{
		DeviceID:        null.Int64From(device.ID),
		Modbusid:        null.IntFrom(device.ModbusID),
		ModuleID:        null.Int64From(device.ModuleID),
		ModelID:         null.Int64From(device.ModelID),
		GWID:            device.GWID,
		PowerCapacity:   null.Float32From(device.PowerCapacity),
		ExtraInfo:       device.ExtraInfo,
		Remark:          device.Remark,
		DeviceUpdatedAt: null.TimeFrom(device.UpdatedAt),
		DeviceUpdatedBy: device.UpdatedBy,
		DeviceDeletedAt: device.DeletedAt,
		DeviceDeletedBy: device.DeletedBy,
	}
	if err = s.repo.Gateway.InsertDeviceLog(tx, deviceLog); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.InsertDeviceLog",
			"err":       err,
		}).Error()
	}
	return
}
