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

const maxNumberOfPowerOutagePeriods = 12

// SettingsService godoc
type SettingsService interface {
	GetBatterySettings(executedUserID int64, gwUUID string) (getBatterySettings *GetBatterySettingsResponse, err error)
	UpdateBatterySettings(executedUserID int64, gwUUID string, body *app.UpdateBatterySettingsBody) (dlData []byte, err error)
	GetMeterSettings(executedUserID int64, gwUUID string) (getMeterSettings *GetMeterSettingsResponse, err error)
	UpdateMeterSettings(executedUserID int64, gwUUID string, body *app.UpdateMeterSettingsBody) (dlData []byte, err error)
	GetPowerOutagePeriods(executedUserID int64, gwUUID string) (getPowerOutagePeriods *GetPowerOutagePeriodsResponse, err error)
	CreatePowerOutagePeriods(executedUserID int64, gwUUID string, body *app.CreatePowerOutagePeriodsBody) (dlData []byte, err error)
}

// GetBatterySettingsResponse godoc
type GetBatterySettingsResponse struct {
	ChargingSources              string `json:"chargingSources"`
	ReservedForGridOutagePercent int    `json:"reservedForGridOutagePercent"`
}

// BatteryDLData godoc
type BatteryDLData struct {
	Type   string                     `json:"type"`
	Values GetBatterySettingsResponse `json:"values"`
}

// GetMeterSettingsResponse godoc
type GetMeterSettingsResponse struct {
	MaxDemandCapacity int `json:"maxDemandCapacity"`
}

// MeterDLData godoc
type MeterDLData struct {
	Type   string                   `json:"type"`
	Values GetMeterSettingsResponse `json:"values"`
}

// GetPowerOutagePeriodsResponse godoc
type GetPowerOutagePeriodsResponse struct {
	Periods []PowerOutagePeriodInfo `json:"periods"`
}

// PowerOutagePeriodInfo godoc
type PowerOutagePeriodInfo struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Ongoing   bool      `json:"ongoing"`
}

// PowerOutagePeriodsDLData godoc
type PowerOutagePeriodsDLData struct {
	AddedPeriods   []PeriodDLInfo `json:"addedPeriods"`
	DeletedPeriods []PeriodDLInfo `json:"deletedPeriods"`
}

