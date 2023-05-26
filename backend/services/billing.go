package services

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// BillingType godoc
type BillingType struct {
	TOULocationID int64
	VoltageType   string
	TOUType       string
}

// BillingParams godoc
type BillingParams struct {
	Timezone string     `json:"timezone"`
	Rates    []RateInfo `json:"rates"`
}

// RateInfo godoc
type RateInfo struct {
	Date             string  `json:"date"`
	Interval         string  `json:"interval"`
	DemandChargeRate float32 `json:"demandChargeRate"`
	TOURate          float32 `json:"touRate"`
}

// BillingService godoc
type BillingService interface {
	GetTOUsOfLocalTime(gwUUID string, t time.Time) (localTime time.Time, tous []*deremsmodels.Tou, err error)
	GetBillingTypeByLocationID(locationID int64) (billingType BillingType, err error)
	GetLocalTime(touLocationID int64, t time.Time) (localTime time.Time, err error)
	GetPeriodTypeOfDay(touLocationID int64, t time.Time) (periodType string)
	IsSummer(voltageType string, t time.Time) bool
	GetPeakType(localTime time.Time, tous []*deremsmodels.Tou) (peakType string, err error)
	GenerateBillingParams(gateway *deremsmodels.Gateway, sendNow bool) (billingParamsJSON []byte, err error)
	GetWeeklyBillingParamsByType(billingType BillingType, localTime time.Time, sendNow bool) (billingParamsJSON []byte, err error)
	GetSundayOfBillingWeek(t time.Time, sendNow bool) (timeOnSunday time.Time)
	SendAIBillingParamsToGateway(cfg *viper.Viper, billingParamsJSON []byte, uuid string)
}

type defaultBillingService struct {
	repo *repository.Repository
}

// NewBillingService godoc
func NewBillingService(repo *repository.Repository) BillingService {
	return &defaultBillingService{repo}
}

// GetTOUsOfLocalTime godoc
func (s defaultBillingService) GetTOUsOfLocalTime(gwUUID string, t time.Time) (localTime time.Time, tous []*deremsmodels.Tou, err error) {
	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(gwUUID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}
	billingType, err := s.GetBillingTypeByLocationID(gateway.LocationID.Int64)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.GetBillingTypeByLocationID",
			"err":       err,
		}).Error()
		return
	}
	localTime, err = s.GetLocalTime(billingType.TOULocationID, t)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.GetLocalTime",
			"err":       err,
		}).Error()
		return
	}
	periodType := s.GetPeriodTypeOfDay(billingType.TOULocationID, localTime)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.GetPeriodTypeOfDay",
			"err":       err,
		}).Error()
		return
	}
	isSummer := s.IsSummer(billingType.VoltageType, localTime)
	tous, err = s.repo.TOU.GetTOUsByTOUInfo(billingType.TOULocationID, billingType.VoltageType, billingType.TOUType, periodType, isSummer, localTime.Format(utils.YYYYMMDD))
	if err == nil && len(tous) == 0 {
		err = e.ErrNewBillingsNotExist
	}
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.TOU.GetTOUsByTOUInfo",
			"err":       err,
		}).Error()
	}
	return
}

// GetBillingTypeByLocationID godoc
func (s defaultBillingService) GetBillingTypeByLocationID(locationID int64) (billingType BillingType, err error) {
	location, err := s.repo.Location.GetLocationByLocationID(locationID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.Location.GetLocationByLocationID",
			"err":       err,
		}).Error()
		return
	}
	billingType = BillingType{
		TOULocationID: location.TOULocationID.Int64,
		VoltageType:   location.VoltageType.String,
		TOUType:       location.TOUType.String,
	}
	return
}

// GetLocalTime godoc
func (s defaultBillingService) GetLocalTime(touLocationID int64, t time.Time) (localTime time.Time, err error) {
	touLocation, err := s.repo.TOU.GetTOULocationByTOULocationID(touLocationID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.TOU.GetTOULocationByTOULocationID",
			"err":       err,
		}).Error()
		return
	}
	localLocation, err := time.LoadLocation(touLocation.Location.String)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "time.LoadLocation",
			"err":       err,
		}).Error()
		return
	}
	localTime = t.In(localLocation)
	return
}

