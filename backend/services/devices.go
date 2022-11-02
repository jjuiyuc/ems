package services

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// LatestAccumulatedInfo godoc
type LatestAccumulatedInfo struct {
	Timestamps                       int
	LoadConsumedLifetimeEnergyACDiff float32
	PvProducedLifetimeEnergyACDiff   float32
	BatteryLifetimeEnergyACDiff      float32
	GridLifetimeEnergyACDiff         float32
	LoadSelfConsumedEnergyPercentAC  float32
}

// LatestComputedDemandState godoc
type LatestComputedDemandState struct {
	Timestamps                      int
	GridLifetimeEnergyACDiffToPower float32
	GridContractPowerAC             float32
}

// RealtimeInfo godoc
type RealtimeInfo struct {
	Timestamps             []int
	LoadAveragePowerACs    []float32
	BatteryAveragePowerACs []float32
	PvAveragePowerACs      []float32
	GridAveragePowerACs    []float32
	BatterySoCs            []float32
	BatteryVoltages        []float32
}

// AccumulatedInfo godoc
type AccumulatedInfo struct {
	Timestamps                        []int
	LoadConsumedLifetimeEnergyACDiffs []float32
	PvProducedLifetimeEnergyACDiffs   []float32
	BatteryLifetimeEnergyACDiffs      []float32
	GridLifetimeEnergyACDiffs         []float32
	LoadSelfConsumedEnergyPercentACs  []float32
}

// AbsFloat32Format godoc
type AbsFloat32Format float32

// MarshalJSON godoc
func (f AbsFloat32Format) MarshalJSON() ([]byte, error) {
	v := math.Abs(float64(f))
	v, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", v), 32)
	return json.Marshal(float32(v))
}

// Float32Format godoc
type Float32Format float32

// MarshalJSON godoc
func (f Float32Format) MarshalJSON() ([]byte, error) {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", f), 32)
	return json.Marshal(float32(v))
}

// Float32ArrayFormat godoc
type Float32ArrayFormat []float32

// MarshalJSON godoc godoc
func (f Float32ArrayFormat) MarshalJSON() ([]byte, error) {
	var formatedArray []float32
	for _, value := range f {
		v, err := strconv.ParseFloat(fmt.Sprintf("%.3f", value), 32)
		if err == nil {
			formatedArray = append(formatedArray, float32(v))
		} else {
			formatedArray = append(formatedArray, 0)
		}
	}
	return json.Marshal(formatedArray)
}

// DevicesEnergyInfoResponse godoc
type DevicesEnergyInfoResponse struct {
	GridIsPeakShaving             int                    `json:"gridIsPeakShaving"`
	LoadGridAveragePowerAC        AbsFloat32Format       `json:"loadGridAveragePowerAC"`
	BatteryGridAveragePowerAC     AbsFloat32Format       `json:"batteryGridAveragePowerAC"`
	GridContractPowerAC           Float32Format          `json:"gridContractPowerAC"`
	LoadPvAveragePowerAC          AbsFloat32Format       `json:"loadPvAveragePowerAC"`
	LoadBatteryAveragePowerAC     AbsFloat32Format       `json:"loadBatteryAveragePowerAC"`
	BatterySoC                    Float32Format          `json:"batterySoC"`
	BatteryProducedAveragePowerAC Float32Format          `json:"batteryProducedAveragePowerAC"`
	BatteryConsumedAveragePowerAC Float32Format          `json:"batteryConsumedAveragePowerAC"`
	BatteryChargingFrom           string                 `json:"batteryChargingFrom"`
	BatteryDischargingTo          string                 `json:"batteryDischargingTo"`
	PvAveragePowerAC              Float32Format          `json:"pvAveragePowerAC"`
	LoadAveragePowerAC            AbsFloat32Format       `json:"loadAveragePowerAC"`
	LoadLinks                     map[string]interface{} `json:"loadLinks"`
	GridLinks                     map[string]interface{} `json:"gridLinks"`
	PVLinks                       map[string]interface{} `json:"pvLinks"`
	BatteryLinks                  map[string]interface{} `json:"batteryLinks"`
	BatteryPvAveragePowerAC       AbsFloat32Format       `json:"batteryPvAveragePowerAC"`
	GridPvAveragePowerAC          AbsFloat32Format       `json:"gridPvAveragePowerAC"`
	GridProducedAveragePowerAC    Float32Format          `json:"gridProducedAveragePowerAC"`
	GridConsumedAveragePowerAC    Float32Format          `json:"gridConsumedAveragePowerAC"`
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
	Timestamps             []int              `json:"timestamps"`
	LoadAveragePowerACs    Float32ArrayFormat `json:"loadAveragePowerACs"`
	PvAveragePowerACs      Float32ArrayFormat `json:"pvAveragePowerACs"`
	BatteryAveragePowerACs Float32ArrayFormat `json:"batteryAveragePowerACs"`
	GridAveragePowerACs    Float32ArrayFormat `json:"gridAveragePowerACs"`
}

// AccumulatedPowerStateResponse godoc
type AccumulatedPowerStateResponse struct {
	Timestamps                        []int              `json:"timestamps"`
	LoadConsumedLifetimeEnergyACDiffs Float32ArrayFormat `json:"loadConsumedLifetimeEnergyACDiffs"`
	PvProducedLifetimeEnergyACDiffs   Float32ArrayFormat `json:"pvProducedLifetimeEnergyACDiffs"`
	BatteryLifetimeEnergyACDiffs      Float32ArrayFormat `json:"batteryLifetimeEnergyACDiffs"`
	GridLifetimeEnergyACDiffs         Float32ArrayFormat `json:"gridLifetimeEnergyACDiffs"`
}

