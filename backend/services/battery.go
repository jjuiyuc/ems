package services

import (
	"time"

	log "github.com/sirupsen/logrus"

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

// BatteryService godoc
type BatteryService interface {
	GetBatteryEnergyInfo(gwUUID string, startTime time.Time) (batteryEnergyInfo *BatteryEnergyInfoResponse)
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
