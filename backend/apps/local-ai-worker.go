package apps

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/utils"
	"der-ems/kafka"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// LocalAIWorker godoc
type LocalAIWorker struct {
	kafka.SimpleConsumer
}

type localAIConsumerHandler struct {
	cfg  *viper.Viper
	repo *repository.Repository
}

func (localAIConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (localAIConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h localAIConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.WithFields(log.Fields{
			"topic":     msg.Topic,
			"partition": msg.Partition,
			"offset":    msg.Offset,
			"value":     string(msg.Value),
		}).Info("consuming")

		if msg.Topic == kafka.ReceiveLocalAIData {
			h.processLocalAIData(msg.Value)
		}

		// Mark a message as consumed
		sess.MarkMessage(msg, "")
	}
	return nil
}

// NewLocalAIWorker godoc
func NewLocalAIWorker(
	ctx context.Context,
	cfg *viper.Viper,
	repo *repository.Repository,
	name string,
) (w *LocalAIWorker) {
	topics := []string{
		kafka.ReceiveLocalAIData,
	}
	handler := localAIConsumerHandler{
		cfg:  cfg,
		repo: repo,
	}

	simpleConsumer, err := kafka.NewSimpleConsumer(ctx, cfg, name, handler, topics)
	if err != nil {
		return
	}

	w = &LocalAIWorker{
		SimpleConsumer: *simpleConsumer,
	}

	return
}

// MainLoop godoc
func (w *LocalAIWorker) MainLoop() {
	w.SimpleConsumer.MainLoop()
}

func (h localAIConsumerHandler) processLocalAIData(msg []byte) {
	utils.PrintFunctionName()
	h.saveLocalAIData(msg)
}

func (h localAIConsumerHandler) saveLocalAIData(msg []byte) (err error) {
	gwIDValue, timestampValue, data, err := utils.AssertGatewayMessage(msg)
	if err != nil {
		return
	}
	gwUUID := gwIDValue.(string)
	logDate := int64(timestampValue.(float64))
	dataJSON, _ := json.Marshal(data)

	aiData := &deremsmodels.AiDatum{
		GWUUID:      gwUUID,
		LogDate:     time.Unix(logDate, 0),
		LocalAiData: null.NewJSON(dataJSON, true),
	}

	gateway, err := h.repo.Gateway.GetGatewayByGatewayUUID(gwUUID)
	if err == nil {
		aiData.GWID = null.NewInt(gateway.ID, true)
		aiData.CustomerID = null.NewInt(gateway.CustomerID, true)
	} else {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Warn()
	}

	log.WithFields(log.Fields{
		deremsmodels.AiDatumColumns.GWUUID:      aiData.GWUUID,
		deremsmodels.AiDatumColumns.LogDate:     aiData.LogDate,
		deremsmodels.AiDatumColumns.CustomerID:  aiData.CustomerID,
		deremsmodels.AiDatumColumns.LocalAiData: string(aiData.LocalAiData.JSON),
	}).Debug("upsert local AI data")
	err = h.repo.AIData.UpsertAIData(aiData)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.AIData.UpsertAIData",
			"err":       err,
		}).Error()
	}
	return
}