// PeriodDLInfo godoc
type PeriodDLInfo struct {
	Type      string `json:"type"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
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

func (s defaultSettingsService) UpdateBatterySettings(executedUserID int64, gwUUID string, body *app.UpdateBatterySettingsBody) (dlData []byte, err error) {
	if !s.fieldManagement.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	if err = s.matchDownlinkRules(tx, gwUUID); err != nil {
		tx.Rollback()
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
	dlData, err = s.getBatteryDLData(body)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultSettingsService) matchDownlinkRules(tx *sql.Tx, gwUUID string) (err error) {
	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(tx, gwUUID)
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
	}
	return
}

func (s defaultSettingsService) getUpdateBatterySettingsInfo(tx *sql.Tx, executedUserID int64, gwUUID string, body *app.UpdateBatterySettingsBody) (device *deremsmodels.Device, err error) {
	device, err = s.getSettingsByGatewayUUIDAndType(tx, gwUUID, repository.Battery)
	if err != nil {
		return
	}
	var batterySettings repository.BatteryExtraInfo
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

func (s defaultSettingsService) getBatteryDLData(body *app.UpdateBatterySettingsBody) (dlData []byte, err error) {
	batteryDLData := &BatteryDLData{
		Type: "batterySetting",
		Values: GetBatterySettingsResponse{
			ChargingSources:              body.ChargingSources,
			ReservedForGridOutagePercent: body.ReservedForGridOutagePercent,
		},
	}
	dlData, err = json.Marshal(batteryDLData)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
		return
	}
	logrus.Debug("batteryDLDataJSON: ", string(dlData))
	return
}

func (s defaultSettingsService) GetMeterSettings(executedUserID int64, gwUUID string) (getMeterSettings *GetMeterSettingsResponse, err error) {
	if !s.fieldManagement.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}
	return s.getMeterSettingsResponse(gwUUID)
}

func (s defaultSettingsService) getMeterSettingsResponse(gwUUID string) (getMeterSettings *GetMeterSettingsResponse, err error) {
	device, err := s.getSettingsByGatewayUUIDAndType(nil, gwUUID, repository.Meter)
	if err != nil {
		return
	}
	getMeterSettings = &GetMeterSettingsResponse{
		MaxDemandCapacity: int(device.PowerCapacity),
	}
	return
}

func (s defaultSettingsService) UpdateMeterSettings(executedUserID int64, gwUUID string, body *app.UpdateMeterSettingsBody) (dlData []byte, err error) {
	if !s.fieldManagement.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	if err = s.matchDownlinkRules(tx, gwUUID); err != nil {
		tx.Rollback()
		return
	}

	device, err := s.getUpdateMeterSettingsInfo(tx, executedUserID, gwUUID, body)
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
	dlData, err = s.getMeterDLData(body)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultSettingsService) getUpdateMeterSettingsInfo(tx *sql.Tx, executedUserID int64, gwUUID string, body *app.UpdateMeterSettingsBody) (device *deremsmodels.Device, err error) {
	device, err = s.getSettingsByGatewayUUIDAndType(tx, gwUUID, repository.Meter)
	if err != nil {
		return
	}
	if device.PowerCapacity == float32(body.MaxDemandCapacity) {
		device = nil
		logrus.Warn("the-same-value-ignored")
		return
	}
	device.PowerCapacity = float32(body.MaxDemandCapacity)
	device.UpdatedAt = time.Now().UTC()
	device.UpdatedBy = null.Int64From(executedUserID)
	return
}

func (s defaultSettingsService) getMeterDLData(body *app.UpdateMeterSettingsBody) (dlData []byte, err error) {
	meterDLData := &MeterDLData{
		Type: "meterSetting",
		Values: GetMeterSettingsResponse{
			MaxDemandCapacity: body.MaxDemandCapacity,
		},
	}
	dlData, err = json.Marshal(meterDLData)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
		return
	}
	logrus.Debug("meterDLDataJSON: ", string(dlData))
	return
}

func (s defaultSettingsService) GetPowerOutagePeriods(executedUserID int64, gwUUID string) (getPowerOutagePeriods *GetPowerOutagePeriodsResponse, err error) {
	if !s.fieldManagement.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}
	return s.getPowerOutagePeriodsResponse(gwUUID)
}

func (s defaultSettingsService) getPowerOutagePeriodsResponse(gwUUID string) (getPowerOutagePeriods *GetPowerOutagePeriodsResponse, err error) {
	periods, err := s.repo.Gateway.GetPowerOutagePeriods(nil, gwUUID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetPowerOutagePeriods",
			"err":       err,
		}).Error()
		return
	}
	var getPeriods []PowerOutagePeriodInfo
	now := time.Now().UTC()
	for _, period := range periods {
		powerOutagePeriodInfo := PowerOutagePeriodInfo{
			ID:        period.ID,
			Type:      period.Type,
			StartTime: period.StartedAt,
			EndTime:   period.EndedAt,
			Ongoing:   period.StartedAt.Before(now),
		}
		getPeriods = append(getPeriods, powerOutagePeriodInfo)
	}
	getPowerOutagePeriods = &GetPowerOutagePeriodsResponse{
		Periods: getPeriods,
	}
	return
}

func (s defaultSettingsService) CreatePowerOutagePeriods(executedUserID int64, gwUUID string, body *app.CreatePowerOutagePeriodsBody) (dlData []byte, err error) {
	if !s.fieldManagement.AuthorizeGatewayUUID(nil, executedUserID, gwUUID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}

	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	if err = s.matchDownlinkRules(tx, gwUUID); err != nil {
		tx.Rollback()
		return
	}

	if err = s.matchPowerOutagePeriodsRules(tx, gwUUID, body); err != nil {
		tx.Rollback()
		return
	}

	if err = s.insertPowerOutagePeriods(tx, executedUserID, gwUUID, body); err != nil {
		tx.Rollback()
		return
	}
	dlData, err = s.getAddedPowerOutagePeriodsDLData(body)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultSettingsService) matchPowerOutagePeriodsRules(tx *sql.Tx, gwUUID string, body *app.CreatePowerOutagePeriodsBody) (err error) {
	existedPeriods, err := s.repo.Gateway.GetPowerOutagePeriods(tx, gwUUID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetPowerOutagePeriods",
			"err":       err,
		}).Error()
		return
	}
	if len(existedPeriods)+len(body.Periods) > maxNumberOfPowerOutagePeriods {
		err = e.ErrNewPowerOutagePeriodsMoreThanMaximum
		return
	}

	return s.checkAllPeriodsNotOverlapping(existedPeriods, body)
}

func (s defaultSettingsService) checkAllPeriodsNotOverlapping(existedPeriods []*deremsmodels.PowerOutagePeriod, body *app.CreatePowerOutagePeriodsBody) (err error) {
	for _, newPeriod := range body.Periods {
		for _, existedPeriod := range existedPeriods {
			if !s.checkTwoPeriodsNotOverlapping(newPeriod.StartTime, newPeriod.EndTime, existedPeriod.StartedAt, existedPeriod.EndedAt) {
				err = e.ErrNewPowerOutagePeriodInvalid
				return
			}

		}
	}
	for i, newPeriod := range body.Periods {
		for j, newPeriod2 := range body.Periods {
			if i == j {
				continue
			}
			if !s.checkTwoPeriodsNotOverlapping(newPeriod.StartTime, newPeriod.EndTime, newPeriod2.StartTime, newPeriod2.EndTime) {
				err = e.ErrNewPowerOutagePeriodInvalid
				return
			}

		}
	}
	return
}

func (s defaultSettingsService) checkTwoPeriodsNotOverlapping(newPeriodStartTime, newPeriodEndTime, existedPeriodStartTime, existedPeriodEndTime time.Time) bool {
	return (newPeriodStartTime.Before(existedPeriodStartTime) && !newPeriodEndTime.After(existedPeriodStartTime)) ||
		(!newPeriodStartTime.Before(existedPeriodEndTime) && newPeriodEndTime.After(existedPeriodEndTime))
}

func (s defaultSettingsService) insertPowerOutagePeriods(tx *sql.Tx, executedUserID int64, gwUUID string, body *app.CreatePowerOutagePeriodsBody) (err error) {
	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(tx, gwUUID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}

	now := time.Now().UTC()
	var powerOutagePeriods []*deremsmodels.PowerOutagePeriod
	for _, period := range body.Periods {
		powerOutagePeriod := &deremsmodels.PowerOutagePeriod{
			GWID:      gateway.ID,
			Type:      period.Type,
			StartedAt: period.StartTime,
			EndedAt:   period.EndTime,
			CreatedAt: now,
			CreatedBy: null.Int64From(executedUserID),
		}
		powerOutagePeriods = append(powerOutagePeriods, powerOutagePeriod)
	}
	if err = s.repo.Gateway.InsertPowerOutagePeriods(tx, powerOutagePeriods); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.Gateway.InsertPowerOutagePeriods",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultSettingsService) getAddedPowerOutagePeriodsDLData(body *app.CreatePowerOutagePeriodsBody) (dlData []byte, err error) {
	var addedPeriods []PeriodDLInfo
	for _, period := range body.Periods {
		addedPeriod := PeriodDLInfo{
			Type:      period.Type,
			StartTime: period.StartTime.Unix(),
			EndTime:   period.EndTime.Unix(),
		}
		addedPeriods = append(addedPeriods, addedPeriod)
	}
	if addedPeriods == nil {
		logrus.Warning("no-added-periods")
		return
	}

	powerOutagePeriodsDLData := PowerOutagePeriodsDLData{
		AddedPeriods: addedPeriods,
	}
	dlData, err = json.Marshal(powerOutagePeriodsDLData)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
		return
	}
	logrus.Debug("powerOutagePeriodsDLDataJSON: ", string(dlData))
	return
}