// PowerSelfSupplyRateResponse godoc
type PowerSelfSupplyRateResponse struct {
	Timestamps                       []int              `json:"timestamps"`
	LoadSelfConsumedEnergyPercentACs Float32ArrayFormat `json:"loadSelfConsumedEnergyPercentACs"`
}

// BatteryUsageInfoResponse godoc
type BatteryUsageInfoResponse struct {
	BatterySoC                    Float32Format `json:"batterySoC"`
	BatteryProducedAveragePowerAC Float32Format `json:"batteryProducedAveragePowerAC"`
	BatteryConsumedAveragePowerAC Float32Format `json:"batteryConsumedAveragePowerAC"`
	BatteryChargingFrom           string        `json:"batteryChargingFrom"`
	BatteryDischargingTo          string        `json:"batteryDischargingTo"`
}

// ChargeInfoResponse godoc
type ChargeInfoResponse struct {
	GridPowerCost              float32       `json:"gridPowerCost"`
	GridPowerCostSavings       float32       `json:"gridPowerCostSavings"`
	GridPowerCostLastMonth     float32       `json:"gridPowerCostLastMonth"`
	GridProducedAveragePowerAC Float32Format `json:"gridProducedAveragePowerAC"`
	GridContractPowerAC        Float32Format `json:"gridContractPowerAC"`
	GridIsPeakShaving          int           `json:"gridIsPeakShaving"`
}

// DemandStateResponse godoc
type DemandStateResponse struct {
	Timestamps                       []int     `json:"timestamps"`
	GridLifetimeEnergyACDiffToPowers []float32 `json:"gridLifetimeEnergyACDiffToPowers"`
	GridContractPowerAC              float32   `json:"gridContractPowerAC"`
}

// SolarEnergyInfoResponse godoc
type SolarEnergyInfoResponse struct {
	PvProducedLifetimeEnergyACDiff        float32 `json:"pvProducedLifetimeEnergyACDiff"`
	LoadPvConsumedEnergyPercentAC         float32 `json:"loadPvConsumedEnergyPercentAC"`
	LoadPvConsumedLifetimeEnergyACDiff    float32 `json:"loadPvConsumedLifetimeEnergyACDiff"`
	BatteryPvConsumedEnergyPercentAC      float32 `json:"batteryPvConsumedEnergyPercentAC"`
	BatteryPvConsumedLifetimeEnergyACDiff float32 `json:"batteryPvConsumedLifetimeEnergyACDiff"`
	GridPvConsumedEnergyPercentAC         float32 `json:"gridPvConsumedEnergyPercentAC"`
	GridPvConsumedLifetimeEnergyACDiff    float32 `json:"gridPvConsumedLifetimeEnergyACDiff"`
	PvEnergyCostSavingsSum                float32 `json:"pvEnergyCostSavingsSum"`
	PvCo2SavingsSum                       float32 `json:"pvCo2SavingsSum"`
}

// SolarPowerStateResponse godoc
type SolarPowerStateResponse struct {
	Timestamps        []int              `json:"timestamps"`
	PvAveragePowerACs Float32ArrayFormat `json:"pvAveragePowerACs"`
	OnPeakTime        map[string]string  `json:"onPeakTime"`
}

// BatteryEnergyInfoResponse godoc
type BatteryEnergyInfoResponse struct {
	BatteryLifetimeOperationCyclesDiff  float32       `json:"batteryLifetimeOperationCyclesDiff"`
	BatteryLifetimeOperationCycles      Float32Format `json:"batteryLifetimeOperationCycles"`
	BatterySoC                          Float32Format `json:"batterySoC"`
	BatteryProducedLifetimeEnergyACDiff float32       `json:"batteryProducedLifetimeEnergyACDiff"`
	BatteryProducedLifetimeEnergyAC     Float32Format `json:"batteryProducedLifetimeEnergyAC"`
	BatteryConsumedLifetimeEnergyACDiff float32       `json:"batteryConsumedLifetimeEnergyACDiff"`
	BatteryConsumedLifetimeEnergyAC     Float32Format `json:"batteryConsumedLifetimeEnergyAC"`
	Model                               string        `json:"model"`
	Capcity                             float32       `json:"capcity"`
	PowerSources                        string        `json:"powerSources"`
	BatteryPower                        float32       `json:"batteryPower"`
	Voltage                             float32       `json:"voltage"`
}

// BatteryPowerStateResponse godoc
type BatteryPowerStateResponse struct {
	Timestamps             []int              `json:"timestamps"`
	BatteryAveragePowerACs Float32ArrayFormat `json:"batteryAveragePowerACs"`
	OnPeakTime             map[string]string  `json:"onPeakTime"`
}

// BatteryChargeVoltageStateResponse godoc
type BatteryChargeVoltageStateResponse struct {
	Timestamps      []int              `json:"timestamps"`
	BatterySoCs     Float32ArrayFormat `json:"batterySoCs"`
	BatteryVoltages Float32ArrayFormat `json:"batteryVoltages"`
	OnPeakTime      map[string]string  `json:"onPeakTime"`
}

// GridEnergyInfoResponse godoc
type GridEnergyInfoResponse struct {
	GridConsumedLifetimeEnergyACDiff float32 `json:"gridConsumedLifetimeEnergyACDiff"`
	GridProducedLifetimeEnergyACDiff float32 `json:"gridProducedLifetimeEnergyACDiff"`
	GridLifetimeEnergyACDiff         float32 `json:"gridLifetimeEnergyACDiff"`
	GridLifetimeEnergyACDiffOfMonth  float32 `json:"gridLifetimeEnergyACDiffOfMonth"`
}

// GridPowerStateResponse godoc
type GridPowerStateResponse struct {
	Timestamps          []int              `json:"timestamps"`
	GridAveragePowerACs Float32ArrayFormat `json:"gridAveragePowerACs"`
	OnPeakTime          map[string]string  `json:"onPeakTime"`
}

