package services

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"

	"der-ems/internal/utils"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// DevicesEnergyInfoResponse godoc
type DevicesEnergyInfoResponse struct {
	GridIsPeakShaving             int                    `json:"gridIsPeakShaving"`
	LoadGridAveragePowerAC        float32                `json:"loadGridAveragePowerAC"`
	BatteryGridAveragePowerAC     float32                `json:"batteryGridAveragePowerAC"`
	GridContractPowerAC           float32                `json:"gridContractPowerAC"`
	LoadPvAveragePowerAC          float32                `json:"loadPvAveragePowerAC"`
	LoadBatteryAveragePowerAC     float32                `json:"loadBatteryAveragePowerAC"`
	BatterySoC                    float32                `json:"batterySoC"`
	BatteryProducedAveragePowerAC float32                `json:"batteryProducedAveragePowerAC"`
	BatteryConsumedAveragePowerAC float32                `json:"batteryConsumedAveragePowerAC"`
	BatteryChargingFrom           string                 `json:"batteryChargingFrom"`
	BatteryDischargingTo          string                 `json:"batteryDischargingTo"`
	PvAveragePowerAC              float32                `json:"pvAveragePowerAC"`
	LoadAveragePowerAC            float32                `json:"loadAveragePowerAC"`
	LoadLinks                     map[string]interface{} `json:"loadLinks"`
	GridLinks                     map[string]interface{} `json:"gridLinks"`
	PVLinks                       map[string]interface{} `json:"pvLinks"`
	BatteryLinks                  map[string]interface{} `json:"batteryLinks"`
	BatteryPvAveragePowerAC       float32                `json:"batteryPvAveragePowerAC"`
	GridPvAveragePowerAC          float32                `json:"gridPvAveragePowerAC"`
	GridProducedAveragePowerAC    float32                `json:"gridProducedAveragePowerAC"`
	GridConsumedAveragePowerAC    float32                `json:"gridConsumedAveragePowerAC"`
}

// EnergyDistributionInfoResponse godoc
type EnergyDistributionInfoResponse struct {
	AllProducedLifetimeEnergyACDiff     float32 `json:"allProducedLifetimeEnergyACDiff"`
	PvProducedEnergyPercentAC           float32 `json:"pvProducedEnergyPercentAC"`
	GridProducedEnergyPercentAC         float32 `json:"gridProducedEnergyPercentAC"`
	BatteryProducedEnergyPercentAC      float32 `json:"batteryProducedEnergyPercentAC"`
	PvProducedLifetimeEnergyACDiff      float32 `json:"pvProducedLifetimeEnergyACDiff"`
	GridProducedLifetimeEnergyACDiff    float32 `json:"gridProducedLifetimeEnergyACDiff"`
	BatteryProducedLifetimeEnergyACDiff float32 `json:"batteryProducedLifetimeEnergyACDiff"`
	AllConsumedLifetimeEnergyACDiff     float32 `json:"allConsumedLifetimeEnergyACDiff"`
	LoadConsumedEnergyPercentAC         float32 `json:"loadConsumedEnergyPercentAC"`
	GridConsumedEnergyPercentAC         float32 `json:"gridConsumedEnergyPercentAC"`
	BatteryConsumedEnergyPercentAC      float32 `json:"batteryConsumedEnergyPercentAC"`
	LoadConsumedLifetimeEnergyACDiff    float32 `json:"loadConsumedLifetimeEnergyACDiff"`
	GridConsumedLifetimeEnergyACDiff    float32 `json:"gridConsumedLifetimeEnergyACDiff"`
	BatteryConsumedLifetimeEnergyACDiff float32 `json:"batteryConsumedLifetimeEnergyACDiff"`
}

// PowerStateResponse godoc
type PowerStateResponse struct {
	Timestamps             []int     `json:"timestamps"`
	LoadAveragePowerACs    []float32 `json:"loadAveragePowerACs"`
	PvAveragePowerACs      []float32 `json:"pvAveragePowerACs"`
	BatteryAveragePowerACs []float32 `json:"batteryAveragePowerACs"`
	GridAveragePowerACs    []float32 `json:"gridAveragePowerACs"`
}

