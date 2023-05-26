package apps

import (
	"context"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"der-ems/internal/utils"
	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
	"der-ems/services"
)

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
		billingParamsJSON, err := billing.GenerateBillingParams(gateway, sendNow)
		if err != nil {
			continue
		}
		kafka.SendDataToGateways(cfg, kafka.SendAIBillingParamsToLocalGW, billingParamsJSON, []string{gateway.UUID})
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