// DevicesService godoc
type DevicesService interface {
	GetLatestDevicesEnergyInfo(gwUUID string) (updatedTime time.Time, devicesEnergyInfo *DevicesEnergyInfoResponse, err error)
	GetEnergyDistributionInfo(param *app.PeriodParam) (energyDistributionInfo *EnergyDistributionInfoResponse)
	GetPowerState(param *app.ZoomableParam) (powerState *PowerStateResponse)
	GetAccumulatedPowerState(param *app.ResolutionWithPeriodParam) (accumulatedPowerState *AccumulatedPowerStateResponse)
	GetPowerSelfSupplyRate(param *app.ResolutionWithPeriodParam) (powerSelfSupplyRate *PowerSelfSupplyRateResponse)
	GetBatteryUsageInfo(param *app.StartTimeParam) (batteryUsageInfo *BatteryUsageInfoResponse)
	GetChargeInfo(param *app.StartTimeParam) (chargeInfo *ChargeInfoResponse)
	GetDemandState(param *app.PeriodParam) (demandState *DemandStateResponse)
	GetSolarEnergyInfo(param *app.StartTimeParam) (solarEnergyInfo *SolarEnergyInfoResponse)
	GetSolarPowerState(param *app.ZoomableParam) (solarPowerState *SolarPowerStateResponse, err error)
	GetBatteryEnergyInfo(param *app.StartTimeParam) (batteryEnergyInfo *BatteryEnergyInfoResponse)
	GetBatteryPowerState(param *app.ZoomableParam) (batteryPowerState *BatteryPowerStateResponse, err error)
	GetBatteryChargeVoltageState(param *app.ZoomableParam) (batteryChargeVoltageState *BatteryChargeVoltageStateResponse, err error)
	GetGridEnergyInfo(param *app.StartTimeParam) (gridEnergyInfo *GridEnergyInfoResponse)
	GetGridPowerState(param *app.ZoomableParam) (gridPowerState *GridPowerStateResponse, err error)
}

type defaultDevicesService struct {
	repo    *repository.Repository
	billing BillingService
}

// NewDevicesService godoc
func NewDevicesService(repo *repository.Repository, billing BillingService) DevicesService {
	return &defaultDevicesService{repo, billing}
}

// GetLatestDevicesEnergyInfo godoc
func (s defaultDevicesService) GetLatestDevicesEnergyInfo(gwUUID string) (logTime time.Time, devicesEnergyInfo *DevicesEnergyInfoResponse, err error) {
	latestLog, err := s.repo.CCData.GetLatestLog(gwUUID, time.Time{}, time.Time{})
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetLatestLog",
			"err":       err,
		}).Error()
		return
	}

	logTime = latestLog.LogDate
	devicesEnergyInfo = &DevicesEnergyInfoResponse{
		GridIsPeakShaving:             latestLog.GridIsPeakShaving.Int,
		LoadGridAveragePowerAC:        AbsFloat32Format(latestLog.LoadGridAveragePowerAC.Float32),
		BatteryGridAveragePowerAC:     AbsFloat32Format(latestLog.BatteryGridAveragePowerAC.Float32),
		GridContractPowerAC:           Float32Format(latestLog.GridContractPowerAC.Float32),
		LoadPvAveragePowerAC:          AbsFloat32Format(latestLog.LoadPvAveragePowerAC.Float32),
		LoadBatteryAveragePowerAC:     AbsFloat32Format(latestLog.LoadBatteryAveragePowerAC.Float32),
		BatterySoC:                    Float32Format(latestLog.BatterySoC.Float32),
		BatteryProducedAveragePowerAC: Float32Format(latestLog.BatteryProducedAveragePowerAC.Float32),
		BatteryConsumedAveragePowerAC: Float32Format(latestLog.BatteryConsumedAveragePowerAC.Float32),
		BatteryChargingFrom:           latestLog.BatteryChargingFrom.String,
		BatteryDischargingTo:          latestLog.BatteryDischargingTo.String,
		PvAveragePowerAC:              Float32Format(latestLog.PvAveragePowerAC.Float32),
		LoadAveragePowerAC:            AbsFloat32Format(latestLog.LoadAveragePowerAC.Float32),
		BatteryPvAveragePowerAC:       AbsFloat32Format(latestLog.BatteryPvAveragePowerAC.Float32),
		GridPvAveragePowerAC:          AbsFloat32Format(latestLog.GridPvAveragePowerAC.Float32),
		GridProducedAveragePowerAC:    Float32Format(latestLog.GridProducedAveragePowerAC.Float32),
		GridConsumedAveragePowerAC:    Float32Format(latestLog.GridConsumedAveragePowerAC.Float32),
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

func (s defaultDevicesService) GetEnergyDistributionInfo(param *app.PeriodParam) (energyDistributionInfo *EnergyDistributionInfoResponse) {
	energyDistributionInfo = &EnergyDistributionInfoResponse{}
	firstLog, err1 := s.repo.CCData.GetFirstLog(param.GatewayUUID, param.Query.StartTime, param.Query.EndTime)
	latestLog, err2 := s.repo.CCData.GetLatestLog(param.GatewayUUID, param.Query.StartTime, param.Query.EndTime)
	if err1 != nil || err2 != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetFirstLog and GetLatestLog",
			"err1":      err1,
			"err2":      err2,
		}).Error()
		return
	}

	log.WithFields(log.Fields{
		"firstLog.LogDate":  firstLog.LogDate,
		"latestLog.LogDate": latestLog.LogDate,
	}).Debug()
	energyDistributionInfo.AllProducedLifetimeEnergyACDiff = utils.Diff(latestLog.AllProducedLifetimeEnergyAC.Float32, firstLog.AllProducedLifetimeEnergyAC.Float32)
	energyDistributionInfo.PvProducedLifetimeEnergyACDiff = utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstLog.PvProducedLifetimeEnergyAC.Float32)
	energyDistributionInfo.GridProducedLifetimeEnergyACDiff = utils.Diff(latestLog.GridProducedLifetimeEnergyAC.Float32, firstLog.GridProducedLifetimeEnergyAC.Float32)
	energyDistributionInfo.BatteryProducedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryProducedLifetimeEnergyAC.Float32, firstLog.BatteryProducedLifetimeEnergyAC.Float32)
	energyDistributionInfo.PvProducedEnergyPercentAC = utils.Percent(
		energyDistributionInfo.PvProducedLifetimeEnergyACDiff,
		energyDistributionInfo.AllProducedLifetimeEnergyACDiff)
	energyDistributionInfo.GridProducedEnergyPercentAC = utils.Percent(
		energyDistributionInfo.GridProducedLifetimeEnergyACDiff,
		energyDistributionInfo.AllProducedLifetimeEnergyACDiff)
	energyDistributionInfo.BatteryProducedEnergyPercentAC = utils.Percent(
		energyDistributionInfo.BatteryProducedLifetimeEnergyACDiff,
		energyDistributionInfo.AllProducedLifetimeEnergyACDiff)
	energyDistributionInfo.AllConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.AllConsumedLifetimeEnergyAC.Float32, firstLog.AllConsumedLifetimeEnergyAC.Float32)
	energyDistributionInfo.LoadConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.LoadConsumedLifetimeEnergyAC.Float32, firstLog.LoadConsumedLifetimeEnergyAC.Float32)
	energyDistributionInfo.GridConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.GridConsumedLifetimeEnergyAC.Float32, firstLog.GridConsumedLifetimeEnergyAC.Float32)
	energyDistributionInfo.BatteryConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryConsumedLifetimeEnergyAC.Float32, firstLog.BatteryConsumedLifetimeEnergyAC.Float32)
	energyDistributionInfo.LoadConsumedEnergyPercentAC = utils.Percent(
		energyDistributionInfo.LoadConsumedLifetimeEnergyACDiff,
		energyDistributionInfo.AllConsumedLifetimeEnergyACDiff)
	energyDistributionInfo.GridConsumedEnergyPercentAC = utils.Percent(
		energyDistributionInfo.GridConsumedLifetimeEnergyACDiff,
		energyDistributionInfo.AllConsumedLifetimeEnergyACDiff)
	energyDistributionInfo.BatteryConsumedEnergyPercentAC = utils.Percent(
		energyDistributionInfo.BatteryConsumedLifetimeEnergyACDiff,
		energyDistributionInfo.AllConsumedLifetimeEnergyACDiff)
	return
}

