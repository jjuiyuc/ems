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
	"der-ems/services"
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
	billing services.BillingService,
	name string,
) {
	// 1. Send at the beginning
	sendAIBillingParams(cfg, repo, billing, true)

	// 2. Send at 04:00 on Saturday in UTC(12:00 on Saturday in UTC+0800)
	c := cron.New()
	c.AddFunc(cfg.GetString("cron.billing"), func() { sendAIBillingParams(cfg, repo, billing, false) })
	c.Start()
	log.Info("serving: ", name)
	<-ctx.Done()
	log.Info("graceful stopping: ", name)
	c.Stop()
	log.Info("stopped: ", name)
}

func sendAIBillingParams(cfg *viper.Viper, repo *repository.Repository, billing services.BillingService, sendNow bool) {
	utils.PrintFunctionName()
	gateways, err := getGateways(repo)
	if err != nil {
		return
	}

	for _, gateway := range gateways {
		billingParamsJSON, err := generateBillingParams(repo, billing, gateway, sendNow)
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

func generateBillingParams(repo *repository.Repository, billing services.BillingService, gateway *deremsmodels.Gateway, sendNow bool) (billingParamsJSON []byte, err error) {
	billingType, err := billing.GetBillingTypeByLocationID(gateway.LocationID.Int64)
	if err != nil {
		return
	}
	localTime, err := billing.GetLocalTime(billingType.TOULocationID, time.Now().UTC())
	if err != nil {
		return
	}
	billingParamsJSON, err = getWeeklyBillingParamsByType(repo, billing, billingType, localTime, sendNow)
	return
}

func getWeeklyBillingParamsByType(repo *repository.Repository, billing services.BillingService, billingType services.BillingType, localTime time.Time, sendNow bool) (billingParamsJSON []byte, err error) {
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
		periodType := billing.GetPeriodTypeOfDay(billingType.TOULocationID, timeOfEachDay)
		log.Debug("periodType: ", periodType)
		// 3-2. The day is summmer or not
		isSummer := billing.IsSummer(billingType.VoltageType, timeOfEachDay)
		log.Debug("isSummer: ", isSummer)
		// 3-3. Get billings
		tous, err := repo.TOU.GetTOUsByTOUInfo(billingType.TOULocationID, billingType.VoltageType, billingType.TOUType, periodType, isSummer, timeOfEachDay.Format(utils.YYYYMMDD))
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "repo.TOU.GetTOUsByTOUInfo",
				"err":       err,
			}).Error()
			break
		}
		for _, tou := range tous {
			rate, err := touToRate(tou, timeOfEachDay)
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

func touToRate(tou *deremsmodels.Tou, timeOfEachDay time.Time) (rate RateInfo, err error) {
	log.WithFields(log.Fields{
		"peak type":    tou.PeakType,
		"period stime": tou.PeriodStime,
		"period etime": tou.PeriodEtime,
		"basic rate":   tou.BasicRate.Float32,
		"flow rate":    tou.FlowRate.Float32,
		"enable at":    tou.EnableAt,
		"disable at":   tou.DisableAt,
	}).Debug()

	interval, err := getTOUInterval(tou.PeriodStime.String, tou.PeriodEtime.String)
	if err != nil {
		return
	}

	rate.Date = timeOfEachDay.Format(utils.YYYYMMDD)
	rate.Interval = interval
	rate.DemandChargeRate = tou.BasicRate.Float32
	rate.TOURate = tou.FlowRate.Float32
	return
}

func getTOUInterval(periodStime, periodEtime string) (interval string, err error) {
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
