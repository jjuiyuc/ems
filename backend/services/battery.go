package services

import (
	"time"

	log "github.com/sirupsen/logrus"

	"der-ems/internal/utils"
	"der-ems/repository"
)

// BatteryEnergyInfoResponse godoc
type BatteryEnergyInfoResponse struct {
	BatteryOperationCycles          float32 `json:"batteryOperationCycles"`
	BatteryLifetimeOperationCycles  float32 `json:"batteryLifetimeOperationCycles"`
	BatterySoC                      float32 `json:"batterySoC"`
	BatteryProducedEnergyAC         float32 `json:"batteryProducedEnergyAC"`
	BatteryProducedLifetimeEnergyAC float32 `json:"batteryProducedLifetimeEnergyAC"`
	BatteryConsumedEnergyAC         float32 `json:"batteryConsumedEnergyAC"`
	BatteryConsumedLifetimeEnergyAC float32 `json:"batteryConsumedLifetimeEnergyAC"`
	Model                           string  `json:"model"`
	Capcity                         float32 `json:"capcity"`
	PowerSources                    string  `json:"powerSources"`
	BatteryPower                    float32 `json:"batteryPower"`
	Voltage                         float32 `json:"voltage"`
}

// BatteryPowerStateResponse godoc
type BatteryPowerStateResponse struct {
	Timestamps             []int             `json:"timestamps"`
	BatteryAveragePowerACs []float32         `json:"batteryAveragePowerACs"`
	OnPeakTime             map[string]string `json:"onPeakTime"`
}

// BatteryChargeVoltageStateResponse godoc
type BatteryChargeVoltageStateResponse struct {
	Timestamps      []int             `json:"timestamps"`
	BatterySoCs     []float32         `json:"batterySoCs"`
	BatteryVoltages []float32         `json:"batteryVoltages"`
	OnPeakTime      map[string]string `json:"onPeakTime"`
}

// BatteryService godoc
type BatteryService interface {
	GetBatteryEnergyInfo(gwUUID string, startTime time.Time) (batteryEnergyInfo *BatteryEnergyInfoResponse)
	GetBatteryPowerState(gwUUID string, startTime, endTime time.Time) (batteryPowerState *BatteryPowerStateResponse, err error)
	GetBatteryChargeVoltageState(gwUUID string, startTime, endTime time.Time) (batteryChargeVoltageState *BatteryChargeVoltageStateResponse, err error)
}

type defaultBatteryService struct {
	repo    *repository.Repository
	billing BillingService
}

// NewBatteryService godoc
func NewBatteryService(repo *repository.Repository, billing BillingService) BatteryService {
	return &defaultBatteryService{repo, billing}
}

// GetBatteryEnergyInfo godoc
func (s defaultBatteryService) GetBatteryEnergyInfo(gwUUID string, startTime time.Time) (batteryEnergyInfo *BatteryEnergyInfoResponse) {
	batteryEnergyInfo = &BatteryEnergyInfoResponse{}
	firstlog, err1 := s.repo.CCData.GetFirstLogByGatewayUUIDAndPeriod(gwUUID, startTime, time.Time{})
	latestLog, err2 := s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod(gwUUID, time.Time{}, time.Time{})
	if err1 == nil && err2 == nil {
		log.Debug("firstlog.LogDate: ", firstlog.LogDate)
		log.Debug("latestLog.LogDate: ", latestLog.LogDate)
		batteryEnergyInfo.BatteryOperationCycles = latestLog.BatteryLifetimeOperationCycles.Float32 - firstlog.BatteryLifetimeOperationCycles.Float32
		batteryEnergyInfo.BatteryLifetimeOperationCycles = latestLog.BatteryLifetimeOperationCycles.Float32
		batteryEnergyInfo.BatterySoC = latestLog.BatterySoC.Float32
		batteryEnergyInfo.BatteryProducedEnergyAC = latestLog.BatteryProducedLifetimeEnergyAC.Float32 - firstlog.BatteryProducedLifetimeEnergyAC.Float32
		batteryEnergyInfo.BatteryProducedLifetimeEnergyAC = latestLog.BatteryProducedLifetimeEnergyAC.Float32
		batteryEnergyInfo.BatteryConsumedEnergyAC = latestLog.BatteryConsumedLifetimeEnergyAC.Float32 - firstlog.BatteryConsumedLifetimeEnergyAC.Float32
		batteryEnergyInfo.BatteryConsumedLifetimeEnergyAC = latestLog.BatteryConsumedLifetimeEnergyAC.Float32
	} else {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetFirstLogByGatewayUUIDAndPeriod and GetLatestLogByGatewayUUIDAndPeriod",
			"err1":      err1,
			"err2":      err2,
		}).Error()
		batteryEnergyInfo.BatteryOperationCycles = 0
		batteryEnergyInfo.BatteryLifetimeOperationCycles = 0
		batteryEnergyInfo.BatterySoC = 0
		batteryEnergyInfo.BatteryProducedEnergyAC = 0
		batteryEnergyInfo.BatteryProducedLifetimeEnergyAC = 0
		batteryEnergyInfo.BatteryConsumedEnergyAC = 0
		batteryEnergyInfo.BatteryConsumedLifetimeEnergyAC = 0
	}

	s.getBatteryInfo(gwUUID, batteryEnergyInfo)
	return
}

