package apps

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

const (
	// ZHHMM ±hhmm
	ZHHMM = "Z0700"
	// YYYY YYYY
	YYYY = "2006"
	// YYYYMMDD YYYY-MM-DD
	YYYYMMDD = "2006-01-02"
	// HHMMSS24h HH:MM:SS
	HHMMSS24h = "15:04:05"
	// HHMM24h HHMM
	HHMM24h = "1504"
)

// BillingType godoc
type BillingType struct {
	TOULocationID int
	VoltageType   string
	TOUType       string
}

// BillingParams godoc
type BillingParams struct {
	Timezone string     `json:"timezone"`
	Rate     []RateInfo `json:"rate"`
}

// RateInfo godoc
type RateInfo struct {
	Date             string  `json:"date"`
	Interval         string  `json:"interval"`
	DemandChargeRate float32 `json:"demandChargeRate"`
	TOURate          float32 `json:"touRate"`
}

// NewBillingWorker godoc
func NewBillingWorker(
	ctx context.Context,
	cfg *viper.Viper,
	repo *repository.Repository,
	name string,
) {
	// 1. Send in beginning
	sendAIBillingParams(cfg, repo, true)

	// 2. Send at 04:00 on Saturday in UTC(12:00 on Saturday in UTC+0800)
	c := cron.New()
	c.AddFunc("0 4 * * 6", func() { sendAIBillingParams(cfg, repo, false) })
	c.Start()
	log.Info("serving: ", name)
	<-ctx.Done()
	log.Info("graceful stopping: ", name)
	c.Stop()
	log.Info("stopped: ", name)
}

func sendAIBillingParams(cfg *viper.Viper, repo *repository.Repository, sendNow bool) {
	// TODO: modify log format
	log.Info("sendAIBillingParams")
	gateways, err := getGateways(repo)
	if err != nil {
		return
	}

	for _, gateway := range gateways {
		billingParamsJSON, err := generateBillingParams(repo, gateway, sendNow)
		if err != nil {
			continue
		}
		sendAIBillingParamsToGateway(cfg, billingParamsJSON, gateway.UUID)
	}
}

func getGateways(repo *repository.Repository) (gateways []*deremsmodels.Gateway, err error) {
	gateways, err = repo.Gateway.GetGateways()
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "repo.Gateway.GetGateways",
			"err":       err,
		}).Error()
	}
	return
}

func generateBillingParams(repo *repository.Repository, gateway *deremsmodels.Gateway, sendNow bool) (billingParamsJSON []byte, err error) {
	billingType, err := getBillingTypeByCustomerID(repo, gateway.CustomerID)
	if err != nil {
		return
	}
	localTime, err := getLocalTime(repo, billingType.TOULocationID)
	if err != nil {
		return
	}
	billingParamsJSON, err = getWeeklyBillingParamsByType(repo, billingType, localTime, sendNow)
	return
}

func getBillingTypeByCustomerID(repo *repository.Repository, customerID int) (billingType BillingType, err error) {
	customer, err := repo.Customer.GetCustomerByCustomerID(customerID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "repo.Customer.GetCustomerByCustomerID",
			"err":       err,
		}).Error()
		return
	}
	billingType = BillingType{
		TOULocationID: customer.TOULocationID.Int,
		VoltageType:   customer.VoltageType.String,
		TOUType:       customer.TOUType.String,
	}
	return
}

