package services

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
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
	PreUbiikCost                     int
	PostUbiikCost                    int
}

// LatestComputedDemandState godoc
type LatestComputedDemandState struct {
	Timestamps                      int
	GridLifetimeEnergyACDiffToPower float32
	GridContractPowerAC             float32
}

// RealtimeInfo godoc
type RealtimeInfo struct {
	Timestamps                     []int
	LoadAveragePowerACs            []float32
	BatteryAveragePowerACs         []float32
	PvAveragePowerACs              []float32
	GridAveragePowerACs            []float32
	BatterySoCs                    []float32
	BatteryVoltages                []float32
	LoadPvConsumedEnergyPercentACs []float32
}

// AccumulatedInfo godoc
type AccumulatedInfo struct {
	Timestamps                        []int
	LoadConsumedLifetimeEnergyACDiffs []float32
	PvProducedLifetimeEnergyACDiffs   []float32
	BatteryLifetimeEnergyACDiffs      []float32
	GridLifetimeEnergyACDiffs         []float32
	LoadSelfConsumedEnergyPercentACs  []float32
	PreUbiikCosts                     []int
	PostUbiikCosts                    []int
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

// ComputedValues godoc
func (r *EnergyDistributionInfoResponse) ComputedValues(firstLog, latestLog *deremsmodels.CCDataLog) {
	r.PvProducedLifetimeEnergyACDiff = utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstLog.PvProducedLifetimeEnergyAC.Float32)
	r.GridProducedLifetimeEnergyACDiff = utils.Diff(latestLog.GridProducedLifetimeEnergyAC.Float32, firstLog.GridProducedLifetimeEnergyAC.Float32)
	r.BatteryProducedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryProducedLifetimeEnergyAC.Float32, firstLog.BatteryProducedLifetimeEnergyAC.Float32)
	// Avoid cc illegal value
	r.PvProducedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(r.PvProducedLifetimeEnergyACDiff)
	r.GridProducedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(r.GridProducedLifetimeEnergyACDiff)
	r.BatteryProducedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(r.BatteryProducedLifetimeEnergyACDiff)
	r.AllProducedLifetimeEnergyACDiff = utils.ThreeDecimalPlaces(
		r.PvProducedLifetimeEnergyACDiff +
			r.GridProducedLifetimeEnergyACDiff +
			r.BatteryProducedLifetimeEnergyACDiff)
	r.PvProducedEnergyPercentAC = utils.Percent(
		r.PvProducedLifetimeEnergyACDiff,
		r.AllProducedLifetimeEnergyACDiff)
	r.GridProducedEnergyPercentAC = utils.Percent(
		r.GridProducedLifetimeEnergyACDiff,
		r.AllProducedLifetimeEnergyACDiff)
	r.BatteryProducedEnergyPercentAC = utils.Percent(
		r.BatteryProducedLifetimeEnergyACDiff,
		r.AllProducedLifetimeEnergyACDiff)
	r.AllConsumedLifetimeEnergyACDiff = r.AllProducedLifetimeEnergyACDiff
	r.GridConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.GridConsumedLifetimeEnergyAC.Float32, firstLog.GridConsumedLifetimeEnergyAC.Float32)
	r.BatteryConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryConsumedLifetimeEnergyAC.Float32, firstLog.BatteryConsumedLifetimeEnergyAC.Float32)
	r.LoadConsumedLifetimeEnergyACDiff = utils.ThreeDecimalPlaces(
		r.AllConsumedLifetimeEnergyACDiff -
			(r.GridConsumedLifetimeEnergyACDiff + r.BatteryConsumedLifetimeEnergyACDiff))
	r.LoadConsumedEnergyPercentAC = utils.Percent(
		r.LoadConsumedLifetimeEnergyACDiff,
		r.AllConsumedLifetimeEnergyACDiff)
	r.GridConsumedEnergyPercentAC = utils.Percent(
		r.GridConsumedLifetimeEnergyACDiff,
		r.AllConsumedLifetimeEnergyACDiff)
	r.BatteryConsumedEnergyPercentAC = utils.Percent(
		r.BatteryConsumedLifetimeEnergyACDiff,
		r.AllConsumedLifetimeEnergyACDiff)
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

// TimeOfUseInfoResponse godoc
type TimeOfUseInfoResponse struct {
	EnergySources map[string]interface{} `json:"energySources"`
	TimeOfUse     map[string]interface{} `json:"timeOfUse"`
}

// SolarEnergyUsageResponse godoc
type SolarEnergyUsageResponse struct {
	Timestamps                     []int              `json:"timestamps"`
	LoadPvConsumedEnergyPercentACs Float32ArrayFormat `json:"loadPvConsumedEnergyPercentACs"`
}

// EnergyCostsInfo godoc
type EnergyCostsInfo struct {
	PreUbiikThisMonth             int `json:"preUbiikThisMonth"`
	PostUbiikThisMonth            int `json:"postUbiikThisMonth"`
	PreUbiikLastMonth             int `json:"preUbiikLastMonth"`
	PostUbiikLastMonth            int `json:"postUbiikLastMonth"`
	PreUbiikTheSameMonthLastYear  int `json:"preUbiikTheSameMonthLastYear"`
	PostUbiikTheSameMonthLastYear int `json:"postUbiikTheSameMonthLastYear"`
}

// EnergyCostsDailyInfo godoc
type EnergyCostsDailyInfo struct {
	Timestamps                    []int `json:"timestamps"`
	PreUbiikThisMonth             []int `json:"preUbiikThisMonth"`
	PostUbiikThisMonth            []int `json:"postUbiikThisMonth"`
	PreUbiikLastMonth             []int `json:"preUbiikLastMonth"`
	PostUbiikLastMonth            []int `json:"postUbiikLastMonth"`
	PreUbiikTheSameMonthLastYear  []int `json:"preUbiikTheSameMonthLastYear"`
	PostUbiikTheSameMonthLastYear []int `json:"postUbiikTheSameMonthLastYear"`
	SavedThisMonth                []int `json:"savedThisMonth"`
	SavedLastMonth                []int `json:"savedLastMonth"`
	SavedTheSameMonthLastYear     []int `json:"savedTheSameMonthLastYear"`
}