// LatestAccumulatedPowerState godoc
type LatestAccumulatedPowerState struct {
	Timestamps                       int
	LoadConsumedLifetimeEnergyACDiff float32
	PvProducedLifetimeEnergyACDiff   float32
	BatteryLifetimeEnergyACDiff      float32
	GridLifetimeEnergyACDiff         float32
}

// AccumulatedPowerStateResponse godoc
type AccumulatedPowerStateResponse struct {
	Timestamps                        []int     `json:"timestamps"`
	LoadConsumedLifetimeEnergyACDiffs []float32 `json:"loadConsumedLifetimeEnergyACDiffs"`
	PvProducedLifetimeEnergyACDiffs   []float32 `json:"pvProducedLifetimeEnergyACDiffs"`
	BatteryLifetimeEnergyACDiffs      []float32 `json:"batteryLifetimeEnergyACDiffs"`
	GridLifetimeEnergyACDiffs         []float32 `json:"gridLifetimeEnergyACDiffs"`
}

// DevicesService godoc
type DevicesService interface {
	GetLatestDevicesEnergyInfo(gwUUID string) (updatedTime time.Time, devicesEnergyInfo *DevicesEnergyInfoResponse, err error)
	GetEnergyDistributionInfo(gwUUID string, startTime, endTime time.Time) (energyDistributionInfo *EnergyDistributionInfoResponse)
	GetPowerState(gwUUID string, startTime, endTime time.Time) (powerState *PowerStateResponse)
	GetAccumulatedPowerState(gwUUID, resolution string, startTime, endTime time.Time) (accumulatedPowerState *AccumulatedPowerStateResponse)
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
	latestLog, err := s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod(gwUUID, time.Time{}, time.Time{})
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod",
			"err":       err,
		}).Error()
		return
	}

	logTime = latestLog.LogDate

	devicesEnergyInfo = &DevicesEnergyInfoResponse{
		GridIsPeakShaving:             latestLog.GridIsPeakShaving.Int,
		LoadGridAveragePowerAC:        latestLog.LoadGridAveragePowerAC.Float32,
		BatteryGridAveragePowerAC:     latestLog.BatteryGridAveragePowerAC.Float32,
		GridContractPowerAC:           latestLog.GridContractPowerAC.Float32,
		LoadPvAveragePowerAC:          latestLog.LoadPvAveragePowerAC.Float32,
		LoadBatteryAveragePowerAC:     latestLog.LoadBatteryAveragePowerAC.Float32,
		BatterySoC:                    latestLog.BatterySoC.Float32,
		BatteryProducedAveragePowerAC: latestLog.BatteryProducedAveragePowerAC.Float32,
		BatteryConsumedAveragePowerAC: latestLog.BatteryConsumedAveragePowerAC.Float32,
		BatteryChargingFrom:           latestLog.BatteryChargingFrom.String,
		BatteryDischargingTo:          latestLog.BatteryDischargingTo.String,
		PvAveragePowerAC:              latestLog.PvAveragePowerAC.Float32,
		LoadAveragePowerAC:            latestLog.LoadAveragePowerAC.Float32,
		BatteryPvAveragePowerAC:       latestLog.BatteryPvAveragePowerAC.Float32,
		GridPvAveragePowerAC:          latestLog.GridPvAveragePowerAC.Float32,
		GridProducedAveragePowerAC:    latestLog.GridProducedAveragePowerAC.Float32,
		GridConsumedAveragePowerAC:    latestLog.GridConsumedAveragePowerAC.Float32,
	}
	if err = json.Unmarshal(latestLog.LoadLinks.JSON, &devicesEnergyInfo.LoadLinks); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal: loadLinksValue",
			"err":       err,
		}).Error()
		return
	}
	if err = json.Unmarshal(latestLog.GridLinks.JSON, &devicesEnergyInfo.GridLinks); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal: gridLinksValue",
			"err":       err,
		}).Error()
		return
	}
	if err = json.Unmarshal(latestLog.PvLinks.JSON, &devicesEnergyInfo.PVLinks); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal: pvLinksValue",
			"err":       err,
		}).Error()
		return
	}
	if err = json.Unmarshal(latestLog.BatteryLinks.JSON, &devicesEnergyInfo.BatteryLinks); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal: batteryLinksValue",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultDevicesService) GetEnergyDistributionInfo(gwUUID string, startTime, endTime time.Time) (energyDistributionInfo *EnergyDistributionInfoResponse) {
	energyDistributionInfo = &EnergyDistributionInfoResponse{}
	firstlog, err1 := s.repo.CCData.GetFirstLogByGatewayUUIDAndPeriod(gwUUID, startTime, endTime)
	latestLog, err2 := s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod(gwUUID, startTime, endTime)

	if err1 != nil || err2 != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetFirstLogByGatewayUUIDAndPeriod and GetLatestLogByGatewayUUIDAndPeriod",
			"err1":      err1,
			"err2":      err2,
		}).Error()
		return
	}

	log.Debug("firstlog.LogDate: ", firstlog.LogDate)
	log.Debug("latestLog.LogDate: ", latestLog.LogDate)
	if firstlog.LogDate == latestLog.LogDate {
		log.WithFields(log.Fields{"caused-by": "firstlog.LogDate == latestLog.LogDate"}).Warn()
		return
	}

	energyDistributionInfo.AllProducedLifetimeEnergyACDiff = utils.Diff(latestLog.AllProducedLifetimeEnergyAC.Float32, firstlog.AllProducedLifetimeEnergyAC.Float32)
	energyDistributionInfo.PvProducedLifetimeEnergyACDiff = utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstlog.PvProducedLifetimeEnergyAC.Float32)
	energyDistributionInfo.GridProducedLifetimeEnergyACDiff = utils.Diff(latestLog.GridProducedLifetimeEnergyAC.Float32, firstlog.GridProducedLifetimeEnergyAC.Float32)
	energyDistributionInfo.BatteryProducedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryProducedLifetimeEnergyAC.Float32, firstlog.BatteryProducedLifetimeEnergyAC.Float32)
	energyDistributionInfo.PvProducedEnergyPercentAC = utils.Percent(energyDistributionInfo.PvProducedLifetimeEnergyACDiff, energyDistributionInfo.AllProducedLifetimeEnergyACDiff)
	energyDistributionInfo.GridProducedEnergyPercentAC = utils.Percent(energyDistributionInfo.GridProducedLifetimeEnergyACDiff, energyDistributionInfo.AllProducedLifetimeEnergyACDiff)
	energyDistributionInfo.BatteryProducedEnergyPercentAC = utils.Percent(energyDistributionInfo.BatteryProducedLifetimeEnergyACDiff, energyDistributionInfo.AllProducedLifetimeEnergyACDiff)

	energyDistributionInfo.AllConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.AllConsumedLifetimeEnergyAC.Float32, firstlog.AllConsumedLifetimeEnergyAC.Float32)
	energyDistributionInfo.LoadConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.LoadConsumedLifetimeEnergyAC.Float32, firstlog.LoadConsumedLifetimeEnergyAC.Float32)
	energyDistributionInfo.GridConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.GridConsumedLifetimeEnergyAC.Float32, firstlog.GridConsumedLifetimeEnergyAC.Float32)
	energyDistributionInfo.BatteryConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryConsumedLifetimeEnergyAC.Float32, firstlog.BatteryConsumedLifetimeEnergyAC.Float32)
	energyDistributionInfo.LoadConsumedEnergyPercentAC = utils.Percent(energyDistributionInfo.LoadConsumedLifetimeEnergyACDiff, energyDistributionInfo.AllConsumedLifetimeEnergyACDiff)
	energyDistributionInfo.GridConsumedEnergyPercentAC = utils.Percent(energyDistributionInfo.GridConsumedLifetimeEnergyACDiff, energyDistributionInfo.AllConsumedLifetimeEnergyACDiff)
	energyDistributionInfo.BatteryConsumedEnergyPercentAC = utils.Percent(energyDistributionInfo.BatteryConsumedLifetimeEnergyACDiff, energyDistributionInfo.AllConsumedLifetimeEnergyACDiff)
	return
}