func getLocalTime(repo *repository.Repository, touLocationID int) (localTime time.Time, err error) {
	touLocation, err := repo.TOU.GetTOULocationByTOULocationID(touLocationID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "repo.TOU.GetTOULocationByTOULocationID",
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
	localTime = time.Now().In(localLocation)
	return
}

func getWeeklyBillingParamsByType(repo *repository.Repository, billingType BillingType, localTime time.Time, sendNow bool) (billingParamsJSON []byte, err error) {
	var billingParams BillingParams
	// 1. Get timezone
	log.Debug("timezone: ", localTime.Format(ZHHMM))
	billingParams.Timezone = localTime.Format(ZHHMM)
	// 2. Get Sunday of billing week
	timeOnSunday := getSundayOfBillingWeek(localTime, sendNow)
	log.Debug("timeOnSunday: ", timeOnSunday)
	// 3. Get one week billing params
	for i := 0; i < 7; i++ {
		timeOfEachDay := timeOnSunday.AddDate(0, 0, i)
		log.Debug("timeOfEachDay: ", timeOfEachDay)
		// 3-1. Get period type
		periodType := getPeriodTypeOfDay(repo, billingType.TOULocationID, timeOfEachDay)
		log.Debug("periodType: ", periodType)
		// 3-2. The day is summmer or not
		isSummer := isSummer(timeOfEachDay)
		log.Debug("isSummer: ", isSummer)
		// 3-3. Get billings
		billings, err := repo.TOU.GetBillingsByTOUInfo(billingType.TOULocationID, billingType.VoltageType, billingType.TOUType, periodType, isSummer, timeOfEachDay.Format(YYYYMMDD))
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "repo.TOU.GetBillingsByTOUInfo",
				"err":       err,
			}).Error()
			break
		}
		for _, billing := range billings {
			log.WithFields(log.Fields{
				"peak type":    billing.PeakType,
				"period stime": billing.PeriodStime,
				"period etime": billing.PeriodEtime,
				"basic rate":   billing.BasicRate.Float32,
				"flow rate":    billing.FlowRate.Float32,
				"enable at":    billing.EnableAt,
				"disable at":   billing.DisableAt,
			}).Debug()

			startTime, err := time.Parse(HHMMSS24h, billing.PeriodStime.String)
			if err != nil {
				log.WithFields(log.Fields{
					"caused-by": "time.Parse",
					"err":       err,
				}).Error()
				break
			}
			startTimeString := startTime.Format(HHMM24h)
			endTime, err := time.Parse(HHMMSS24h, billing.PeriodEtime.String)
			if err != nil {
				log.WithFields(log.Fields{
					"caused-by": "time.Parse",
					"err":       err,
				}).Error()
				break
			}
			endTimeString := endTime.Format(HHMM24h)
			if endTimeString == "0000" {
				endTimeString = "2400"
			}
			interval := startTimeString + "-" + endTimeString

			var rate RateInfo
			rate.Date = timeOfEachDay.Format(YYYYMMDD)
			rate.Interval = interval
			rate.DemandChargeRate = billing.BasicRate.Float32
			rate.TOURate = billing.FlowRate.Float32
			billingParams.Rate = append(billingParams.Rate, rate)
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

func getSundayOfBillingWeek(t time.Time, sendNow bool) (timeOnSunday time.Time) {
	weekDay := t.Weekday()
	var offset = 0
	if sendNow {
		// Get Sunday of this week
		if weekDay != time.Sunday {
			offset = int(time.Sunday - weekDay)
		}
	} else {
		// Get Sunday of next week (Check on Saturday)
		offset = 1
	}
	timeOnSunday = t.AddDate(0, 0, offset)
	return
}

func getPeriodTypeOfDay(repo *repository.Repository, touLocationID int, t time.Time) (periodType string) {
	// The day is holiday or not
	count, _ := repo.TOU.GetHolidayByDay(touLocationID, t.Format(YYYY), t.Format(YYYYMMDD))
	if count > 0 {
		periodType = "Sunday & Holiday"
		return
	}

	switch t.Weekday() {
	case time.Sunday:
		periodType = "Sunday & Holiday"
	case time.Saturday:
		periodType = "Saturday"
	default:
		periodType = "Weekdays"
	}
	return
}

func isSummer(t time.Time) bool {
	// XXX: Hardcode TPC summer is 06/30~09/30
	if t.Month() == time.June || t.Month() == time.July || t.Month() == time.August || t.Month() == time.September {
		return true
	}
	return false
}

func sendAIBillingParamsToGateway(cfg *viper.Viper, billingParamsJSON []byte, uuid string) {
	sendAIBillingParamsToLocalGW := strings.Replace(kafka.SendAIBillingParamsToLocalGW, "{gw-id}", uuid, 1)
	log.Debug("sendAIBillingParamsToLocalGW: ", sendAIBillingParamsToLocalGW)
	kafka.Produce(cfg, sendAIBillingParamsToLocalGW, string(billingParamsJSON))
}