// TimeOfUseEnergyCostResponse godoc
type TimeOfUseEnergyCostResponse struct {
	EnergyCosts      EnergyCostsInfo      `json:"energyCosts"`
	EnergyDailyCosts EnergyCostsDailyInfo `json:"energyDailyCosts"`
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
	PvProducedLifetimeEnergyACDiff        float32       `json:"pvProducedLifetimeEnergyACDiff"`
	LoadPvConsumedEnergyPercentAC         float32       `json:"loadPvConsumedEnergyPercentAC"`
	LoadPvConsumedLifetimeEnergyACDiff    float32       `json:"loadPvConsumedLifetimeEnergyACDiff"`
	BatteryPvConsumedEnergyPercentAC      float32       `json:"batteryPvConsumedEnergyPercentAC"`
	BatteryPvConsumedLifetimeEnergyACDiff float32       `json:"batteryPvConsumedLifetimeEnergyACDiff"`
	GridPvConsumedEnergyPercentAC         float32       `json:"gridPvConsumedEnergyPercentAC"`
	GridPvConsumedLifetimeEnergyACDiff    float32       `json:"gridPvConsumedLifetimeEnergyACDiff"`
	PvEnergyCostSavingsSum                int           `json:"pvEnergyCostSavingsSum"`
	PvCo2SavingsSum                       Float32Format `json:"pvCo2SavingsSum"`
}

// ComputedValues godoc
func (r *SolarEnergyInfoResponse) ComputedValues(firstLogOfDay, firstLogOfMonth, latestLog *deremsmodels.CCDataLog, logsOfMonth []*deremsmodels.CCDataLog) {
	r.PvProducedLifetimeEnergyACDiff = utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstLogOfDay.PvProducedLifetimeEnergyAC.Float32)
	r.LoadPvConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.LoadPvConsumedLifetimeEnergyAC.Float32, firstLogOfDay.LoadPvConsumedLifetimeEnergyAC.Float32)
	r.BatteryPvConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryPvConsumedLifetimeEnergyAC.Float32, firstLogOfDay.BatteryPvConsumedLifetimeEnergyAC.Float32)
	r.GridPvConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.GridPvConsumedLifetimeEnergyAC.Float32, firstLogOfDay.GridPvConsumedLifetimeEnergyAC.Float32)
	// Avoid cc illegal value
	r.PvProducedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(r.PvProducedLifetimeEnergyACDiff)
	r.LoadPvConsumedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(r.LoadPvConsumedLifetimeEnergyACDiff)
	r.BatteryPvConsumedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(r.BatteryPvConsumedLifetimeEnergyACDiff)
	r.GridPvConsumedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(r.GridPvConsumedLifetimeEnergyACDiff)
	// Percent and value are recomputed by PvProducedLifetimeEnergyACDiff
	sumOfPvConsumedLifetimeEnergyAC := r.LoadPvConsumedLifetimeEnergyACDiff + r.BatteryPvConsumedLifetimeEnergyACDiff + r.GridPvConsumedLifetimeEnergyACDiff
	if sumOfPvConsumedLifetimeEnergyAC == 0 {
		r.PvProducedLifetimeEnergyACDiff = 0
	} else {
		r.LoadPvConsumedLifetimeEnergyACDiff = utils.ThreeDecimalPlaces(r.PvProducedLifetimeEnergyACDiff * utils.Division(r.LoadPvConsumedLifetimeEnergyACDiff, sumOfPvConsumedLifetimeEnergyAC))
		r.BatteryPvConsumedLifetimeEnergyACDiff = utils.ThreeDecimalPlaces(r.PvProducedLifetimeEnergyACDiff * utils.Division(r.BatteryPvConsumedLifetimeEnergyACDiff, sumOfPvConsumedLifetimeEnergyAC))
		r.GridPvConsumedLifetimeEnergyACDiff = utils.ThreeDecimalPlaces(r.PvProducedLifetimeEnergyACDiff * utils.Division(r.GridPvConsumedLifetimeEnergyACDiff, sumOfPvConsumedLifetimeEnergyAC))
	}
	r.LoadPvConsumedEnergyPercentAC = utils.Percent(
		r.LoadPvConsumedLifetimeEnergyACDiff,
		r.PvProducedLifetimeEnergyACDiff)
	r.BatteryPvConsumedEnergyPercentAC = utils.Percent(
		r.BatteryPvConsumedLifetimeEnergyACDiff,
		r.PvProducedLifetimeEnergyACDiff)
	r.GridPvConsumedEnergyPercentAC = utils.Percent(
		r.GridPvConsumedLifetimeEnergyACDiff,
		r.PvProducedLifetimeEnergyACDiff)

	sumOfPvEnergyCostSavings := utils.Diff(latestLog.PvEnergyCostSavings.Float32, firstLogOfMonth.PvEnergyCostSavings.Float32)
	sumOfPvCo2Savings := utils.Diff(latestLog.PvCo2Savings.Float32, firstLogOfMonth.PvCo2Savings.Float32)
	r.PvEnergyCostSavingsSum = int(sumOfPvEnergyCostSavings)
	r.PvCo2SavingsSum = Float32Format(sumOfPvCo2Savings)
}

// SolarPowerStateResponse godoc
type SolarPowerStateResponse struct {
	Timestamps        []int                  `json:"timestamps"`
	PvAveragePowerACs Float32ArrayFormat     `json:"pvAveragePowerACs"`
	TimeOfUse         map[string]interface{} `json:"timeOfUse"`
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

// GetBatteryInfo godoc
func (r *BatteryEnergyInfoResponse) GetBatteryInfo(gwUUID string) {
	// XXX: Hardcode battery information by gateway UUID
	const (
		Huayu      = "0324DE7B51B262F3B11A643CBA8E12CE"
		Serenegray = "0E0BA27A8175AF978C49396BDE9D7A1E"
		CHTMiaoli  = "018F1623ADD8E739F7C6CBE62A7DF3C0"
	)
	switch gwUUID {
	case Huayu:
		r.Model = "PR2116 Poweroad Battery"
		r.Capcity = 48
		r.BatteryPower = 30
		r.Voltage = 480
	case Serenegray:
		r.Model = "L051100-A UZ-Energy Battery"
		r.Capcity = 30
		r.BatteryPower = 20
		r.Voltage = 51.2
	case CHTMiaoli:
		r.Model = "L051100-A UZ-Energy Battery"
		r.Capcity = 10
		r.BatteryPower = 5
		r.Voltage = 51.2
	}
	r.PowerSources = "Solar + Grid"
}

// ComputedValues godoc
func (r *BatteryEnergyInfoResponse) ComputedValues(firstLog, latestLog *deremsmodels.CCDataLog) {
	r.BatteryLifetimeOperationCyclesDiff = utils.Diff(latestLog.BatteryLifetimeOperationCycles.Float32, firstLog.BatteryLifetimeOperationCycles.Float32)
	r.BatteryLifetimeOperationCycles = Float32Format(latestLog.BatteryLifetimeOperationCycles.Float32)
	r.BatterySoC = Float32Format(latestLog.BatterySoC.Float32)
	r.BatteryProducedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryProducedLifetimeEnergyAC.Float32, firstLog.BatteryProducedLifetimeEnergyAC.Float32)
	r.BatteryProducedLifetimeEnergyAC = Float32Format(latestLog.BatteryProducedLifetimeEnergyAC.Float32)
	r.BatteryConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.BatteryConsumedLifetimeEnergyAC.Float32, firstLog.BatteryConsumedLifetimeEnergyAC.Float32)
	r.BatteryConsumedLifetimeEnergyAC = Float32Format(latestLog.BatteryConsumedLifetimeEnergyAC.Float32)
}