func (s defaultDevicesService) GetPowerState(gwUUID string, startTime, endTime time.Time) (powerState *PowerStateResponse) {
	powerState = &PowerStateResponse{}
	startTimeIndex := startTime
	endTimeIndex := startTime.Add(1 * time.Hour)

	for startTimeIndex.Before(endTime) {
		latestLog, latestLogErr := s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod(gwUUID, startTimeIndex, endTimeIndex)
		if latestLogErr == nil {
			log.WithFields(log.Fields{
				"log_date":              latestLog.LogDate,
				"loadAveragePowerAC":    latestLog.LoadAveragePowerAC,
				"batteryAveragePowerAC": latestLog.BatteryAveragePowerAC,
				"pvAveragePowerAC":      latestLog.PvAveragePowerAC,
				"gridAveragePowerAC":    latestLog.GridAveragePowerAC,
			}).Debug()
			powerState.Timestamps = append(powerState.Timestamps, int(latestLog.LogDate.Unix()))
			powerState.LoadAveragePowerACs = append(powerState.LoadAveragePowerACs, latestLog.LoadAveragePowerAC.Float32)
			powerState.BatteryAveragePowerACs = append(powerState.BatteryAveragePowerACs, latestLog.BatteryAveragePowerAC.Float32)
			powerState.PvAveragePowerACs = append(powerState.PvAveragePowerACs, latestLog.PvAveragePowerAC.Float32)
			powerState.GridAveragePowerACs = append(powerState.GridAveragePowerACs, latestLog.GridAveragePowerAC.Float32)
		} else {
			log.WithFields(log.Fields{
				"caused-by":      "s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod",
				"err":            latestLogErr,
				"startTimeIndex": startTimeIndex,
				"endTimeIndex":   endTimeIndex,
			}).Warn()
			if endTimeIndex == endTime {
				powerState.Timestamps = append(powerState.Timestamps, int(endTimeIndex.Unix()))
			} else {
				powerState.Timestamps = append(powerState.Timestamps, int(endTimeIndex.Add(-1*time.Second).Unix()))
			}
			powerState.LoadAveragePowerACs = append(powerState.LoadAveragePowerACs, 0)
			powerState.BatteryAveragePowerACs = append(powerState.BatteryAveragePowerACs, 0)
			powerState.PvAveragePowerACs = append(powerState.PvAveragePowerACs, 0)
			powerState.GridAveragePowerACs = append(powerState.GridAveragePowerACs, 0)
		}

		startTimeIndex = endTimeIndex
		endTimeIndex = startTimeIndex.Add(+1 * time.Hour)
		if endTimeIndex.After(endTime) {
			endTimeIndex = endTime
		}
	}
	return
}

