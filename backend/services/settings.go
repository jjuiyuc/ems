package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// SettingsService godoc
type SettingsService interface {
	GetBatterySettings(executedUserID int64, gwUUID string) (getBatterySettings *GetBatterySettingsResponse, err error)
	UpdateBatterySettings(executedUserID int64, gwUUID string, body *app.UpdateBatterySettingsBody) (err error)
}

// GetBatterySettingsResponse godoc
type GetBatterySettingsResponse struct {
	ChargingSources              string `json:"chargingSources"`
	ReservedForGridOutagePercent int    `json:"reservedForGridOutagePercent"`
}

// BatterySettings godoc
type BatterySettings struct {
	Voltage                      float32 `json:"voltage"`
	EnergyCapacity               float32 `json:"energyCapacity"`
	ChargingSources              string  `json:"chargingSources"`
	ReservedForGridOutagePercent int     `json:"reservedForGridOutagePercent"`
}

type defaultSettingsService struct {
	repo            *repository.Repository
	fieldManagement FieldManagementService
}

// NewSettingsService godoc
func NewSettingsService(repo *repository.Repository, fieldManagement FieldManagementService) SettingsService {
	return &defaultSettingsService{repo, fieldManagement}
}

func (s defaultSettingsService) GetBatterySettings(executedUserID int64, gwUUID string) (getBatterySettings *GetBatterySettingsResponse, err error) {
	if !s.fieldManagement.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}
	return s.getBatterySettingsResponse(gwUUID)
}

func (s defaultSettingsService) getBatterySettingsResponse(gwUUID string) (getBatterySettings *GetBatterySettingsResponse, err error) {
	device, err := s.getSettingsByGatewayUUIDAndType(nil, gwUUID, repository.Battery)
	if err != nil {
		return
	}
	if err = json.Unmarshal(device.ExtraInfo.JSON, &getBatterySettings); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultSettingsService) getSettingsByGatewayUUIDAndType(tx *sql.Tx, gwUUID string, modelType repository.DeviceModelType) (device *deremsmodels.Device, err error) {
	device, err = s.repo.Gateway.GetDeviceByGatewayUUIDAndType(tx, gwUUID, modelType)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetDeviceByGatewayUUIDAndType",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultSettingsService) UpdateBatterySettings(executedUserID int64, gwUUID string, body *app.UpdateBatterySettingsBody) (err error) {
	if !s.fieldManagement.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	device, err := s.getUpdateBatterySettingsInfo(tx, executedUserID, gwUUID, body)
	if err != nil || device == nil {
		tx.Rollback()
		return
	}
	if err = s.repo.Gateway.UpdateDevice(tx, device); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.UpdateDevice",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}
	if err = s.updateDeviceLog(tx, device); err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultSettingsService) getUpdateBatterySettingsInfo(tx *sql.Tx, executedUserID int64, gwUUID string, body *app.UpdateBatterySettingsBody) (device *deremsmodels.Device, err error) {
	device, err = s.getSettingsByGatewayUUIDAndType(tx, gwUUID, repository.Battery)
	if err != nil {
		return
	}
	var batterySettings BatterySettings
	if err = json.Unmarshal(device.ExtraInfo.JSON, &batterySettings); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}
	if batterySettings.ChargingSources == body.ChargingSources &&
		batterySettings.ReservedForGridOutagePercent == body.ReservedForGridOutagePercent {
		device = nil
		logrus.Warn("the-same-values-ignored")
		return
	}
	batterySettings.ChargingSources = body.ChargingSources
	batterySettings.ReservedForGridOutagePercent = body.ReservedForGridOutagePercent
	batterySettingsJSON, err := json.Marshal(batterySettings)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
		return
	}
	device.ExtraInfo = null.JSONFrom(batterySettingsJSON)
	device.UpdatedAt = time.Now().UTC()
	device.UpdatedBy = null.Int64From(executedUserID)
	return
}

func (s defaultSettingsService) updateDeviceLog(tx *sql.Tx, device *deremsmodels.Device) (err error) {
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