func (s defaultDevicesService) GetPowerState(param *app.ZoomableParam) (powerState *PowerStateResponse) {
	realtimeInfo := s.getRealtimeInfo(param)
	powerState = &PowerStateResponse{
		Timestamps:             realtimeInfo.Timestamps,
		LoadAveragePowerACs:    realtimeInfo.LoadAveragePowerACs,
		BatteryAveragePowerACs: realtimeInfo.BatteryAveragePowerACs,
		PvAveragePowerACs:      realtimeInfo.PvAveragePowerACs,
		GridAveragePowerACs:    realtimeInfo.GridAveragePowerACs,
	}
	return
}

func (s defaultDevicesService) GetAccumulatedPowerState(param *app.ResolutionWithPeriodParam) (accumulatedPowerState *AccumulatedPowerStateResponse) {
	accumulatedInfo := s.getAccumulatedInfo(param)
	accumulatedPowerState = &AccumulatedPowerStateResponse{
		Timestamps:                        accumulatedInfo.Timestamps,
		LoadConsumedLifetimeEnergyACDiffs: accumulatedInfo.LoadConsumedLifetimeEnergyACDiffs,
		PvProducedLifetimeEnergyACDiffs:   accumulatedInfo.PvProducedLifetimeEnergyACDiffs,
		BatteryLifetimeEnergyACDiffs:      accumulatedInfo.BatteryLifetimeEnergyACDiffs,
		GridLifetimeEnergyACDiffs:         accumulatedInfo.GridLifetimeEnergyACDiffs,
	}
	return
}

func (s defaultDevicesService) GetChargeInfo(param *app.StartTimeParam) (chargeInfo *ChargeInfoResponse) {
	chargeInfo = &ChargeInfoResponse{}
	latestLog, err := s.repo.CCData.GetLatestLog(param.GatewayUUID, param.Query.StartTime, time.Now().UTC())
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetLatestLog",
			"err":       err,
		}).Error()
		return
	}

	log.Debug("latestLog.LogDate: ", latestLog.LogDate)
	// XXX: Hardcode gridPowerCost/gridPowerCostSavings/gridPowerCostLastMonth are return default value 0 now
	chargeInfo.GridPowerCost = 0
	chargeInfo.GridPowerCostSavings = 0
	chargeInfo.GridPowerCostLastMonth = 0
	chargeInfo.GridProducedAveragePowerAC = Float32Format(latestLog.GridProducedAveragePowerAC.Float32)
	chargeInfo.GridContractPowerAC = Float32Format(latestLog.GridContractPowerAC.Float32)
	chargeInfo.GridIsPeakShaving = latestLog.GridIsPeakShaving.Int
	return
}

func (s defaultDevicesService) GetPowerSelfSupplyRate(param *app.ResolutionWithPeriodParam) (powerSelfSupplyRate *PowerSelfSupplyRateResponse) {
	accumulatedInfo := s.getAccumulatedInfo(param)
	powerSelfSupplyRate = &PowerSelfSupplyRateResponse{
		Timestamps:                       accumulatedInfo.Timestamps,
		LoadSelfConsumedEnergyPercentACs: accumulatedInfo.LoadSelfConsumedEnergyPercentACs,
	}
	return
}

