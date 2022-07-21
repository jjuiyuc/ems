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
	"der-ems/utils"
)

// LocalCCWorker godoc
type LocalCCWorker struct {
	kafka.SimpleConsumer
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
			h.processLocalCCData(msg.Value)
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

func (h localCCConsumerHandler) processLocalCCData(msg []byte) {
	utils.PrintFunctionName()
	h.saveLocalCCData(msg)
}

func (h localCCConsumerHandler) saveLocalCCData(msg []byte) (err error) {
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
		err = e.ErrNewKeyNotExist(gwID)
		log.WithFields(log.Fields{
			"caused-by": gwID,
			"err":       err,
		}).Error()
		return
	}
	timestampValue := data[timestamp]
	if timestampValue == nil {
		err = e.ErrNewKeyNotExist(timestamp)
		log.WithFields(log.Fields{
			"caused-by": timestamp,
			"err":       err,
		}).Error()
		return
	}
	gwUUID := gwIDValue.(string)
	logDate := int64(timestampValue.(float64))
	dataJSON, _ := json.Marshal(data)

	ccData := &deremsmodels.CCDatum{
		GWUUID:      gwUUID,
		LogDate:     time.Unix(logDate, 0),
		LocalCCData: null.NewJSON(dataJSON, true),
	}

	gateway, err := h.repo.Gateway.GetGatewayByGatewayUUID(gwUUID)
	if err == nil {
		ccData.CustomerID = null.NewInt(gateway.CustomerID, true)
	} else {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.Gateway.GetGatewayByGatewayUUID",
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