// BatteryPowerStateResponse godoc
type BatteryPowerStateResponse struct {
	Timestamps             []int                  `json:"timestamps"`
	BatteryAveragePowerACs Float32ArrayFormat     `json:"batteryAveragePowerACs"`
	TimeOfUse              map[string]interface{} `json:"timeOfUse"`
}

// BatteryChargeVoltageStateResponse godoc
type BatteryChargeVoltageStateResponse struct {
	Timestamps      []int                  `json:"timestamps"`
	BatterySoCs     Float32ArrayFormat     `json:"batterySoCs"`
	BatteryVoltages Float32ArrayFormat     `json:"batteryVoltages"`
	TimeOfUse       map[string]interface{} `json:"timeOfUse"`
}

// GridEnergyInfoResponse godoc
type GridEnergyInfoResponse struct {
	GridConsumedLifetimeEnergyACDiff float32 `json:"gridConsumedLifetimeEnergyACDiff"`
	GridProducedLifetimeEnergyACDiff float32 `json:"gridProducedLifetimeEnergyACDiff"`
	GridLifetimeEnergyACDiff         float32 `json:"gridLifetimeEnergyACDiff"`
	GridLifetimeEnergyACDiffOfMonth  float32 `json:"gridLifetimeEnergyACDiffOfMonth"`
}

// ComputedValues godoc
func (r *GridEnergyInfoResponse) ComputedValues(firstLogOfDay, firstLogOfMonth, latestLog *deremsmodels.CCDataLog) {
	r.GridConsumedLifetimeEnergyACDiff = utils.Diff(latestLog.GridConsumedLifetimeEnergyAC.Float32, firstLogOfDay.GridConsumedLifetimeEnergyAC.Float32)
	r.GridProducedLifetimeEnergyACDiff = utils.Diff(latestLog.GridProducedLifetimeEnergyAC.Float32, firstLogOfDay.GridProducedLifetimeEnergyAC.Float32)
	r.GridLifetimeEnergyACDiff = utils.Diff(latestLog.GridLifetimeEnergyAC.Float32, firstLogOfDay.GridLifetimeEnergyAC.Float32)
	r.GridLifetimeEnergyACDiffOfMonth = utils.Diff(latestLog.GridLifetimeEnergyAC.Float32, firstLogOfMonth.GridLifetimeEnergyAC.Float32)
}

// GridPowerStateResponse godoc
type GridPowerStateResponse struct {
	Timestamps          []int                  `json:"timestamps"`
	GridAveragePowerACs Float32ArrayFormat     `json:"gridAveragePowerACs"`
	TimeOfUse           map[string]interface{} `json:"timeOfUse"`
}

// DevicesService godoc
type DevicesService interface {
	GetLatestDevicesEnergyInfo(gwUUID string) (updatedTime time.Time, devicesEnergyInfo *DevicesEnergyInfoResponse, err error)
	GetEnergyDistributionInfo(param *app.PeriodParam) (energyDistributionInfo *EnergyDistributionInfoResponse)
	GetPowerState(param *app.ZoomableParam) (powerState *PowerStateResponse)
	GetAccumulatedPowerState(param *app.ResolutionWithPeriodParam) (accumulatedPowerState *AccumulatedPowerStateResponse)
	GetPowerSelfSupplyRate(param *app.ResolutionWithPeriodParam) (powerSelfSupplyRate *PowerSelfSupplyRateResponse)
	GetBatteryUsageInfo(param *app.StartTimeParam) (batteryUsageInfo *BatteryUsageInfoResponse)
	GetTimeOfUseInfo(param *app.StartTimeParam) (timeOfUseInfo *TimeOfUseInfoResponse, err error)
	GetSolarEnergyUsage(param *app.ZoomableParam) (solarEnergyUsage *SolarEnergyUsageResponse)
	GetTimeOfUseEnergyCost(param *app.PeriodParam) (timeOfUseEnergyCost *TimeOfUseEnergyCostResponse)
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
		// To avoid frontend crash, return default values if no data in database
		err = nil
		latestLog = &deremsmodels.CCDataLog{}
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
	if err1 := json.Unmarshal(latestLog.LoadLinks.JSON, &devicesEnergyInfo.LoadLinks); err1 != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal: loadLinksValue",
			"err1":      err1,
		}).Error()
	}
	if devicesEnergyInfo.LoadLinks == nil {
		devicesEnergyInfo.LoadLinks = map[string]interface{}{
			"pv":      0,
			"grid":    0,
			"battery": 0,
		}
	}
	if err1 := json.Unmarshal(latestLog.GridLinks.JSON, &devicesEnergyInfo.GridLinks); err1 != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal: gridLinksValue",
			"err1":      err1,
		}).Error()
	}
	if devicesEnergyInfo.GridLinks == nil {
		devicesEnergyInfo.GridLinks = map[string]interface{}{
			"pv":      0,
			"load":    0,
			"battery": 0,
		}
	}
	if err1 := json.Unmarshal(latestLog.PvLinks.JSON, &devicesEnergyInfo.PVLinks); err1 != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal: pvLinksValue",
			"err1":      err1,
		}).Error()
	}
	if devicesEnergyInfo.PVLinks == nil {
		devicesEnergyInfo.PVLinks = map[string]interface{}{
			"grid":    0,
			"load":    0,
			"battery": 0,
		}
	}
	if err1 := json.Unmarshal(latestLog.BatteryLinks.JSON, &devicesEnergyInfo.BatteryLinks); err1 != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal: batteryLinksValue",
			"err1":      err1,
		}).Error()
	}
	if devicesEnergyInfo.BatteryLinks == nil {
		devicesEnergyInfo.BatteryLinks = map[string]interface{}{
			"pv":   0,
			"grid": 0,
			"load": 0,
		}
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
	energyDistributionInfo.ComputedValues(firstLog, latestLog)
	return
}

