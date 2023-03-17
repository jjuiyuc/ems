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
	h.saveLocalCCDataLog(msg)
}

func (h localCCConsumerHandler) saveLocalCCData(msg []byte) (err error) {
	gwIDValue, timestampValue, data, err := utils.AssertGatewayMessage(msg)
	if err != nil {
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
		ccData.GWID = null.NewInt64(gateway.ID, true)
		ccData.LocationID = gateway.LocationID
	} else {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Warn()
	}

	log.WithFields(log.Fields{
		deremsmodels.CCDatumColumns.GWUUID:      ccData.GWUUID,
		deremsmodels.CCDatumColumns.LogDate:     ccData.LogDate,
		deremsmodels.CCDatumColumns.LocationID:  ccData.LocationID,
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

func (h localCCConsumerHandler) saveLocalCCDataLog(msg []byte) (err error) {
	gwIDValue, timestampValue, data, err := utils.AssertGatewayMessage(msg)
	if err != nil {
		return
	}
	delete(data, "gwID")
	delete(data, "timestamp")

	dataJSON, err := json.Marshal(data)
	var ccDataLog deremsmodels.CCDataLog
	if err = json.Unmarshal(dataJSON, &ccDataLog); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}

	ccDataLog.GWUUID = gwIDValue.(string)
	ccDataLog.LogDate = time.Unix(int64(timestampValue.(float64)), 0)
	gateway, err := h.repo.Gateway.GetGatewayByGatewayUUID(gwIDValue.(string))
	if err == nil {
		ccDataLog.GWID = null.NewInt64(gateway.ID, true)
		ccDataLog.LocationID = gateway.LocationID
	} else {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Warn()
	}

	log.Debug("upsert local CC data log")
	err = h.repo.CCData.UpsertCCDataLog(&ccDataLog)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "h.repo.CCData.UpsertCCDataLog",
			"err":       err,
		}).Error()
	}
	return
}
