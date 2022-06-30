package apps

import (
	"context"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// BillingType godoc
type BillingType struct {
	PowerCompany string
	VoltageType  string
	TouType      string
}

// NewBillingWorker godoc
func NewBillingWorker(
	ctx context.Context,
	cfg *viper.Viper,
	repo *repository.Repository,
	name string,
) {
	// 1. Send in beginning
	sendAIBillingParams(cfg, repo)

	// 2. Send at 12:00 on Saturday in UTC+0800
	nyc, _ := time.LoadLocation("Asia/Hong_Kong")
	c := cron.New(cron.WithLocation(nyc))
	// TODO: remove test code
	//c.AddFunc("0 12 * * 6", func() { sendAIBillingParams(cfg, repo) })
	c.AddFunc("*/1 * * * *", func() { sendAIBillingParams(cfg, repo) })
	c.Start()
	log.Info("serving: ", name)
	<-ctx.Done()
	log.Info("graceful stopping: ", name)
	c.Stop()
	log.Info("stopped: ", name)
}

func sendAIBillingParams(cfg *viper.Viper, repo *repository.Repository) {
	// TODO: modify log format
	log.Info("sendAIBillingParams")
	gateways, err := getGateways(repo)
	if err != nil {
		return
	}

	for _, gateway := range gateways {
		billingJSON, _ := generateBillingParams(repo, gateway)
		sendAIBillingParamsToGateway(cfg, billingJSON, gateway.UUID)
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

func generateBillingParams(repo *repository.Repository, gateway *deremsmodels.Gateway) (billingJSON []byte, err error) {
	billingType, err := getBillingTypeByCustomerID(repo, gateway.CustomerID)
	if err != nil {
		return
	}
	billingJSON, err = getWeeklyBillingParamsByType(billingType)
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
		PowerCompany: customer.PowerCompany.String,
		VoltageType:  customer.VoltageType.String,
		TouType:      customer.TouType.String,
	}
	return
}

func getWeeklyBillingParamsByType(billingType BillingType) (billingJSON []byte, err error) {
	// TODO: remove test code and implement
	billingJSON = []byte("test")
	return
}

func sendAIBillingParamsToGateway(cfg *viper.Viper, billingJSON []byte, uuid string) {
	sendAIBillingParamsToLocalGW := strings.Replace(kafka.SendAIBillingParamsToLocalGW, "{gw-id}", uuid, 1)
	log.Debug("sendAIBillingParamsToLocalGW: ", sendAIBillingParamsToLocalGW)
	kafka.Produce(cfg, sendAIBillingParamsToLocalGW, string(billingJSON))
}