// GetPeriodTypeOfDay godoc
func (s defaultBillingService) GetPeriodTypeOfDay(touLocationID int64, t time.Time) (periodType string) {
	// The day is holiday or not
	count, _ := s.repo.TOU.CountHolidayByDay(touLocationID, t.Format(utils.YYYY), t.Format(utils.YYYYMMDD))

	switch {
	case count > 0 || t.Weekday() == time.Sunday:
		periodType = "Sunday & Holiday"
	case t.Weekday() == time.Saturday:
		periodType = "Saturday"
	default:
		periodType = "Weekdays"
	}
	return
}

// IsSummer godoc
func (s defaultBillingService) IsSummer(voltageType string, t time.Time) bool {
	/* XXX: Hardcode TPC ~2022 summer is 06/30~09/30
	   TPC 2023~ summer low voltage is 06/30~09/30 and the other is 05/16~10/15 */
	if (t.Year() <= 2022) || (t.Year() > 2022 && voltageType == "Low voltage") {
		switch t.Month() {
		case time.June, time.July, time.August, time.September:
			return true
		}
		return false
	} else {
		billingDate, _ := time.Parse(utils.YYYYMMDD, t.Format(utils.YYYYMMDD))
		summerStartDate, _ := time.Parse(utils.YYYYMMDD, strconv.Itoa(t.Year())+"-05-16")
		summerEndDate, _ := time.Parse(utils.YYYYMMDD, strconv.Itoa(t.Year())+"-10-15")
		if (summerStartDate.Before(billingDate) && summerEndDate.After(billingDate)) || summerStartDate.Equal(billingDate) || summerEndDate.Equal(billingDate) {
			return true
		}
		return false
	}
}

// GetPeakType godoc
func (s defaultBillingService) GetPeakType(localTime time.Time, tous []*deremsmodels.Tou) (peakType string, err error) {
	loc := time.FixedZone(localTime.Zone())
	localTime, err = time.ParseInLocation(utils.HHMMSS24h, localTime.Format(utils.HHMMSS24h), loc)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "time.ParseInLocation",
			"err":       err,
		}).Error()
		return
	}

	for _, tou := range tous {
		startTime, err := time.ParseInLocation(utils.HHMMSS24h, tou.PeriodStime.String, loc)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "time.ParseInLocation",
				"err":       err,
			}).Error()
			break
		}
		endTime, err := time.ParseInLocation(utils.HHMMSS24h, tou.PeriodEtime.String, loc)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "time.ParseInLocation",
				"err":       err,
			}).Error()
			break
		}
		if tou.PeriodEtime.String == "00:00:00" {
			endTime = endTime.AddDate(0, 0, 1)
		}
		if localTime.After(startTime) && localTime.Before(endTime) {
			peakType = tou.PeakType.String
			break
		}
	}
	return
}

func (s defaultBillingService) GenerateBillingParams(gateway *deremsmodels.Gateway, sendNow bool) (billingParamsJSON []byte, err error) {
	billingType, err := s.GetBillingTypeByLocationID(gateway.LocationID.Int64)
	if err != nil {
		return
	}
	localTime, err := s.GetLocalTime(billingType.TOULocationID, time.Now().UTC())
	if err != nil {
		return
	}
	billingParamsJSON, err = s.GetWeeklyBillingParamsByType(billingType, localTime, sendNow)
	return
}