func (s defaultDevicesService) GetPowerState(param *app.ZoomableParam) (powerState *PowerStateResponse) {
	realtimeInfo := s.getRealtimeInfo(param, false)
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

func (s defaultDevicesService) GetTimeOfUseInfo(param *app.StartTimeParam) (timeOfUseInfo *TimeOfUseInfoResponse, err error) {
	timeOfUseInfo = &TimeOfUseInfoResponse{}
	localStartTime, tous, err := s.billing.GetTOUsOfLocalTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	// 1. energySources
	energySources, err := s.getEnergySourcesInfo(param.GatewayUUID, localStartTime, tous)
	if err != nil {
		return
	}
	timeOfUseInfo.EnergySources = energySources

	// 2. timeOfUse
	timeOfUse, err := s.getTimeOfUseOfDay(localStartTime, tous)
	if err != nil {
		return
	}
	timeOfUseInfo.TimeOfUse = timeOfUse

	log.WithFields(log.Fields{
		"energySources": energySources,
		"timeOfUse":     timeOfUse,
	}).Debug()
	return
}

func (s defaultDevicesService) getEnergySourcesInfo(gwUUID string, localStartTime time.Time, tous []*deremsmodels.Tou) (energySources map[string]interface{}, err error) {
	energySources = make(map[string]interface{})

	onPeak, err := s.getEnergySourceDistributionByPeakType("On-peak", gwUUID, localStartTime, tous)
	if err != nil {
		return
	}
	midPeak, err := s.getEnergySourceDistributionByPeakType("Mid-peak", gwUUID, localStartTime, tous)
	if err != nil {
		return
	}
	offPeak, err := s.getEnergySourceDistributionByPeakType("Off-peak", gwUUID, localStartTime, tous)
	if err != nil {
		return
	}
	if len(onPeak) > 0 {
		energySources["onPeak"] = onPeak
	}
	if len(midPeak) > 0 {
		energySources["midPeak"] = midPeak
	}
	if len(offPeak) > 0 {
		energySources["offPeak"] = offPeak
	}
	return
}

func (s defaultDevicesService) getEnergySourceDistributionByPeakType(peakType, gwUUID string, localStartTime time.Time, tous []*deremsmodels.Tou) (energySourceDistribution map[string]float32, err error) {
	energySourceDistribution = make(map[string]float32)

	loc := time.FixedZone(localStartTime.Zone())
	for _, tou := range tous {
		if tou.PeakType.String != peakType {
			continue
		}

		startTime, err := time.ParseInLocation(utils.HHMMSS24h, tou.PeriodStime.String, loc)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "time.ParseInLocation",
				"err":       err,
			}).Error()
			break
		}
		startTimeInUTC := time.Date(localStartTime.Year(), localStartTime.Month(), localStartTime.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, loc).In(time.UTC)
		endTime, err := time.ParseInLocation(utils.HHMMSS24h, tou.PeriodEtime.String, loc)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "time.ParseInLocation",
				"err":       err,
			}).Error()
			break
		}
		endTimeInUTC := time.Date(localStartTime.Year(), localStartTime.Month(), localStartTime.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, loc).In(time.UTC)
		if tou.PeriodEtime.String == "00:00:00" {
			endTimeInUTC = endTimeInUTC.AddDate(0, 0, 1)
		}
		firstLog, err1 := s.repo.CCData.GetFirstLog(gwUUID, startTimeInUTC, endTimeInUTC)
		latestLog, err2 := s.repo.CCData.GetLatestLog(gwUUID, startTimeInUTC, endTimeInUTC)
		if err1 != nil || err2 != nil {
			log.WithFields(log.Fields{
				"caused-by":      "s.repo.CCData.GetFirstLog and GetLatestLog",
				"err1":           err1,
				"err2":           err2,
				"startTimeInUTC": startTimeInUTC,
				"endTimeInUTC":   endTimeInUTC,
			}).Error()
			continue
		}

		pvProducedLifetimeEnergyACDiff := utils.Diff(latestLog.PvProducedLifetimeEnergyAC.Float32, firstLog.PvProducedLifetimeEnergyAC.Float32)
		gridProducedLifetimeEnergyACDiff := utils.Diff(latestLog.GridProducedLifetimeEnergyAC.Float32, firstLog.GridProducedLifetimeEnergyAC.Float32)
		batteryProducedLifetimeEnergyACDiff := utils.Diff(latestLog.BatteryProducedLifetimeEnergyAC.Float32, firstLog.BatteryProducedLifetimeEnergyAC.Float32)
		// Avoid cc illegal value
		pvProducedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(pvProducedLifetimeEnergyACDiff)
		gridProducedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(gridProducedLifetimeEnergyACDiff)
		batteryProducedLifetimeEnergyACDiff = utils.GetZeroForNegativeValue(batteryProducedLifetimeEnergyACDiff)
		energySourceDistribution["pvProducedLifetimeEnergyACDiff"] += pvProducedLifetimeEnergyACDiff
		energySourceDistribution["gridProducedLifetimeEnergyACDiff"] += gridProducedLifetimeEnergyACDiff
		energySourceDistribution["batteryProducedLifetimeEnergyACDiff"] += batteryProducedLifetimeEnergyACDiff
	}

	if len(energySourceDistribution) > 0 {
		energySourceDistribution["pvProducedLifetimeEnergyACDiff"] = utils.ThreeDecimalPlaces(energySourceDistribution["pvProducedLifetimeEnergyACDiff"])
		energySourceDistribution["gridProducedLifetimeEnergyACDiff"] = utils.ThreeDecimalPlaces(energySourceDistribution["gridProducedLifetimeEnergyACDiff"])
		energySourceDistribution["batteryProducedLifetimeEnergyACDiff"] = utils.ThreeDecimalPlaces(energySourceDistribution["batteryProducedLifetimeEnergyACDiff"])
		energySourceDistribution["allProducedLifetimeEnergyACDiff"] = utils.ThreeDecimalPlaces(
			energySourceDistribution["pvProducedLifetimeEnergyACDiff"] +
				energySourceDistribution["gridProducedLifetimeEnergyACDiff"] +
				energySourceDistribution["batteryProducedLifetimeEnergyACDiff"])
		energySourceDistribution["pvProducedEnergyPercentAC"] = utils.Percent(
			energySourceDistribution["pvProducedLifetimeEnergyACDiff"],
			energySourceDistribution["allProducedLifetimeEnergyACDiff"])
		energySourceDistribution["gridProducedEnergyPercentAC"] = utils.Percent(
			energySourceDistribution["gridProducedLifetimeEnergyACDiff"],
			energySourceDistribution["allProducedLifetimeEnergyACDiff"])
		energySourceDistribution["batteryProducedEnergyPercentAC"] = utils.Percent(
			energySourceDistribution["batteryProducedLifetimeEnergyACDiff"],
			energySourceDistribution["allProducedLifetimeEnergyACDiff"])
	}
	return
}

