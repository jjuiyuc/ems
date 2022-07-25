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
	GetBatteryEnergyInfo(gwUUID string) (batteryEnergyInfo *BatteryEnergyInfoResponse, err error)
}

type defaultBatteryService struct {
	repo *repository.Repository
}

// NewBatteryService godoc
func NewBatteryService(repo *repository.Repository) BatteryService {
	return &defaultBatteryService{repo}
}

// GetBatteryEnergyInfo godoc
func (s defaultBatteryService) GetBatteryEnergyInfo(gwUUID string) (batteryEnergyInfo *BatteryEnergyInfoResponse, err error) {
	currentTime := time.Now().UTC()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	firstlog, err := s.repo.CCData.GetFirstLogByGatewayUUIDAndStartTime(gwUUID, startTime)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetFirstLogByGatewayUUIDAndStartTime",
			"err":       err,
		}).Error()
		return
	}
	log.Debug("firstlog.LogDate: ", firstlog.LogDate)
	latestLog, err := s.repo.CCData.GetLatestLogByGatewayUUID(gwUUID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetLatestLogByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}
	log.Debug("latestLog.LogDate: ", latestLog.LogDate)

	batteryEnergyInfo = &BatteryEnergyInfoResponse{
		BatteryOperationCycles:          latestLog.BatteryLifetimeOperationCycles.Float32 - firstlog.BatteryLifetimeOperationCycles.Float32,
		BatteryLifetimeOperationCycles:  latestLog.BatteryLifetimeOperationCycles.Float32,
		BatterySoC:                      latestLog.BatterySoC.Float32,
		BatteryProducedEnergyAC:         latestLog.BatteryProducedLifetimeEnergyAC.Float32 - firstlog.BatteryProducedLifetimeEnergyAC.Float32,
		BatteryProducedLifetimeEnergyAC: latestLog.BatteryProducedLifetimeEnergyAC.Float32,
		BatteryConsumedEnergyAC:         latestLog.BatteryConsumedLifetimeEnergyAC.Float32 - firstlog.BatteryConsumedLifetimeEnergyAC.Float32,
		BatteryConsumedLifetimeEnergyAC: latestLog.BatteryConsumedLifetimeEnergyAC.Float32,
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
