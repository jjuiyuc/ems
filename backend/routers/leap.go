package routers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/kafka"
)

// LeapNotification godoc
type LeapNotification struct {
	TestNotification bool                     `json:"test_notification"`
	MeterDispatches  []map[string]interface{} `json:"meter_dispatches"`
}

// GetLeapBiddingDispatch godoc
func (w *APIWorker) GetLeapBiddingDispatch(c *gin.Context) {
	// XXX: Hardcode for demo
	const (
		MeterID     = "436ca983-d52e-4d8e-8c82-2d5021d495bf"
		GatewayUUID = "0E0BA27A8175AF978C49396BDE9D7A1E"
	)
	appG := app.Gin{c}
	appG.Response(http.StatusOK, e.Success, nil)

	body := new(LeapNotification)
	if err := c.BindJSON(&body); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "c.BindJSON",
			"err":       err,
		}).Error()
		return
	}

	meterExists := false
	for _, dispatch := range body.MeterDispatches {
		meterIDs, ok := dispatch["meter_ids"]
		if ok {
			// Market group notification
			if reflect.TypeOf(meterIDs).Kind() == reflect.Slice {
				s := reflect.ValueOf(meterIDs)
				for i := 0; i < s.Len(); i++ {
					meterID := (s.Index(i).Interface()).(string)
					if meterID == MeterID {
						meterExists = true
						break
					}
				}
			}
		} else {
			// Meter notification
			if dispatch["meter_id"] == MeterID {
				meterExists = true
				break
			}
		}
	}
	if meterExists {
		bodyJSON, _ := json.Marshal(body)
		log.Debug("bodyJSON : ", string(bodyJSON))
		sendLeapBiddingDispatchToGateway(w.Cfg, bodyJSON, GatewayUUID)
	}
}

func sendLeapBiddingDispatchToGateway(cfg *viper.Viper, biddingDispatchJSON []byte, uuid string) {
	biddingDispatchTopic := strings.Replace(kafka.SendLeapBiddingDispatchToLocalGW, "{gw-id}", uuid, 1)
	log.Debug("biddingDispatchTopic: ", biddingDispatchTopic)
	kafka.Produce(cfg, biddingDispatchTopic, string(biddingDispatchJSON))
}