func (s defaultDevicesService) GetBatteryUsageInfo(param *app.StartTimeParam) (batteryUsageInfo *BatteryUsageInfoResponse) {
	batteryUsageInfo = &BatteryUsageInfoResponse{}
	latestLog, err := s.repo.CCData.GetLatestLog(param.GatewayUUID, param.Query.StartTime, time.Now().UTC())
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetLatestLog",
			"err":       err,
		}).Error()
		return
	}

	log.Debug("latestLog.LogDate: ", latestLog.LogDate)
	batteryUsageInfo.BatterySoC = Float32Format(latestLog.BatterySoC.Float32)
	batteryUsageInfo.BatteryProducedAveragePowerAC = Float32Format(latestLog.BatteryProducedAveragePowerAC.Float32)
	batteryUsageInfo.BatteryConsumedAveragePowerAC = Float32Format(latestLog.BatteryConsumedAveragePowerAC.Float32)
	batteryUsageInfo.BatteryChargingFrom = latestLog.BatteryChargingFrom.String
	batteryUsageInfo.BatteryDischargingTo = latestLog.BatteryDischargingTo.String
	return
}

func (s defaultDevicesService) GetDemandState(param *app.PeriodParam) (demandState *DemandStateResponse) {
	demandState = &DemandStateResponse{}
	startTimeIndex := param.Query.StartTime
	endTimeIndex := param.Query.StartTime.Add(15 * time.Minute)

	for startTimeIndex.Before(param.Query.EndTime) {
		latestComputedDemandState := s.getLatestComputedDemandState(param.GatewayUUID, startTimeIndex, endTimeIndex, param.Query.EndTime)
		log.Debug("latestComputedDemandState: ", latestComputedDemandState)
		demandState.Timestamps = append(demandState.Timestamps, latestComputedDemandState.Timestamps)
		demandState.GridLifetimeEnergyACDiffToPowers = append(demandState.GridLifetimeEnergyACDiffToPowers, latestComputedDemandState.GridLifetimeEnergyACDiffToPower)
		if latestComputedDemandState.GridContractPowerAC != 0 {
			demandState.GridContractPowerAC = latestComputedDemandState.GridContractPowerAC
		}

		startTimeIndex = endTimeIndex
		endTimeIndex = startTimeIndex.Add(15 * time.Minute)
		if endTimeIndex.After(param.Query.EndTime) {
			endTimeIndex = param.Query.EndTime
		}
	}
	return
}

func (s defaultDevicesService) GetSolarEnergyInfo(param *app.StartTimeParam) (solarEnergyInfo *SolarEnergyInfoResponse) {
	solarEnergyInfo = &SolarEnergyInfoResponse{}
	now := time.Now().UTC()
	startTimeThisMonth := param.Query.StartTime.AddDate(0, 0, -param.Query.StartTime.Day())
	firstLogOfDay, err1 := s.repo.CCData.GetFirstLog(param.GatewayUUID, param.Query.StartTime, now)
	firstLogOfMonth, err2 := s.repo.CCData.GetFirstLog(param.GatewayUUID, startTimeThisMonth, now)
	latestLog, err3 := s.repo.CCData.GetLatestLog(param.GatewayUUID, param.Query.StartTime, now)
	logsOfMonth, err4 := s.repo.CCData.GetLogs(param.GatewayUUID, startTimeThisMonth, now)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetFirstLog:Day&Month or GetLatestLog or GetLogs",
			"err1":      err1,
			"err2":      err2,
			"err3":      err3,
			"err4":      err4,
		}).Error()
		return
	}

	log.WithFields(log.Fields{
		"firstLogOfDay.LogDate":   firstLogOfDay.LogDate,
		"firstLogOfMonth.LogDate": firstLogOfMonth.LogDate,
		"latestLog.LogDate":       latestLog.LogDate,
	}).Debug()
	solarEnergyInfo.PvProducedLifetimeEnergyACDiff = utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstLogOfDay.PvProducedLifetimeEnergyAC.Float32)
	solarEnergyInfo.LoadPvConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.LoadPvConsumedLifetimeEnergyAC.Float32, firstLogOfDay.LoadPvConsumedLifetimeEnergyAC.Float32)
	solarEnergyInfo.LoadPvConsumedEnergyPercentAC = utils.Percent(
		solarEnergyInfo.LoadPvConsumedLifetimeEnergyACDiff,
		utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstLogOfDay.PvProducedLifetimeEnergyAC.Float32))
	solarEnergyInfo.BatteryPvConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryPvConsumedLifetimeEnergyAC.Float32, firstLogOfDay.BatteryPvConsumedLifetimeEnergyAC.Float32)
	solarEnergyInfo.BatteryPvConsumedEnergyPercentAC = utils.Percent(
		solarEnergyInfo.BatteryPvConsumedLifetimeEnergyACDiff,
		utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstLogOfDay.PvProducedLifetimeEnergyAC.Float32))
	solarEnergyInfo.GridPvConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.GridPvConsumedLifetimeEnergyAC.Float32, firstLogOfDay.GridPvConsumedLifetimeEnergyAC.Float32)
	solarEnergyInfo.GridPvConsumedEnergyPercentAC = utils.Percent(
		solarEnergyInfo.GridPvConsumedLifetimeEnergyACDiff,
		utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstLogOfDay.PvProducedLifetimeEnergyAC.Float32))
	var sumOfPvEnergyCostSavings, sumOfPvCo2Savings float32
	for _, logOfMonth := range logsOfMonth {
		sumOfPvEnergyCostSavings = sumOfPvEnergyCostSavings + logOfMonth.PvEnergyCostSavings.Float32
		sumOfPvCo2Savings = sumOfPvCo2Savings + logOfMonth.PvCo2Savings.Float32
	}
	solarEnergyInfo.PvEnergyCostSavingsSum = utils.ThreeDecimalPlaces(sumOfPvEnergyCostSavings)
	solarEnergyInfo.PvCo2SavingsSum = utils.ThreeDecimalPlaces(sumOfPvCo2Savings)
	return
}

