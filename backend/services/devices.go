package services

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"

	"der-ems/repository"
)

// DevicesEnergyInfoResponse godoc
type DevicesEnergyInfoResponse struct {
	GridIsPeakShaving             int                    `json:"gridIsPeakShaving"`
	LoadGridAveragePowerAc        float32                `json:"loadGridAveragePowerAC"`
	BatteryGridAveragePowerAc     float32                `json:"batteryGridAveragePowerAC"`
	GridContractPowerAc           float32                `json:"gridContractPowerAC"`
	LoadPVAveragePowerAc          float32                `json:"loadPvAveragePowerAC"`
	LoadBatteryAveragePowerAc     float32                `json:"loadBatteryAveragePowerAC"`
	BatterySoc                    float32                `json:"batterySoC"`
	BatteryProducedAveragePowerAc float32                `json:"batteryProducedAveragePowerAC"`
	BatteryConsumedAveragePowerAc float32                `json:"batteryConsumedAveragePowerAC"`
	BatteryChargingFrom           string                 `json:"batteryChargingFrom"`
	BatteryDischargingTo          string                 `json:"batteryDischargingTo"`
	PVAveragePowerAc              float32                `json:"pvAveragePowerAC"`
	LoadAveragePowerAc            float32                `json:"loadAveragePowerAC"`
	LoadLinks                     map[string]interface{} `json:"loadLinks"`
	GridLinks                     map[string]interface{} `json:"gridLinks"`
	PVLinks                       map[string]interface{} `json:"pvLinks"`
	BatteryLinks                  map[string]interface{} `json:"batteryLinks"`
	BatteryPVAveragePowerAc       float32                `json:"batteryPvAveragePowerAC"`
	GridPVAveragePowerAc          float32                `json:"gridPvAveragePowerAC"`
	GridProducedAveragePowerAc    float32                `json:"gridProducedAveragePowerAC"`
	GridConsumedAveragePowerAc    float32                `json:"gridConsumedAveragePowerAC"`
}

// DevicesService godoc
type DevicesService interface {
	GetLatestDevicesEnergyInfo(gwUUID string) (updatedTime time.Time, devicesEnergyInfo *DevicesEnergyInfoResponse, err error)
}

type defaultDevicesService struct {
	repo *repository.Repository
}

// NewDevicesService godoc
func NewDevicesService(repo *repository.Repository) DevicesService {
	return &defaultDevicesService{repo}
}

// GetLatestDevicesEnergyInfo godoc
func (s defaultDevicesService) GetLatestDevicesEnergyInfo(gwUUID string) (logTime time.Time, devicesEnergyInfo *DevicesEnergyInfoResponse, err error) {
	latestLog, err := s.repo.CCData.GetLatestLogByGatewayUUID(gwUUID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetLatestLogByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}

	logTime = latestLog.LogDate

	devicesEnergyInfo = &DevicesEnergyInfoResponse{
		GridIsPeakShaving:             latestLog.GridIsPeakShaving.Int,
		LoadGridAveragePowerAc:        latestLog.LoadGridAveragePowerAc.Float32,
		BatteryGridAveragePowerAc:     latestLog.BatteryGridAveragePowerAc.Float32,
		GridContractPowerAc:           latestLog.GridContractPowerAc.Float32,
		LoadPVAveragePowerAc:          latestLog.LoadPVAveragePowerAc.Float32,
		LoadBatteryAveragePowerAc:     latestLog.LoadBatteryAveragePowerAc.Float32,
		BatterySoc:                    latestLog.BatterySoc.Float32,
		BatteryProducedAveragePowerAc: latestLog.BatteryProducedAveragePowerAc.Float32,
		BatteryConsumedAveragePowerAc: latestLog.BatteryConsumedAveragePowerAc.Float32,
		BatteryChargingFrom:           latestLog.BatteryChargingFrom.String,
		BatteryDischargingTo:          latestLog.BatteryDischargingTo.String,
		PVAveragePowerAc:              latestLog.PVAveragePowerAc.Float32,
		LoadAveragePowerAc:            latestLog.LoadAveragePowerAc.Float32,
		BatteryPVAveragePowerAc:       latestLog.BatteryPVAveragePowerAc.Float32,
		GridPVAveragePowerAc:          latestLog.GridPVAveragePowerAc.Float32,
		GridProducedAveragePowerAc:    latestLog.GridProducedAveragePowerAc.Float32,
		GridConsumedAveragePowerAc:    latestLog.GridConsumedAveragePowerAc.Float32,
	}
	var loadLinksValue map[string]interface{}
	err = json.Unmarshal(latestLog.LoadLinks.JSON, &loadLinksValue)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}
	devicesEnergyInfo.LoadLinks = loadLinksValue
	var gridLinksValue map[string]interface{}
	err = json.Unmarshal(latestLog.GridLinks.JSON, &gridLinksValue)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}
	devicesEnergyInfo.GridLinks = gridLinksValue
	var pvLinksValue map[string]interface{}
	err = json.Unmarshal(latestLog.PVLinks.JSON, &pvLinksValue)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}
	devicesEnergyInfo.PVLinks = pvLinksValue
	var batteryLinksValue map[string]interface{}
	err = json.Unmarshal(latestLog.BatteryLinks.JSON, &batteryLinksValue)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}
	devicesEnergyInfo.BatteryLinks = batteryLinksValue

	return
}
