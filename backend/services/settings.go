package services

import (
	"database/sql"
	"encoding/json"

	"github.com/sirupsen/logrus"

	"der-ems/internal/e"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// SettingsService godoc
type SettingsService interface {
	GetBatterySettings(executedUserID int64, gwUUID string) (getBatterySettings *GetBatterySettingsResponse, err error)
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