func (s defaultDevicesService) GetSolarPowerState(param *app.ZoomableParam) (solarPowerState *SolarPowerStateResponse, err error) {
	solarPowerState = &SolarPowerStateResponse{}
	onPeakTime, err := s.getOnPeakTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	solarPowerState.OnPeakTime = onPeakTime
	realtimeInfo := s.getRealtimeInfo(param)
	solarPowerState.Timestamps = realtimeInfo.Timestamps
	solarPowerState.PvAveragePowerACs = realtimeInfo.PvAveragePowerACs
	return
}

func (s defaultDevicesService) GetBatteryEnergyInfo(param *app.StartTimeParam) (batteryEnergyInfo *BatteryEnergyInfoResponse) {
	batteryEnergyInfo = &BatteryEnergyInfoResponse{}
	s.getBatteryInfo(param.GatewayUUID, batteryEnergyInfo)
	firstLog, err1 := s.repo.CCData.GetFirstLog(param.GatewayUUID, param.Query.StartTime, time.Time{})
	latestLog, err2 := s.repo.CCData.GetLatestLog(param.GatewayUUID, time.Time{}, time.Time{})
	if err1 != nil || err2 != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetFirstLog and GetLatestLog",
			"err1":      err1,
			"err2":      err2,
		}).Error()
		return
	}

	log.WithFields(log.Fields{
		"firstLog.LogDate":  firstLog.LogDate,
		"latestLog.LogDate": latestLog.LogDate,
	}).Debug()
	batteryEnergyInfo.BatteryLifetimeOperationCyclesDiff = utils.Diff(latestLog.BatteryLifetimeOperationCycles.Float32, firstLog.BatteryLifetimeOperationCycles.Float32)
	batteryEnergyInfo.BatteryLifetimeOperationCycles = Float32Format(latestLog.BatteryLifetimeOperationCycles.Float32)
	batteryEnergyInfo.BatterySoC = Float32Format(latestLog.BatterySoC.Float32)
	batteryEnergyInfo.BatteryProducedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryProducedLifetimeEnergyAC.Float32, firstLog.BatteryProducedLifetimeEnergyAC.Float32)
	batteryEnergyInfo.BatteryProducedLifetimeEnergyAC = Float32Format(latestLog.BatteryProducedLifetimeEnergyAC.Float32)
	batteryEnergyInfo.BatteryConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryConsumedLifetimeEnergyAC.Float32, firstLog.BatteryConsumedLifetimeEnergyAC.Float32)
	batteryEnergyInfo.BatteryConsumedLifetimeEnergyAC = Float32Format(latestLog.BatteryConsumedLifetimeEnergyAC.Float32)
	return
}

func (s defaultDevicesService) GetBatteryPowerState(param *app.ZoomableParam) (batteryPowerState *BatteryPowerStateResponse, err error) {
	batteryPowerState = &BatteryPowerStateResponse{}
	onPeakTime, err := s.getOnPeakTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	batteryPowerState.OnPeakTime = onPeakTime
	realtimeInfo := s.getRealtimeInfo(param)
	batteryPowerState.Timestamps = realtimeInfo.Timestamps
	batteryPowerState.BatteryAveragePowerACs = realtimeInfo.BatteryAveragePowerACs
	return
}

func (s defaultDevicesService) GetBatteryChargeVoltageState(param *app.ZoomableParam) (batteryChargeVoltageState *BatteryChargeVoltageStateResponse, err error) {
	batteryChargeVoltageState = &BatteryChargeVoltageStateResponse{}
	onPeakTime, err := s.getOnPeakTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	batteryChargeVoltageState.OnPeakTime = onPeakTime
	realtimeInfo := s.getRealtimeInfo(param)
	batteryChargeVoltageState.Timestamps = realtimeInfo.Timestamps
	batteryChargeVoltageState.BatterySoCs = realtimeInfo.BatterySoCs
	batteryChargeVoltageState.BatteryVoltages = realtimeInfo.BatteryVoltages
	return
}

func (s defaultDevicesService) GetGridEnergyInfo(param *app.StartTimeParam) (gridEnergyInfo *GridEnergyInfoResponse) {
	gridEnergyInfo = &GridEnergyInfoResponse{}
	now := time.Now().UTC()
	startTimeThisMonth := param.Query.StartTime.AddDate(0, 0, -param.Query.StartTime.Day())
	firstLogOfDay, err1 := s.repo.CCData.GetFirstLog(param.GatewayUUID, param.Query.StartTime, now)
	firstLogOfMonth, err2 := s.repo.CCData.GetFirstLog(param.GatewayUUID, startTimeThisMonth, now)
	latestLog, err3 := s.repo.CCData.GetLatestLog(param.GatewayUUID, param.Query.StartTime, now)
	if err1 != nil || err2 != nil || err3 != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetFirstLog:Day&Month and GetLatestLog",
			"err1":      err1,
			"err2":      err2,
			"err3":      err3,
		}).Error()
		return
	}

	log.WithFields(log.Fields{
		"firstLogOfDay.LogDate":   firstLogOfDay.LogDate,
		"firstLogOfMonth.LogDate": firstLogOfMonth.LogDate,
		"latestLog.LogDate":       latestLog.LogDate,
	}).Debug()
	gridEnergyInfo.GridConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.GridConsumedLifetimeEnergyAC.Float32, firstLogOfDay.GridConsumedLifetimeEnergyAC.Float32)
	gridEnergyInfo.GridProducedLifetimeEnergyACDiff = utils.Diff(latestLog.GridProducedLifetimeEnergyAC.Float32, firstLogOfDay.GridProducedLifetimeEnergyAC.Float32)
	gridEnergyInfo.GridLifetimeEnergyACDiff = utils.Diff(latestLog.GridLifetimeEnergyAC.Float32, firstLogOfDay.GridLifetimeEnergyAC.Float32)
	gridEnergyInfo.GridLifetimeEnergyACDiffOfMonth = utils.Diff(latestLog.GridLifetimeEnergyAC.Float32, firstLogOfMonth.GridLifetimeEnergyAC.Float32)
	return
}