func (s defaultDevicesService) getTimeOfUseOfDay(localStartTime time.Time, tous []*deremsmodels.Tou) (timeOfUse map[string]interface{}, err error) {
	timeOfUse = make(map[string]interface{})

	// 1. timezone
	timeOfUse["timezone"] = localStartTime.Format(utils.ZHHMM)

	// 2. onPeak, midPeak, offPeak
	timeOfUse["onPeak"] = s.getPeriodsByPeakType("On-peak", tous)
	timeOfUse["midPeak"] = s.getPeriodsByPeakType("Mid-peak", tous)
	timeOfUse["offPeak"] = s.getPeriodsByPeakType("Off-peak", tous)

	// 3. currentPeakType
	loc := time.FixedZone(localStartTime.Zone())
	peakType, err := s.billing.GetPeakType(time.Now().In(loc), tous)
	if err != nil {
		return
	}
	switch peakType {
	case "Off-peak":
		timeOfUse["currentPeakType"] = "offPeak"
	case "On-peak":
		timeOfUse["currentPeakType"] = "onPeak"
	case "Mid-peak":
		timeOfUse["currentPeakType"] = "midPeak"
	}
	return
}

func (s defaultDevicesService) getPeriodsByPeakType(peakType string, tous []*deremsmodels.Tou) (periods []map[string]interface{}) {
	for _, tou := range tous {
		if tou.PeakType.String != peakType {
			continue
		}

		period := map[string]interface{}{
			"start":   tou.PeriodStime.String,
			"end":     tou.PeriodEtime.String,
			"touRate": tou.FlowRate.Float32,
		}
		if period["end"] == "00:00:00" {
			period["end"] = "24:00:00"
		}
		periods = append(periods, period)
	}
	return
}

func (s defaultDevicesService) GetSolarEnergyUsage(param *app.ZoomableParam) (solarEnergyUsage *SolarEnergyUsageResponse) {
	realtimeInfo := s.getRealtimeInfo(param, true)
	solarEnergyUsage = &SolarEnergyUsageResponse{
		Timestamps:                     realtimeInfo.Timestamps,
		LoadPvConsumedEnergyPercentACs: realtimeInfo.LoadPvConsumedEnergyPercentACs,
	}
	return
}

func (s defaultDevicesService) GetTimeOfUseEnergyCost(param *app.PeriodParam) (timeOfUseEnergyCost *TimeOfUseEnergyCostResponse) {
	timeOfUseEnergyCost = &TimeOfUseEnergyCostResponse{}

	accumulatedParam := &app.ResolutionWithPeriodParam{
		GatewayUUID: param.GatewayUUID,
		Query: app.ResolutionWithPeriodQuery{
			Resolution: "day",
			StartTime:  param.Query.StartTime,
			EndTime:    param.Query.EndTime,
		},
	}
	// This month
	log.WithFields(log.Fields{
		"This month - StartTime": accumulatedParam.Query.StartTime,
		"This month - EndTime":   accumulatedParam.Query.EndTime,
	}).Debug()
	accumulatedInfoThisMonth := s.getAccumulatedInfo(accumulatedParam)
	// Last month
	accumulatedParam.Query.StartTime = utils.AddDate(param.Query.StartTime, 0, -1, 0)
	accumulatedParam.Query.EndTime = utils.AddDate(param.Query.EndTime, 0, -1, 0)
	log.WithFields(log.Fields{
		"Last month - StartTime": accumulatedParam.Query.StartTime,
		"Last month - EndTime":   accumulatedParam.Query.EndTime,
	}).Debug()
	accumulatedInfoLastMonth := s.getAccumulatedInfo(accumulatedParam)
	// The same month last year
	accumulatedParam.Query.StartTime = utils.AddDate(param.Query.StartTime, -1, 0, 0)
	accumulatedParam.Query.EndTime = utils.AddDate(param.Query.EndTime, -1, 0, 0)
	log.WithFields(log.Fields{
		"The same month last year - StartTime": accumulatedParam.Query.StartTime,
		"The same month last year - EndTime":   accumulatedParam.Query.EndTime,
	}).Debug()
	accumulatedInfoTheSameMonthLastYear := s.getAccumulatedInfo(accumulatedParam)

	energyCostsInfo := EnergyCostsInfo{
		PreUbiikThisMonth:             utils.SumOfArray(accumulatedInfoThisMonth.PreUbiikCosts),
		PostUbiikThisMonth:            utils.SumOfArray(accumulatedInfoThisMonth.PostUbiikCosts),
		PreUbiikLastMonth:             utils.SumOfArray(accumulatedInfoLastMonth.PreUbiikCosts),
		PostUbiikLastMonth:            utils.SumOfArray(accumulatedInfoLastMonth.PostUbiikCosts),
		PreUbiikTheSameMonthLastYear:  utils.SumOfArray(accumulatedInfoTheSameMonthLastYear.PreUbiikCosts),
		PostUbiikTheSameMonthLastYear: utils.SumOfArray(accumulatedInfoTheSameMonthLastYear.PostUbiikCosts),
	}
	timeOfUseEnergyCost.EnergyCosts = energyCostsInfo
	energyCostsDailyInfo := EnergyCostsDailyInfo{
		Timestamps:                    accumulatedInfoThisMonth.Timestamps,
		PreUbiikThisMonth:             accumulatedInfoThisMonth.PreUbiikCosts,
		PostUbiikThisMonth:            accumulatedInfoThisMonth.PostUbiikCosts,
		PreUbiikLastMonth:             accumulatedInfoLastMonth.PreUbiikCosts,
		PostUbiikLastMonth:            accumulatedInfoLastMonth.PostUbiikCosts,
		PreUbiikTheSameMonthLastYear:  accumulatedInfoTheSameMonthLastYear.PreUbiikCosts,
		PostUbiikTheSameMonthLastYear: accumulatedInfoTheSameMonthLastYear.PostUbiikCosts,
		SavedThisMonth:                utils.DiffTwoArrays(accumulatedInfoThisMonth.PreUbiikCosts, accumulatedInfoThisMonth.PostUbiikCosts),
		SavedLastMonth:                utils.DiffTwoArrays(accumulatedInfoLastMonth.PreUbiikCosts, accumulatedInfoLastMonth.PostUbiikCosts),
		SavedTheSameMonthLastYear:     utils.DiffTwoArrays(accumulatedInfoTheSameMonthLastYear.PreUbiikCosts, accumulatedInfoTheSameMonthLastYear.PostUbiikCosts),
	}
	timeOfUseEnergyCost.EnergyDailyCosts = energyCostsDailyInfo
	return
}