func (s defaultBatteryService) getBatteryInfo(gwUUID string, batteryEnergyInfo *BatteryEnergyInfoResponse) {
	// XXX: Hardcode battery information by gateway UUID
	const (
		Huayu      = "0324DE7B51B262F3B11A643CBA8E12CE"
		Serenegray = "04F1FD6D9C6F64C3352285CCEAF59EE1"
	)
	switch gwUUID {
	case Huayu:
		batteryEnergyInfo.Model = "PR2116 Poweroad Battery"
		batteryEnergyInfo.Capcity = 48
		batteryEnergyInfo.BatteryPower = 30
		batteryEnergyInfo.Voltage = 480
	case Serenegray:
		batteryEnergyInfo.Model = "L051100-A UZ-Energy Battery"
		batteryEnergyInfo.Capcity = 30
		batteryEnergyInfo.BatteryPower = 24
		batteryEnergyInfo.Voltage = 153.6
	}
	batteryEnergyInfo.PowerSources = "Solar + Grid"
	return
}

// GetBatteryPowerState godoc
func (s defaultBatteryService) GetBatteryPowerState(gwUUID string, periodStartTime, periodEndTime time.Time) (batteryPowerState *BatteryPowerStateResponse, err error) {
	batteryPowerState = &BatteryPowerStateResponse{}
	startTimeIndex := periodStartTime.Add(-1 * time.Hour)
	endTimeIndex := periodStartTime
	for endTimeIndex.Before(periodEndTime) || endTimeIndex == periodEndTime {
		latestLog, latestLogErr := s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod(gwUUID, startTimeIndex, endTimeIndex)
		if latestLogErr == nil {
			log.WithFields(log.Fields{
				"log_date":              latestLog.LogDate,
				"batteryAveragePowerAC": latestLog.BatteryAveragePowerAC,
			}).Debug()
			batteryPowerState.Timestamps = append(batteryPowerState.Timestamps, int(latestLog.LogDate.Unix()))
			batteryPowerState.BatteryAveragePowerACs = append(batteryPowerState.BatteryAveragePowerACs, latestLog.BatteryAveragePowerAC.Float32)
		} else {
			log.WithFields(log.Fields{
				"caused-by":      "s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod",
				"err":            latestLogErr,
				"startTimeIndex": startTimeIndex,
				"endTimeIndex":   endTimeIndex,
			}).Warn()
			batteryPowerState.Timestamps = append(batteryPowerState.Timestamps, int(endTimeIndex.Unix()))
			batteryPowerState.BatteryAveragePowerACs = append(batteryPowerState.BatteryAveragePowerACs, 0)
		}

		startTimeIndex = endTimeIndex
		endTimeIndex = startTimeIndex.Add(+1 * time.Hour)
	}

	onPeakTime, err := s.getOnPeakTime(gwUUID, startTimeIndex)
	if err != nil {
		return
	}
	batteryPowerState.OnPeakTime = onPeakTime

	return
}