func (s defaultDevicesService) GetAccumulatedPowerState(gwUUID, resolution string, startTime, endTime time.Time) (accumulatedPowerState *AccumulatedPowerStateResponse) {
	accumulatedPowerState = &AccumulatedPowerStateResponse{}
	startTimeIndex := startTime
	var endTimeIndex time.Time
	switch resolution {
	case "day":
		endTimeIndex = startTime.AddDate(0, 0, 1)
	case "month":
		endTimeIndex = startTime.AddDate(0, 0, 1).AddDate(0, 1, 0).AddDate(0, 0, -1)
	}

	for startTimeIndex.Before(endTime) {
		latestAccumulatedPowerState := s.getLatestAccumulatedPowerState(gwUUID, resolution, startTimeIndex, endTimeIndex, endTime)
		log.Debug("latestAccumulatedPowerState: ", latestAccumulatedPowerState)
		accumulatedPowerState.Timestamps = append(accumulatedPowerState.Timestamps, latestAccumulatedPowerState.Timestamps)
		accumulatedPowerState.LoadConsumedLifetimeEnergyACDiffs = append(accumulatedPowerState.LoadConsumedLifetimeEnergyACDiffs, latestAccumulatedPowerState.LoadConsumedLifetimeEnergyACDiff)
		accumulatedPowerState.PvProducedLifetimeEnergyACDiffs = append(accumulatedPowerState.PvProducedLifetimeEnergyACDiffs, latestAccumulatedPowerState.PvProducedLifetimeEnergyACDiff)
		accumulatedPowerState.BatteryLifetimeEnergyACDiffs = append(accumulatedPowerState.BatteryLifetimeEnergyACDiffs, latestAccumulatedPowerState.BatteryLifetimeEnergyACDiff)
		accumulatedPowerState.GridLifetimeEnergyACDiffs = append(accumulatedPowerState.GridLifetimeEnergyACDiffs, latestAccumulatedPowerState.GridLifetimeEnergyACDiff)

		startTimeIndex = endTimeIndex
		switch resolution {
		case "day":
			endTimeIndex = startTimeIndex.AddDate(0, 0, 1)
		case "month":
			endTimeIndex = startTimeIndex.AddDate(0, 0, 1).AddDate(0, 1, 0).AddDate(0, 0, -1)
		}
		if endTimeIndex.After(endTime) {
			endTimeIndex = endTime
		}
	}
	return
}