func (s defaultDevicesService) GetDemandState(param *app.PeriodParam) (demandState *DemandStateResponse) {
	demandState = &DemandStateResponse{}
	startTimeIndex := param.Query.StartTime
	endTimeIndex := param.Query.StartTime.Add(15 * time.Minute)
	if endTimeIndex.After(param.Query.EndTime) {
		endTimeIndex = param.Query.EndTime
	}

	for startTimeIndex.Before(param.Query.EndTime) {
		latestComputedDemandState := s.getLatestComputedDemandState(param.GatewayUUID, startTimeIndex, endTimeIndex, param.Query.EndTime)
		if latestComputedDemandState != nil {
			log.Debug("latestComputedDemandState: ", latestComputedDemandState)
			demandState.Timestamps = append(demandState.Timestamps, latestComputedDemandState.Timestamps)
			demandState.GridLifetimeEnergyACDiffToPowers = append(demandState.GridLifetimeEnergyACDiffToPowers, latestComputedDemandState.GridLifetimeEnergyACDiffToPower)
			if latestComputedDemandState.GridContractPowerAC != 0 {
				demandState.GridContractPowerAC = latestComputedDemandState.GridContractPowerAC
			}
		}

		startTimeIndex = endTimeIndex
		endTimeIndex = startTimeIndex.Add(15 * time.Minute)
		if endTimeIndex.After(param.Query.EndTime) {
			endTimeIndex = param.Query.EndTime
		}
	}

	// avoid frontend error to send default line data for chart display
	if demandState.Timestamps == nil {
		startTimeIndex = param.Query.StartTime
		endTimeIndex = param.Query.StartTime.Add(15 * time.Minute)
		if endTimeIndex.After(param.Query.EndTime) {
			endTimeIndex = param.Query.EndTime
		}
		for startTimeIndex.Before(param.Query.EndTime) {
			demandState.Timestamps = append(demandState.Timestamps, int(endTimeIndex.Add(-1*time.Second).Unix()))
			demandState.GridLifetimeEnergyACDiffToPowers = append(demandState.GridLifetimeEnergyACDiffToPowers, 0)

			startTimeIndex = endTimeIndex
			endTimeIndex = startTimeIndex.Add(15 * time.Minute)
			if endTimeIndex.After(param.Query.EndTime) {
				endTimeIndex = param.Query.EndTime
			}
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
	solarEnergyInfo.ComputedValues(firstLogOfDay, firstLogOfMonth, latestLog, logsOfMonth)
	return
}

func (s defaultDevicesService) GetSolarPowerState(param *app.ZoomableParam) (solarPowerState *SolarPowerStateResponse, err error) {
	solarPowerState = &SolarPowerStateResponse{}
	timeOfUseOnPeakTime, err := s.getTimeOfUseOnPeakTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	solarPowerState.TimeOfUse = timeOfUseOnPeakTime
	realtimeInfo := s.getRealtimeInfo(param, false)
	solarPowerState.Timestamps = realtimeInfo.Timestamps
	solarPowerState.PvAveragePowerACs = realtimeInfo.PvAveragePowerACs
	return
}

func (s defaultDevicesService) GetBatteryEnergyInfo(param *app.StartTimeParam) (batteryEnergyInfo *BatteryEnergyInfoResponse) {
	batteryEnergyInfo = &BatteryEnergyInfoResponse{}
	batteryEnergyInfo.GetBatteryInfo(param.GatewayUUID)
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
	batteryEnergyInfo.ComputedValues(firstLog, latestLog)
	return
}

func (s defaultDevicesService) GetBatteryPowerState(param *app.ZoomableParam) (batteryPowerState *BatteryPowerStateResponse, err error) {
	batteryPowerState = &BatteryPowerStateResponse{}
	timeOfUseOnPeakTime, err := s.getTimeOfUseOnPeakTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	batteryPowerState.TimeOfUse = timeOfUseOnPeakTime
	realtimeInfo := s.getRealtimeInfo(param, false)
	batteryPowerState.Timestamps = realtimeInfo.Timestamps
	batteryPowerState.BatteryAveragePowerACs = realtimeInfo.BatteryAveragePowerACs
	return
}

func (s defaultDevicesService) GetBatteryChargeVoltageState(param *app.ZoomableParam) (batteryChargeVoltageState *BatteryChargeVoltageStateResponse, err error) {
	batteryChargeVoltageState = &BatteryChargeVoltageStateResponse{}
	timeOfUseOnPeakTime, err := s.getTimeOfUseOnPeakTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	batteryChargeVoltageState.TimeOfUse = timeOfUseOnPeakTime
	realtimeInfo := s.getRealtimeInfo(param, false)
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
	gridEnergyInfo.ComputedValues(firstLogOfDay, firstLogOfMonth, latestLog)
	return
}

func (s defaultDevicesService) GetGridPowerState(param *app.ZoomableParam) (gridPowerState *GridPowerStateResponse, err error) {
	gridPowerState = &GridPowerStateResponse{}
	timeOfUseOnPeakTime, err := s.getTimeOfUseOnPeakTime(param.GatewayUUID, param.Query.StartTime)
	if err != nil {
		return
	}

	gridPowerState.TimeOfUse = timeOfUseOnPeakTime
	realtimeInfo := s.getRealtimeInfo(param, false)
	gridPowerState.Timestamps = realtimeInfo.Timestamps
	gridPowerState.GridAveragePowerACs = realtimeInfo.GridAveragePowerACs
	return
}

func (s defaultDevicesService) getRealtimeInfo(param *app.ZoomableParam, includedComputedData bool) (realtimeInfo *RealtimeInfo) {
	realtimeInfo = &RealtimeInfo{}
	startTimeIndex := param.Query.StartTime
	endTimeIndex := param.GetEndTimeIndex()
	if endTimeIndex.After(param.Query.EndTime) {
		endTimeIndex = param.Query.EndTime
	}

	for startTimeIndex.Before(param.Query.EndTime) {
		latestRealtimeInfo := s.getLatestRealtimeInfo(param.GatewayUUID, startTimeIndex, endTimeIndex, param.Query.EndTime)
		if latestRealtimeInfo != nil {
			log.Debug("latestRealtimeInfo.LogDate: ", latestRealtimeInfo.LogDate)
			realtimeInfo.Timestamps = append(realtimeInfo.Timestamps, int(latestRealtimeInfo.LogDate.Unix()))
			realtimeInfo.LoadAveragePowerACs = append(realtimeInfo.LoadAveragePowerACs, utils.ThreeDecimalPlaces(latestRealtimeInfo.LoadAveragePowerAC.Float32))
			realtimeInfo.BatteryAveragePowerACs = append(realtimeInfo.BatteryAveragePowerACs, utils.ThreeDecimalPlaces(latestRealtimeInfo.BatteryAveragePowerAC.Float32))
			realtimeInfo.PvAveragePowerACs = append(realtimeInfo.PvAveragePowerACs, utils.ThreeDecimalPlaces(latestRealtimeInfo.PvAveragePowerAC.Float32))
			realtimeInfo.GridAveragePowerACs = append(realtimeInfo.GridAveragePowerACs, utils.ThreeDecimalPlaces(latestRealtimeInfo.GridAveragePowerAC.Float32))
			realtimeInfo.BatterySoCs = append(realtimeInfo.BatterySoCs, utils.ThreeDecimalPlaces(latestRealtimeInfo.BatterySoC.Float32))
			realtimeInfo.BatteryVoltages = append(realtimeInfo.BatteryVoltages, utils.ThreeDecimalPlaces(latestRealtimeInfo.BatteryVoltage.Float32))

			if includedComputedData {
				firstRealtimeInfo := s.getFirstLogRealtimeInfo(param.GatewayUUID, startTimeIndex, endTimeIndex, param.Query.EndTime)
				if firstRealtimeInfo != nil {
					loadPvConsumedEnergyPercentAC := s.computeLoadPvConsumedEnergyPercentACValue(latestRealtimeInfo)
					realtimeInfo.LoadPvConsumedEnergyPercentACs = append(realtimeInfo.LoadPvConsumedEnergyPercentACs, loadPvConsumedEnergyPercentAC)
				}
			}
		}

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

	// avoid frontend error to send default line data for chart display
	if realtimeInfo.Timestamps == nil {
		startTimeIndex = param.Query.StartTime
		endTimeIndex = param.GetEndTimeIndex()
		if endTimeIndex.After(param.Query.EndTime) {
			endTimeIndex = param.Query.EndTime
		}
		for startTimeIndex.Before(param.Query.EndTime) {
			realtimeInfo.Timestamps = append(realtimeInfo.Timestamps, int(endTimeIndex.Add(-1*time.Second).Unix()))
			realtimeInfo.LoadAveragePowerACs = append(realtimeInfo.LoadAveragePowerACs, 0)
			realtimeInfo.BatteryAveragePowerACs = append(realtimeInfo.BatteryAveragePowerACs, 0)
			realtimeInfo.PvAveragePowerACs = append(realtimeInfo.PvAveragePowerACs, 0)
			realtimeInfo.GridAveragePowerACs = append(realtimeInfo.GridAveragePowerACs, 0)
			realtimeInfo.BatterySoCs = append(realtimeInfo.BatterySoCs, 0)
			realtimeInfo.BatteryVoltages = append(realtimeInfo.BatteryVoltages, 0)
			realtimeInfo.LoadPvConsumedEnergyPercentACs = append(realtimeInfo.LoadPvConsumedEnergyPercentACs, 0)

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
	}
	return
}

func (s defaultDevicesService) computeLoadPvConsumedEnergyPercentACValue(latestRealtimeInfo *deremsmodels.CCDataLog) (loadPvConsumedEnergyPercentAC float32) {
	// Compute by power values
	loadPvConsumedEnergyPercentAC = utils.Percent(
		latestRealtimeInfo.LoadPvAveragePowerAC.Float32,
		latestRealtimeInfo.LoadPvAveragePowerAC.Float32+latestRealtimeInfo.BatteryPvAveragePowerAC.Float32+latestRealtimeInfo.GridPvAveragePowerAC.Float32)
	loadPvConsumedEnergyPercentAC = float32(math.Abs(float64(loadPvConsumedEnergyPercentAC)))
	return
}

func (s defaultDevicesService) getAccumulatedInfo(param *app.ResolutionWithPeriodParam) (accumulatedInfo *AccumulatedInfo) {
	accumulatedInfo = &AccumulatedInfo{}
	latestAccumulatedInfoArray := s.getLatestAccumulatedInfoArray(param.GatewayUUID, param.Query.Resolution, param.Query.StartTime, param.Query.EndTime)

	startTimeIndex := param.Query.StartTime
	endTimeIndex := param.GetEndTimeIndex()
	for startTimeIndex.Before(param.Query.EndTime) {
		latestAccumulatedInfo := &LatestAccumulatedInfo{}
		dataExists := false
		for _, info := range latestAccumulatedInfoArray {
			if info.Timestamps >= int(startTimeIndex.Unix()) && info.Timestamps < int(endTimeIndex.Unix()) {
				*latestAccumulatedInfo = *info
				dataExists = true
				break
			}
		}
		if !dataExists {
			latestAccumulatedInfo.Timestamps = int(endTimeIndex.Add(-1 * time.Second).Unix())
		}
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
			accumulatedInfo.PreUbiikCosts = append(accumulatedInfo.PreUbiikCosts, latestAccumulatedInfo.PreUbiikCost)
			accumulatedInfo.PostUbiikCosts = append(accumulatedInfo.PostUbiikCosts, latestAccumulatedInfo.PostUbiikCost)
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
		return nil
	}
	return
}

func (s defaultDevicesService) getFirstLogRealtimeInfo(gwUUID string, startTimeIndex, endTimeIndex, endTime time.Time) (firstLog *deremsmodels.CCDataLog) {
	firstLog, err := s.repo.CCData.GetFirstLog(gwUUID, startTimeIndex, endTimeIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by":      "s.repo.CCData.GetFirstLog",
			"err":            err,
			"startTimeIndex": startTimeIndex,
			"endTimeIndex":   endTimeIndex,
		}).Error()
		return nil
	}
	return
}

func (s defaultDevicesService) getLatestAccumulatedInfoArray(gwUUID, resolution string, startTime, endTime time.Time) (latestAccumulatedInfoArray []*LatestAccumulatedInfo) {
	logs, err := s.repo.CCData.GetCalculatedLogs(gwUUID, resolution, startTime, endTime)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.CCData.GetCalculatedLogs",
			"err":       err,
			"startTime": startTime,
			"endTime":   endTime,
		}).Error()
		return
	}

	if reflect.TypeOf(logs).Kind() == reflect.Slice {
		s := reflect.ValueOf(logs)
		for i := 0; i < s.Len(); i++ {
			switch resolution {
			case "day":
				latestLogDaily := (s.Index(i).Interface()).(*deremsmodels.CCDataLogCalculatedDaily)
				latestAccumulatedInfo := &LatestAccumulatedInfo{}
				latestAccumulatedInfo.Timestamps = int(latestLogDaily.LatestLogDate.Unix())
				latestAccumulatedInfo.LoadConsumedLifetimeEnergyACDiff = latestLogDaily.LoadConsumedLifetimeEnergyACDiff.Float32
				latestAccumulatedInfo.PvProducedLifetimeEnergyACDiff = latestLogDaily.PvProducedLifetimeEnergyACDiff.Float32
				latestAccumulatedInfo.BatteryLifetimeEnergyACDiff = latestLogDaily.BatteryLifetimeEnergyACDiff.Float32
				latestAccumulatedInfo.GridLifetimeEnergyACDiff = latestLogDaily.GridLifetimeEnergyACDiff.Float32
				latestAccumulatedInfo.LoadSelfConsumedEnergyPercentAC = latestLogDaily.LoadSelfConsumedEnergyPercentAC.Float32
				latestAccumulatedInfo.PreUbiikCost = int(latestLogDaily.OffPeakPeriodPreUbiikCost.Float32 +
					latestLogDaily.OnPeakPeriodPreUbiikCost.Float32 +
					latestLogDaily.MidPeakPeriodPreUbiikCost.Float32)
				latestAccumulatedInfo.PostUbiikCost = int(latestLogDaily.OffPeakPeriodPostUbiikCost.Float32 +
					latestLogDaily.OnPeakPeriodPostUbiikCost.Float32 +
					latestLogDaily.MidPeakPeriodPostUbiikCost.Float32)
				latestAccumulatedInfoArray = append(latestAccumulatedInfoArray, latestAccumulatedInfo)
			case "month":
				latestLogMonthly := (s.Index(i).Interface()).(*deremsmodels.CCDataLogCalculatedMonthly)
				latestAccumulatedInfo := &LatestAccumulatedInfo{}
				latestAccumulatedInfo.Timestamps = int(latestLogMonthly.LatestLogDate.Unix())
				latestAccumulatedInfo.LoadConsumedLifetimeEnergyACDiff = latestLogMonthly.LoadConsumedLifetimeEnergyACDiff.Float32
				latestAccumulatedInfo.PvProducedLifetimeEnergyACDiff = latestLogMonthly.PvProducedLifetimeEnergyACDiff.Float32
				latestAccumulatedInfo.BatteryLifetimeEnergyACDiff = latestLogMonthly.BatteryLifetimeEnergyACDiff.Float32
				latestAccumulatedInfo.GridLifetimeEnergyACDiff = latestLogMonthly.GridLifetimeEnergyACDiff.Float32
				latestAccumulatedInfo.LoadSelfConsumedEnergyPercentAC = latestLogMonthly.LoadSelfConsumedEnergyPercentAC.Float32
				latestAccumulatedInfoArray = append(latestAccumulatedInfoArray, latestAccumulatedInfo)
			}
		}
	}
	return
}

