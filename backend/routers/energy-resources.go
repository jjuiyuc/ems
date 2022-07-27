package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
)

// GetBatteryEnergyInfo godoc
// @Summary     Show the detailed information and current state about a battery
// @Description get battery by token, gateway UUID and startTime
// @Tags        energy resources
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       startTime      query     string true "UTC time in ISO-8601" format(date-time)
// @Produce     json
// @Success     200            {object}  app.Response{data=services.BatteryEnergyInfoResponse}
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/battery/energy-info [get]
func (w *APIWorker) GetBatteryEnergyInfo(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if userID == nil {
		log.WithFields(log.Fields{"caused-by": "error token"}).Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	gatewayUUID := c.Param("gwid")
	log.Debug("gatewayUUID: ", gatewayUUID)

	batteryEnergyInfo := w.Services.Battery.GetBatteryEnergyInfo(gatewayUUID)
	appG.Response(http.StatusOK, e.Success, batteryEnergyInfo)
}

// GetBatteryPowerState provides today's hourly power state of a battery
// @Summary Provide today's hourly power state of a battery
// @Tags Energy Resources
// @Security ApiKeyAuth
// @Param Authorization header string true "Input user's access token" default(Bearer <Add access token here>)
// @Produce application/json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /{gwid}/devices/battery/power-state [get]
func (w *APIWorker) GetBatteryPowerState(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if userID == nil {
		log.WithFields(log.Fields{"caused-by": "error token"}).Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	gatewayUUID := c.Param("gwid")
	log.Debug("gatewayUUID: ", gatewayUUID)

	startTime, endTime, ok := w.checkZoomableParams(c)
	if !ok {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	batteryPowerState, err := w.Services.Battery.GetBatteryPowerState(gatewayUUID, startTime, endTime)
	if err != nil {
		log.WithFields(log.Fields{"caused-by": "generate battery power state"}).Error()
		appG.Response(http.StatusInternalServerError, e.ErrBatteryPowerStateGen, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.Success, batteryPowerState)
}

func (w *APIWorker) checkZoomableParams(c *gin.Context) (startTime, endTime time.Time, ok bool) {
	// TODO: Only supports hour now
	if c.Query("resolution") != "hour" {
		return
	}
	startTime, ok = utils.IsDateFormat(time.RFC3339, c.Query("startTime"))
	if !ok {
		return
	}
	endTime, ok = utils.IsDateFormat(time.RFC3339, c.Query("endTime"))
	if !ok || startTime == endTime || endTime.Before(startTime) {
		ok = false
		return
	}
	ok = true
	return
}
