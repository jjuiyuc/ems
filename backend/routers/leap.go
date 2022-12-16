package routers

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// LeapNotification godoc
type LeapNotification struct {
	TestNotification bool                     `json:"test_notification"`
	MeterDispatches  []map[string]interface{} `json:"meter_dispatches"`
}

// GetLeapBiddingDispatch godoc
func (w *APIWorker) GetLeapBiddingDispatch(c *gin.Context) {
	// XXX: Hardcode for demo
	const MeterID = "436ca983-d52e-4d8e-8c82-2d5021d495bf"
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
		// TODO: Send notification content to gateway local AI
		log.Debug("Send notification content to gateway")
	}
}