func (s defaultBillingService) GetWeeklyBillingParamsByType(billingType BillingType, localTime time.Time, sendNow bool) (billingParamsJSON []byte, err error) {
	var billingParams BillingParams
	// 1. Get timezone
	log.Debug("timezone: ", localTime.Format(utils.ZHHMM))
	billingParams.Timezone = localTime.Format(utils.ZHHMM)
	// 2. Get Sunday of billing week
	timeOnSunday := s.GetSundayOfBillingWeek(localTime, sendNow)
	log.Debug("timeOnSunday: ", timeOnSunday)
	// 3. Get one week billing params
	for i := 0; i < 7; i++ {
		timeOfEachDay := timeOnSunday.AddDate(0, 0, i)
		log.Debug("timeOfEachDay: ", timeOfEachDay)
		// 3-1. Get period type
		periodType := s.GetPeriodTypeOfDay(billingType.TOULocationID, timeOfEachDay)
		log.Debug("periodType: ", periodType)
		// 3-2. The day is summmer or not
		isSummer := s.IsSummer(billingType.VoltageType, timeOfEachDay)
		log.Debug("isSummer: ", isSummer)
		// 3-3. Get billings
		tous, err := s.repo.TOU.GetTOUsByTOUInfo(billingType.TOULocationID, billingType.VoltageType, billingType.TOUType, periodType, isSummer, timeOfEachDay.Format(utils.YYYYMMDD))
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "repo.TOU.GetTOUsByTOUInfo",
				"err":       err,
			}).Error()
			break
		}
		for _, tou := range tous {
			rate, err := s.touToRate(tou, timeOfEachDay)
			if err != nil {
				break
			}
			billingParams.Rates = append(billingParams.Rates, rate)
		}
	}
	if err != nil {
		return
	}
	billingParamsJSON, err = json.Marshal(billingParams)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
	}
	log.Debug("billingParamsJSON: ", string(billingParamsJSON))
	return
}

func (s defaultBillingService) GetSundayOfBillingWeek(t time.Time, sendNow bool) (timeOnSunday time.Time) {
	weekDay := t.Weekday()
	var offset = 0
	// Get Sunday of this week
	if weekDay != time.Sunday {
		offset = int(time.Sunday - weekDay)
	}
	// Get Sunday of next week
	if !sendNow {
		offset += 7
	}
	timeOnSunday = t.AddDate(0, 0, offset)
	return
}

func (s defaultBillingService) touToRate(tou *deremsmodels.Tou, timeOfEachDay time.Time) (rate RateInfo, err error) {
	log.WithFields(log.Fields{
		"peak type":    tou.PeakType,
		"period stime": tou.PeriodStime,
		"period etime": tou.PeriodEtime,
		"basic rate":   tou.BasicRate.Float32,
		"flow rate":    tou.FlowRate.Float32,
		"enable at":    tou.EnableAt,
		"disable at":   tou.DisableAt,
	}).Debug()

	interval, err := s.getTOUInterval(tou.PeriodStime.String, tou.PeriodEtime.String)
	if err != nil {
		return
	}

	rate.Date = timeOfEachDay.Format(utils.YYYYMMDD)
	rate.Interval = interval
	rate.DemandChargeRate = tou.BasicRate.Float32
	rate.TOURate = tou.FlowRate.Float32
	return
}

func (s defaultBillingService) getTOUInterval(periodStime, periodEtime string) (interval string, err error) {
	startTime, err := time.Parse(utils.HHMMSS24h, periodStime)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "time.Parse",
			"err":       err,
		}).Error()
		return
	}
	startTimeString := startTime.Format(utils.HHMM24h)
	endTime, err := time.Parse(utils.HHMMSS24h, periodEtime)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "time.Parse",
			"err":       err,
		}).Error()
		return
	}
	endTimeString := endTime.Format(utils.HHMM24h)
	if endTimeString == "0000" {
		endTimeString = "2400"
	}
	interval = startTimeString + "-" + endTimeString
	return
}

func (s defaultBillingService) SendAIBillingParamsToGateway(cfg *viper.Viper, billingParamsJSON []byte, uuid string) {
	sendAIBillingParamsToLocalGW := strings.Replace(kafka.SendAIBillingParamsToLocalGW, "{gw-id}", uuid, 1)
	log.Debug("sendAIBillingParamsToLocalGW: ", sendAIBillingParamsToLocalGW)
	kafka.Produce(cfg, sendAIBillingParamsToLocalGW, string(billingParamsJSON))
}
