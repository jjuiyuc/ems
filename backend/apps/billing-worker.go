package apps

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"der-ems/internal/utils"
	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

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

// NewBillingWorker godoc
func NewBillingWorker(
	ctx context.Context,
	cfg *viper.Viper,
	repo *repository.Repository,
	name string,
) {
	// 1. Send at the beginning
	sendAIBillingParams(cfg, repo, true)

	// 2. Send at 04:00 on Saturday in UTC(12:00 on Saturday in UTC+0800)
	c := cron.New()
	c.AddFunc(cfg.GetString("cron.billing"), func() { sendAIBillingParams(cfg, repo, false) })
	c.Start()
	log.Info("serving: ", name)
	<-ctx.Done()
	log.Info("graceful stopping: ", name)
	c.Stop()
	log.Info("stopped: ", name)
}

func sendAIBillingParams(cfg *viper.Viper, repo *repository.Repository, sendNow bool) {
	utils.PrintFunctionName()
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
	billingType, err := utils.GetBillingTypeByCustomerID(repo, gateway.CustomerID)
	if err != nil {
		return
	}
	localTime, err := utils.GetLocalTime(repo, billingType.TOULocationID, time.Now().UTC())
	if err != nil {
		return
	}
	billingParamsJSON, err = getWeeklyBillingParamsByType(repo, billingType, localTime, sendNow)
	return
}

func getWeeklyBillingParamsByType(repo *repository.Repository, billingType utils.BillingType, localTime time.Time, sendNow bool) (billingParamsJSON []byte, err error) {
	var billingParams BillingParams
	// 1. Get timezone
	log.Debug("timezone: ", localTime.Format(utils.ZHHMM))
	billingParams.Timezone = localTime.Format(utils.ZHHMM)
	// 2. Get Sunday of billing week
	timeOnSunday := getSundayOfBillingWeek(localTime, sendNow)
	log.Debug("timeOnSunday: ", timeOnSunday)
	// 3. Get one week billing params
	for i := 0; i < 7; i++ {
		timeOfEachDay := timeOnSunday.AddDate(0, 0, i)
		log.Debug("timeOfEachDay: ", timeOfEachDay)
		// 3-1. Get period type
		periodType := utils.GetPeriodTypeOfDay(repo, billingType.TOULocationID, timeOfEachDay)
		log.Debug("periodType: ", periodType)
		// 3-2. The day is summmer or not
		isSummer := utils.IsSummer(timeOfEachDay)
		log.Debug("isSummer: ", isSummer)
		// 3-3. Get billings
		billings, err := repo.TOU.GetBillingsByTOUInfo(billingType.TOULocationID, billingType.VoltageType, billingType.TOUType, periodType, isSummer, timeOfEachDay.Format(utils.YYYYMMDD))
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

			interval, err := getBillingInterval(billing.PeriodStime.String, billing.PeriodEtime.String)
			if err != nil {
				break
			}

			var rate RateInfo
			rate.Date = timeOfEachDay.Format(utils.YYYYMMDD)
			rate.Interval = interval
			rate.DemandChargeRate = billing.BasicRate.Float32
			rate.TOURate = billing.FlowRate.Float32
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

func getSundayOfBillingWeek(t time.Time, sendNow bool) (timeOnSunday time.Time) {
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

func getBillingInterval(periodStime, periodEtime string) (interval string, err error) {
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

func sendAIBillingParamsToGateway(cfg *viper.Viper, billingParamsJSON []byte, uuid string) {
	sendAIBillingParamsToLocalGW := strings.Replace(kafka.SendAIBillingParamsToLocalGW, "{gw-id}", uuid, 1)
	log.Debug("sendAIBillingParamsToLocalGW: ", sendAIBillingParamsToLocalGW)
	kafka.Produce(cfg, sendAIBillingParamsToLocalGW, string(billingParamsJSON))
}