// GetBatteryChargeVoltageState godoc
func (s defaultBatteryService) GetBatteryChargeVoltageState(gwUUID string, periodStartTime, periodEndTime time.Time) (batteryChargeVoltageState *BatteryChargeVoltageStateResponse, err error) {
	batteryChargeVoltageState = &BatteryChargeVoltageStateResponse{}
	startTimeIndex := periodStartTime.Add(-1 * time.Hour)
	endTimeIndex := periodStartTime
	for endTimeIndex.Before(periodEndTime) || endTimeIndex == periodEndTime {
		latestLog, latestLogErr := s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod(gwUUID, startTimeIndex, endTimeIndex)
		if latestLogErr == nil {
			log.WithFields(log.Fields{
				"log_date":              latestLog.LogDate,
				"batteryAveragePowerAC": latestLog.BatteryAveragePowerAC,
			}).Debug()
			batteryChargeVoltageState.Timestamps = append(batteryChargeVoltageState.Timestamps, int(latestLog.LogDate.Unix()))
			batteryChargeVoltageState.BatterySoCs = append(batteryChargeVoltageState.BatterySoCs, latestLog.BatterySoC.Float32)
			batteryChargeVoltageState.BatteryVoltages = append(batteryChargeVoltageState.BatteryVoltages, latestLog.BatteryVoltage.Float32)
		} else {
			log.WithFields(log.Fields{
				"caused-by":      "s.repo.CCData.GetLatestLogByGatewayUUIDAndPeriod",
				"err":            latestLogErr,
				"startTimeIndex": startTimeIndex,
				"endTimeIndex":   endTimeIndex,
			}).Warn()
			batteryChargeVoltageState.Timestamps = append(batteryChargeVoltageState.Timestamps, int(endTimeIndex.Unix()))
			batteryChargeVoltageState.BatterySoCs = append(batteryChargeVoltageState.BatterySoCs, 0)
			batteryChargeVoltageState.BatteryVoltages = append(batteryChargeVoltageState.BatteryVoltages, 0)
		}

		startTimeIndex = endTimeIndex
		endTimeIndex = startTimeIndex.Add(+1 * time.Hour)
	}

	onPeakTime, err := s.getOnPeakTime(gwUUID, startTimeIndex)
	if err != nil {
		return
	}
	batteryChargeVoltageState.OnPeakTime = onPeakTime

	return
}

func (s defaultBatteryService) getOnPeakTime(gwUUID string, t time.Time) (onPeakTime map[string]string, err error) {
	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(gwUUID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}
	billingType, err := s.billing.GetBillingTypeByCustomerID(gateway.CustomerID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.billing.GetBillingTypeByCustomerID",
			"err":       err,
		}).Error()
		return
	}
	localTime, err := s.billing.GetLocalTime(billingType.TOULocationID, t)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.billing.GetLocalTime",
			"err":       err,
		}).Error()
		return
	}
	periodType := s.billing.GetPeriodTypeOfDay(billingType.TOULocationID, localTime)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.billing.GetPeriodTypeOfDay",
			"err":       err,
		}).Error()
		return
	}
	isSummer := s.billing.IsSummer(localTime)
	billings, err := s.repo.TOU.GetBillingsByTOUInfo(billingType.TOULocationID, billingType.VoltageType, billingType.TOUType, periodType, isSummer, localTime.Format(utils.YYYYMMDD))
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.TOU.GetBillingsByTOUInfo",
			"err":       err,
		}).Error()
		return
	}

	onPeakTime = map[string]string{}
	for _, billing := range billings {
		if billing.PeakType.String == "On-peak" {
			log.WithFields(log.Fields{
				"localTime":           localTime,
				"timezone":            localTime.Format(utils.ZHHMM),
				"billing.PeakType":    billing.PeakType,
				"billing.PeriodStime": billing.PeriodStime,
				"billing.PeriodEtime": billing.PeriodEtime,
			}).Debug()
			onPeakTime["timezone"] = localTime.Format(utils.ZHHMM)
			onPeakTime["start"] = billing.PeriodStime.String
			onPeakTime["end"] = billing.PeriodEtime.String
			break
		}
	}
	return
}