func (s defaultDevicesService) getLatestAccumulatedPowerState(gwUUID, resolution string, startTimeIndex, endTimeIndex, endTime time.Time) (latestAccumulatedPowerState *LatestAccumulatedPowerState) {
	latestAccumulatedPowerState = &LatestAccumulatedPowerState{}
	latestLog, latestLogErr := s.repo.CCData.GetLatestCalculatedLog(gwUUID, resolution, startTimeIndex, endTimeIndex)
	if latestLogErr != nil {
		log.WithFields(log.Fields{
			"caused-by":      "s.repo.CCData.GetLatestCalculatedLog",
			"err":            latestLogErr,
			"startTimeIndex": startTimeIndex,
			"endTimeIndex":   endTimeIndex,
		}).Warn()
		if endTimeIndex == endTime {
			latestAccumulatedPowerState.Timestamps = int(endTimeIndex.Unix())
		} else {
			latestAccumulatedPowerState.Timestamps = int(endTimeIndex.Add(-1 * time.Second).Unix())
		}
		return
	}

	switch resolution {
	case "day":
		latestLogDaily, _ := (latestLog).(*deremsmodels.CCDataLogCalculatedDaily)
		latestAccumulatedPowerState.Timestamps = int(latestLogDaily.LatestLogDate.Unix())
		latestAccumulatedPowerState.LoadConsumedLifetimeEnergyACDiff = latestLogDaily.LoadConsumedLifetimeEnergyACDiff.Float32
		latestAccumulatedPowerState.PvProducedLifetimeEnergyACDiff = latestLogDaily.PvProducedLifetimeEnergyACDiff.Float32
		latestAccumulatedPowerState.BatteryLifetimeEnergyACDiff = latestLogDaily.BatteryLifetimeEnergyACDiff.Float32
		latestAccumulatedPowerState.GridLifetimeEnergyACDiff = latestLogDaily.GridLifetimeEnergyACDiff.Float32
	case "month":
		latestLogMonthly, _ := (latestLog).(*deremsmodels.CCDataLogCalculatedMonthly)
		latestAccumulatedPowerState.Timestamps = int(latestLogMonthly.LatestLogDate.Unix())
		latestAccumulatedPowerState.LoadConsumedLifetimeEnergyACDiff = latestLogMonthly.LoadConsumedLifetimeEnergyACDiff.Float32
		latestAccumulatedPowerState.PvProducedLifetimeEnergyACDiff = latestLogMonthly.PvProducedLifetimeEnergyACDiff.Float32
		latestAccumulatedPowerState.BatteryLifetimeEnergyACDiff = latestLogMonthly.BatteryLifetimeEnergyACDiff.Float32
		latestAccumulatedPowerState.GridLifetimeEnergyACDiff = latestLogMonthly.GridLifetimeEnergyACDiff.Float32
	}
	return
}