func (s defaultDevicesService) getLatestComputedDemandState(gwUUID string, startTimeIndex, endTimeIndex, endTime time.Time) (latestComputedDemandState *LatestComputedDemandState) {
	if endTimeIndex == endTime && startTimeIndex.Add(15*time.Minute).Unix() > endTimeIndex.Unix() {
		return nil
	}
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
		return nil
	}

	latestComputedDemandState.Timestamps = int(latestLog.LogDate.Unix())
	latestComputedDemandState.GridLifetimeEnergyACDiffToPower = utils.Division(
		utils.Diff(latestLog.GridLifetimeEnergyAC.Float32,
			firstLog.GridLifetimeEnergyAC.Float32), (15.0 / 60.0))
	latestComputedDemandState.GridContractPowerAC = utils.ThreeDecimalPlaces(latestLog.GridContractPowerAC.Float32)
	return
}

func (s defaultDevicesService) getTimeOfUseOnPeakTime(gwUUID string, t time.Time) (timeOfUseOnPeakTime map[string]interface{}, err error) {
	timeOfUseOnPeakTime = make(map[string]interface{})
	localStartTime, tous, err := s.billing.GetTOUsOfLocalTime(gwUUID, t)
	if err != nil {
		return
	}

	// 1. timezone
	timeOfUseOnPeakTime["timezone"] = localStartTime.Format(utils.ZHHMM)
	// 2. onPeak
	periods := s.getPeriodsByPeakType("On-peak", tous)
	for _, period := range periods {
		delete(period, "touRate")
	}
	timeOfUseOnPeakTime["onPeak"] = periods
	return
}