func (s defaultDevicesService) GetGridPowerState(param *app.ZoomableParam) (gridPowerState *GridPowerStateResponse, err error) {
	gridPowerState = &GridPowerStateResponse{}
	onPeakTime, err := s.getOnPeakTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	gridPowerState.OnPeakTime = onPeakTime
	realtimeInfo := s.getRealtimeInfo(param)
	gridPowerState.Timestamps = realtimeInfo.Timestamps
	gridPowerState.GridAveragePowerACs = realtimeInfo.GridAveragePowerACs
	return
}

func (s defaultDevicesService) getRealtimeInfo(param *app.ZoomableParam) (realtimeInfo *RealtimeInfo) {
	realtimeInfo = &RealtimeInfo{}
	startTimeIndex := param.Query.StartTime
	endTimeIndex := param.GetEndTimeIndex()

	for startTimeIndex.Before(param.Query.EndTime) {
		latestRealtimeInfo := s.getLatestRealtimeInfo(param.GatewayUUID, startTimeIndex, endTimeIndex, param.Query.EndTime)
		log.Debug("latestRealtimeInfo.LogDate: ", latestRealtimeInfo.LogDate)
		realtimeInfo.Timestamps = append(realtimeInfo.Timestamps, int(latestRealtimeInfo.LogDate.Unix()))
		realtimeInfo.LoadAveragePowerACs = append(realtimeInfo.LoadAveragePowerACs, latestRealtimeInfo.LoadAveragePowerAC.Float32)
		realtimeInfo.BatteryAveragePowerACs = append(realtimeInfo.BatteryAveragePowerACs, latestRealtimeInfo.BatteryAveragePowerAC.Float32)
		realtimeInfo.PvAveragePowerACs = append(realtimeInfo.PvAveragePowerACs, latestRealtimeInfo.PvAveragePowerAC.Float32)
		realtimeInfo.GridAveragePowerACs = append(realtimeInfo.GridAveragePowerACs, latestRealtimeInfo.GridAveragePowerAC.Float32)
		realtimeInfo.BatterySoCs = append(realtimeInfo.BatterySoCs, latestRealtimeInfo.BatterySoC.Float32)
		realtimeInfo.BatteryVoltages = append(realtimeInfo.BatteryVoltages, latestRealtimeInfo.BatteryVoltage.Float32)

		startTimeIndex = endTimeIndex
		switch param.Query.Resolution {
		case "hour":
			endTimeIndex = startTimeIndex.Add(1 * time.Hour)
		case "5minute":
			endTimeIndex = startTimeIndex.Add(5 * time.Minute)
		}
		if endTimeIndex.After(param.Query.EndTime) {
			endTimeIndex = param.Query.EndTime
		}
	}
	return
}

func (s defaultDevicesService) getAccumulatedInfo(param *app.ResolutionWithPeriodParam) (accumulatedInfo *AccumulatedInfo) {
	accumulatedInfo = &AccumulatedInfo{}
	startTimeIndex := param.Query.StartTime
	endTimeIndex := param.GetEndTimeIndex()

	for startTimeIndex.Before(param.Query.EndTime) {
		latestAccumulatedInfo := s.getLatestAccumulatedInfo(param.GatewayUUID, param.Query.Resolution, startTimeIndex, endTimeIndex, param.Query.EndTime)
		log.Debug("latestAccumulatedInfo: ", latestAccumulatedInfo)
		accumulatedInfo.Timestamps = append(accumulatedInfo.Timestamps, latestAccumulatedInfo.Timestamps)
		accumulatedInfo.LoadConsumedLifetimeEnergyACDiffs = append(accumulatedInfo.LoadConsumedLifetimeEnergyACDiffs, latestAccumulatedInfo.LoadConsumedLifetimeEnergyACDiff)
		accumulatedInfo.PvProducedLifetimeEnergyACDiffs = append(accumulatedInfo.PvProducedLifetimeEnergyACDiffs, latestAccumulatedInfo.PvProducedLifetimeEnergyACDiff)
		accumulatedInfo.BatteryLifetimeEnergyACDiffs = append(accumulatedInfo.BatteryLifetimeEnergyACDiffs, latestAccumulatedInfo.BatteryLifetimeEnergyACDiff)
		accumulatedInfo.GridLifetimeEnergyACDiffs = append(accumulatedInfo.GridLifetimeEnergyACDiffs, latestAccumulatedInfo.GridLifetimeEnergyACDiff)
		accumulatedInfo.LoadSelfConsumedEnergyPercentACs = append(accumulatedInfo.LoadSelfConsumedEnergyPercentACs, latestAccumulatedInfo.LoadSelfConsumedEnergyPercentAC)

		startTimeIndex = endTimeIndex
		switch param.Query.Resolution {
		case "day":
			endTimeIndex = startTimeIndex.AddDate(0, 0, 1)
		case "month":
			endTimeIndex = startTimeIndex.AddDate(0, 0, 1).AddDate(0, 1, 0).AddDate(0, 0, -1)
		}
		if endTimeIndex.After(param.Query.EndTime) {
			endTimeIndex = param.Query.EndTime
		}
	}
	return
}

