package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// ZoomableQuery godoc
type ZoomableQuery struct {
	Resolution string `form:"resolution" validate:"required" enums:"hour"`
	StartTime  string `form:"startTime" validate:"required" example:"UTC time in ISO-8601" format:"date-time"`
	EndTime    string `form:"endTime" validate:"required" example:"UTC time in ISO-8601" format:"date-time"`
}

// BatteryState godoc
type BatteryState int

const (
	// Power godoc
	Power BatteryState = iota
	// ChargeVoltage godoc
	ChargeVoltage
)

// GetBatteryEnergyInfo godoc
// @Summary     Show the detailed information and current state about a battery
// @Description get battery by token, gateway UUID and startTime
// @Tags        energy resources
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       startTime      query     string true "Example : UTC time in ISO-8601" format(date-time)
// @Produce     json
// @Success     200            {object}  app.Response{data=services.BatteryEnergyInfoResponse}
// @Failure     400            {object}  app.Response
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

	startTime, err := time.Parse(time.RFC3339, c.Query("startTime"))
	if err != nil {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	batteryEnergyInfo := w.Services.Battery.GetBatteryEnergyInfo(gatewayUUID, startTime)
	appG.Response(http.StatusOK, e.Success, batteryEnergyInfo)
}

// GetBatteryPowerState godoc
// @Summary     Show today's hourly power state of a battery
// @Description get battery by token, gateway UUID, resolution, startTime and endTime
// @Tags        energy resources
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     ZoomableQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.BatteryPowerStateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /{gwid}/devices/battery/power-state [get]
func (w *APIWorker) GetBatteryPowerState(c *gin.Context) {
	w.getBatteryState(c, Power)
}

// GetBatteryChargeVoltageState godoc
// @Summary     Show today's hourly charge and voltage state of a battery
// @Description get battery by token, gateway UUID, resolution, startTime and endTime
// @Tags        energy resources
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     ZoomableQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.BatteryChargeVoltageStateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router /{gwid}/devices/battery/charge-voltage-state [get]
func (w *APIWorker) GetBatteryChargeVoltageState(c *gin.Context) {
	w.getBatteryState(c, ChargeVoltage)
}

func (w *APIWorker) getBatteryState(c *gin.Context, batteryState BatteryState) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if userID == nil {
		log.WithFields(log.Fields{"caused-by": "error token"}).Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	gatewayUUID := c.Param("gwid")
	log.Debug("gatewayUUID: ", gatewayUUID)

	var q ZoomableQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	startTime, endTime, ok := q.Validate()
	if !ok {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	switch batteryState {
	case Power:
		batteryPowerState, err := w.Services.Battery.GetBatteryPowerState(gatewayUUID, startTime, endTime)
		if err != nil {
			log.WithFields(log.Fields{"caused-by": "generate battery power state"}).Error()
			appG.Response(http.StatusInternalServerError, e.ErrBatteryPowerStateGen, err.Error())
			return
		}
		appG.Response(http.StatusOK, e.Success, batteryPowerState)
	case ChargeVoltage:
		batteryChargeVoltageState, err := w.Services.Battery.GetBatteryChargeVoltageState(gatewayUUID, startTime, endTime)
		if err != nil {
			log.WithFields(log.Fields{"caused-by": "generate battery charge and voltage state"}).Error()
			appG.Response(http.StatusInternalServerError, e.ErrBatteryChargeVoltageStateGen, err.Error())
			return
		}
		appG.Response(http.StatusOK, e.Success, batteryChargeVoltageState)
	}
}

// Validate godoc
func (zoomableQuery *ZoomableQuery) Validate() (periodStartTime, periodEndTime time.Time, ok bool) {
	// TODO: Only supports hour now
	if zoomableQuery.Resolution != "hour" {
		return
	}
	startTime, err := time.Parse(time.RFC3339, zoomableQuery.StartTime)
	if err != nil {
		return
	}
	endTime, err := time.Parse(time.RFC3339, zoomableQuery.EndTime)
	if err != nil || startTime == endTime || endTime.Before(startTime) {
		return
	}
	periodStartTime, periodEndTime, err = zoomableQuery.getStatePeriod(startTime, endTime)
	if err != nil {
		return
	}
	ok = true
	return
}

func (zoomableQuery *ZoomableQuery) getStatePeriod(startTime, endTime time.Time) (periodStartTime, periodEndTime time.Time, err error) {
	periodStartTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), 0, 0, 0, startTime.Location())
	log.Debug("periodStartTime: ", periodStartTime)
	periodEndTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour(), 0, 0, 0, endTime.Location())
	log.Debug("periodEndTime: ", periodEndTime)
	if periodStartTime == periodEndTime {
		err = e.ErrNewUnexpectedTimeRange
		log.WithFields(log.Fields{
			"caused-by": "s.getStatePeriod",
			"err":       err,
		}).Error()
	}
	return
}
