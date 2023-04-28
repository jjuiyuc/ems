package utils

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"der-ems/internal/e"
)

const (
	gwID      = "gwID"
	timestamp = "timestamp"
	msgType   = "type"
)

// AssertGatewayMessage godoc
func AssertGatewayMessage(msg []byte) (gwIDValue, timestampValue interface{}, data map[string]interface{}, err error) {
	if err = json.Unmarshal(msg, &data); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "json.Unmarshal",
			"err":       err,
		}).Error()
		return
	}

	gwIDValue, ok := data[gwID]
	if !ok {
		err = e.ErrNewKeyNotExist(gwID)
		log.WithFields(log.Fields{
			"caused-by": gwID,
			"err":       err,
		}).Error()
		return
	}
	if _, ok = gwIDValue.(string); !ok {
		err = e.ErrNewKeyUnexpectedValue(gwID)
		log.WithFields(log.Fields{
			"caused-by": gwID,
			"err":       err,
		}).Error()
		return
	}
	timestampValue, ok = data[timestamp]
	if !ok {
		err = e.ErrNewKeyNotExist(timestamp)
		log.WithFields(log.Fields{
			"caused-by": timestamp,
			"err":       err,
		}).Error()
		return
	}
	if _, ok = timestampValue.(float64); !ok {
		err = e.ErrNewKeyUnexpectedValue(timestamp)
		log.WithFields(log.Fields{
			"caused-by": timestamp,
			"err":       err,
		}).Error()
	}
	return
}

// AssertGatewayMessageType godoc
func AssertGatewayMessageType(data map[string]interface{}) (typeValue interface{}, err error) {
	typeValue, ok := data[msgType]
	if !ok {
		err = e.ErrNewKeyNotExist(msgType)
		log.WithFields(log.Fields{
			"caused-by": msgType,
			"err":       err,
		}).Error()
		return
	}
	if _, ok = typeValue.(string); !ok {
		err = e.ErrNewKeyUnexpectedValue(msgType)
		log.WithFields(log.Fields{
			"caused-by": msgType,
			"err":       err,
		}).Error()
	}
	return
}