func (s defaultDevicesService) getLatestRealtimeInfo(gwUUID string, startTimeIndex, endTimeIndex, endTime time.Time) (latestLog *deremsmodels.CCDataLog) {
	latestLog, err := s.repo.CCData.GetLatestLog(gwUUID, startTimeIndex, endTimeIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by":      "s.repo.CCData.GetLatestLog",
			"err":            err,
			"startTimeIndex": startTimeIndex,
			"endTimeIndex":   endTimeIndex,
		}).Error()
		latestLog = &deremsmodels.CCDataLog{
			LogDate: endTimeIndex.Add(-1 * time.Second),
		}
	}
	return
}

func (s defaultDevicesService) getLatestAccumulatedInfo(gwUUID, resolution string, startTimeIndex, endTimeIndex, endTime time.Time) (latestAccumulatedInfo *LatestAccumulatedInfo) {
	latestAccumulatedInfo = &LatestAccumulatedInfo{}
	latestLog, err := s.repo.CCData.GetLatestCalculatedLog(gwUUID, resolution, startTimeIndex, endTimeIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by":      "s.repo.CCData.GetLatestCalculatedLog",
			"err":            err,
			"startTimeIndex": startTimeIndex,
			"endTimeIndex":   endTimeIndex,
		}).Error()
		latestAccumulatedInfo.Timestamps = int(endTimeIndex.Add(-1 * time.Second).Unix())
		return
	}

	switch resolution {
	case "day":
		latestLogDaily, _ := (latestLog).(*deremsmodels.CCDataLogCalculatedDaily)
		latestAccumulatedInfo.Timestamps = int(latestLogDaily.LatestLogDate.Unix())
		latestAccumulatedInfo.LoadConsumedLifetimeEnergyACDiff = latestLogDaily.LoadConsumedLifetimeEnergyACDiff.Float32
		latestAccumulatedInfo.PvProducedLifetimeEnergyACDiff = latestLogDaily.PvProducedLifetimeEnergyACDiff.Float32
		latestAccumulatedInfo.BatteryLifetimeEnergyACDiff = latestLogDaily.BatteryLifetimeEnergyACDiff.Float32
		latestAccumulatedInfo.GridLifetimeEnergyACDiff = latestLogDaily.GridLifetimeEnergyACDiff.Float32
		latestAccumulatedInfo.LoadSelfConsumedEnergyPercentAC = latestLogDaily.LoadSelfConsumedEnergyPercentAC.Float32
	case "month":
		latestLogMonthly, _ := (latestLog).(*deremsmodels.CCDataLogCalculatedMonthly)
		latestAccumulatedInfo.Timestamps = int(latestLogMonthly.LatestLogDate.Unix())
		latestAccumulatedInfo.LoadConsumedLifetimeEnergyACDiff = latestLogMonthly.LoadConsumedLifetimeEnergyACDiff.Float32
		latestAccumulatedInfo.PvProducedLifetimeEnergyACDiff = latestLogMonthly.PvProducedLifetimeEnergyACDiff.Float32
		latestAccumulatedInfo.BatteryLifetimeEnergyACDiff = latestLogMonthly.BatteryLifetimeEnergyACDiff.Float32
		latestAccumulatedInfo.GridLifetimeEnergyACDiff = latestLogMonthly.GridLifetimeEnergyACDiff.Float32
		latestAccumulatedInfo.LoadSelfConsumedEnergyPercentAC = latestLogMonthly.LoadSelfConsumedEnergyPercentAC.Float32
	}
	return
}

func (s defaultDevicesService) getLatestComputedDemandState(gwUUID string, startTimeIndex, endTimeIndex, endTime time.Time) (latestComputedDemandState *LatestComputedDemandState) {
	latestComputedDemandState = &LatestComputedDemandState{}
	firstLog, err1 := s.repo.CCData.GetFirstLog(gwUUID, startTimeIndex, endTimeIndex)
	latestLog, err2 := s.repo.CCData.GetLatestLog(gwUUID, startTimeIndex, endTimeIndex)
	if err1 != nil || err2 != nil {
		log.WithFields(log.Fields{
			"caused-by":      "s.repo.CCData.GetFirstLog and GetLatestLog",
			"err1":           err1,
			"err2":           err2,
			"startTimeIndex": startTimeIndex,
			"endTimeIndex":   endTimeIndex,
		}).Error()
		latestComputedDemandState.Timestamps = int(endTimeIndex.Add(-1 * time.Second).Unix())
		return
	}

	latestComputedDemandState.Timestamps = int(latestLog.LogDate.Unix())
	latestComputedDemandState.GridLifetimeEnergyACDiffToPower = utils.Division(
		utils.Diff(latestLog.GridLifetimeEnergyAC.Float32,
			firstLog.GridLifetimeEnergyAC.Float32), (15.0 / 60.0))
	latestComputedDemandState.GridContractPowerAC = utils.ThreeDecimalPlaces(latestLog.GridContractPowerAC.Float32)
	return
}

func (s defaultDevicesService) getBatteryInfo(gwUUID string, batteryEnergyInfo *BatteryEnergyInfoResponse) {
	// XXX: Hardcode battery information by gateway UUID
	const (
		Huayu      = "0324DE7B51B262F3B11A643CBA8E12CE"
		Serenegray = "0E0BA27A8175AF978C49396BDE9D7A1E"
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
		batteryEnergyInfo.Voltage = 51.2
	}
	batteryEnergyInfo.PowerSources = "Solar + Grid"
	return
}

func (s defaultDevicesService) getOnPeakTime(gwUUID string, t time.Time) (onPeakTime map[string]string, err error) {
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
	if err == nil && len(billings) == 0 {
		err = e.ErrNewBillingsNotExist
	}
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
