package apps

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/e"
	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// LocalCCWorker godoc
type LocalCCWorker struct {
	kafka.SimpleConsumer
}

var CloudCCParams = []string{
	"gwID", "timestamp", "gridAveragePowerAC", "gridProducedEnergyAC", "gridConsumedEnergyAC",
	"loadInstantaneousPowerAC", "loadAveragePowerAC", "loadBatteryAveragePowerAC", "loadPvAveragePowerAC", "loadBatteryConsumedEnergyAC",
	"loadPvConsumedEnergyAC", "loadConsumedEnergyAC", "batteryInstantaneousPowerAC", "batteryAveragePowerAC", "batteryProducedEnergyAC",
	"batteryConsumedEnergyAC", "batteryLifetimeEnergyAC", "batteryNameplateEnergy", "batteryVoltage", "batterySoC",
	"batterySoH", "pvInstantaneousPowerAC", "pvAveragePowerAC", "pvProducedEnergyAC", "allProducedEnergyAC", "allConsumedEnergyAC",
}

type localCCConsumerHandler struct {
	cfg  *viper.Viper
	repo *repository.Repository
}

func (localCCConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (localCCConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h localCCConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.WithFields(log.Fields{
			"topic":     msg.Topic,
			"partition": msg.Partition,
			"offset":    msg.Offset,
			"value":     string(msg.Value),
		}).Info("consuming")

		if msg.Topic == kafka.ReceiveLocalCCData {
			h.ProcessLocalCCData(msg.Value)
		}

		// Mark a message as consumed
		sess.MarkMessage(msg, "")
	}
	return nil
}

// NewLocalCCWorker godoc
func NewLocalCCWorker(
	ctx context.Context,
	cfg *viper.Viper,
	repo *repository.Repository,
	name string,
) (w *LocalCCWorker) {
	topics := []string{
		kafka.ReceiveLocalCCData,
	}
	handler := localCCConsumerHandler{
		cfg:  cfg,
		repo: repo,
	}

	simpleConsumer, err := kafka.NewSimpleConsumer(ctx, cfg, name, handler, topics)
	if err != nil {
		return
	}

	w = &LocalCCWorker{
		SimpleConsumer: *simpleConsumer,
	}

	return
}

// MainLoop godoc
func (w *LocalCCWorker) MainLoop() {
	w.SimpleConsumer.MainLoop()
}

// ProcessLocalCCData godoc
func (h localCCConsumerHandler) ProcessLocalCCData(msg []byte) {
	log.Debug("ProcessLocalCCData")
	err := h.SaveLocalCCData(msg)
	if err != nil {
		return
	}
	cloudCCParamsDataJson, err := h.GenerateCloudCCSendingInfo(msg)
	if err != nil {
		return
	}
	h.sendParamsToCloudCC(cloudCCParamsDataJson)
}

// SaveLocalCCData godoc
func (h localCCConsumerHandler) SaveLocalCCData(msg []byte) (err error) {
	const (
		gwID      = "gwID"
		timestamp = "timestamp"
	)
	var data map[string]interface{}
	err = json.Unmarshal(msg, &data)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}

	gwIDValue := data[gwID]
	if gwIDValue == nil {
		err = e.NewKeyNotExistError(gwID)
		log.WithFields(log.Fields{
			"caused-by": gwID,
			"err":       err,
		}).Error()
		return
	}
	timestampValue := data[timestamp]
	if timestampValue == nil {
		err = e.NewKeyNotExistError(timestamp)
		log.WithFields(log.Fields{
			"caused-by": timestamp,
			"err":       err,
		}).Error()
		return
	}
	gwUUID := gwIDValue.(string)
	logDate := int64(timestampValue.(float64))
	dataJson, _ := json.Marshal(data)

	ccData := &deremsmodels.CCDatum{
		GWUUID:      gwUUID,
		LogDate:     time.Unix(logDate, 0),
		LocalCCData: null.NewJSON(dataJson, true),
	}

	gateway, err := h.repo.Gateway.GetCustomerIDByGatewayUUID(gwUUID)
	if err == nil {
		ccData.CustomerID = null.NewInt(gateway.CustomerID, true)
	} else {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.Gateway.GetCustomerIDByGatewayUUID",
			"err":       err,
		}).Warn()
	}

	log.WithFields(log.Fields{
		deremsmodels.CCDatumColumns.GWUUID:      ccData.GWUUID,
		deremsmodels.CCDatumColumns.LogDate:     ccData.LogDate,
		deremsmodels.CCDatumColumns.CustomerID:  ccData.CustomerID,
		deremsmodels.CCDatumColumns.LocalCCData: string(ccData.LocalCCData.JSON),
	}).Debug("upsert local CC data")
	err = h.repo.CCData.UpsertCCData(ccData)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.CCData.UpsertCCData",
			"err":       err,
		}).Error()
	}
	return
}

// GenerateCloudCCSendingInfo godoc
func (h localCCConsumerHandler) GenerateCloudCCSendingInfo(msg []byte) (cloudCCParamsDataJson []byte, err error) {
	var data map[string]interface{}
	err = json.Unmarshal(msg, &data)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}

	cloudCCParamsData := make(map[string]interface{})
	for _, param := range CloudCCParams {
		value := data[param]
		if value == nil {
			err = e.NewKeyNotExistError(param)
			log.WithFields(log.Fields{
				"caused-by": param,
				"err":       err,
			}).Error()
			return
		}
		cloudCCParamsData[param] = value
	}

	cloudCCParamsDataJson, err = json.Marshal(cloudCCParamsData)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Marshal",
			"err":       err,
		}).Error()
	}
	return
}

func (h localCCConsumerHandler) sendParamsToCloudCC(cloudCCParamsDataJson []byte) {
	log.WithFields(log.Fields{
		"message": string(cloudCCParamsDataJson),
	}).Debug("send parameters to cloud CC")
	kafka.Produce(h.cfg, kafka.SendParamsToCloudCC, string(cloudCCParamsDataJson))
}
