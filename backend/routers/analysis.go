package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// ZoomableType godoc
type ZoomableType int

const (
	// PowerState godoc
	PowerState ZoomableType = iota
	// SolarPowerState godoc
	SolarPowerState
	// BatteryPowerState godoc
	BatteryPowerState
	// BatteryChargeVoltageState godoc
	BatteryChargeVoltageState
)

// GetEnergyDistributionInfo godoc
// @Summary     Show the distribution of energy sources and distinations
// @Description get energy distribution by token, gateway UUID, startTime and endTime
// @Tags        analysis
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     PeriodQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.EnergyDistributionInfoResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/energy-distribution-info [get]
func (w *APIWorker) GetEnergyDistributionInfo(c *gin.Context) {
	appG := app.Gin{c}
	param := &PeriodParam{}
	if err := param.validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetEnergyDistributionInfo(param.GatewayUUID, param.Query.StartTime, param.Query.EndTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetPowerState godoc
// @Summary     Show today's hourly power state of Load/Solar/Battery/Grid
// @Description get power state by token, gateway UUID, resolution, startTime and endTime
// @Tags        analysis
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     ZoomableQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.PowerStateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/power-state [get]
func (w *APIWorker) GetPowerState(c *gin.Context) {
	w.getZoomableInfo(c, PowerState)
}

// GetAccumulatedPowerState godoc
// @Summary     Show daily/monthly accumulated power state of Load/Solar/Battery/Grid
// @Description get power state by token, gateway UUID, resolution, startTime and endTime
// @Tags        analysis
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     ResolutionWithPeriodQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.AccumulatedPowerStateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/accumulated-power-state [get]
func (w *APIWorker) GetAccumulatedPowerState(c *gin.Context) {
	appG := app.Gin{c}
	param := &ResolutionWithPeriodParam{}
	if err := param.validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetAccumulatedPowerState(param.GatewayUUID, param.Query.Resolution, param.Query.StartTime, param.Query.EndTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetPowerSelfSupplyRate godoc
// @Summary     Show daily/monthly power self supply rate
// @Description get power self supply rate by token, gateway UUID, resolution, startTime and endTime
// @Tags        analysis
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     ResolutionWithPeriodQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.PowerSelfSupplyRateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/power-self-supply-rate [get]
func (w *APIWorker) GetPowerSelfSupplyRate(c *gin.Context) {
	appG := app.Gin{c}
	param := &ResolutionWithPeriodParam{}
	if err := param.validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetPowerSelfSupplyRate(param.GatewayUUID, param.Query.Resolution, param.Query.StartTime, param.Query.EndTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}

func (w *APIWorker) getZoomableInfo(c *gin.Context, zoomableType ZoomableType) {
	appG := app.Gin{c}
	gatewayUUID := c.Param("gwid")
	log.Debug("gatewayUUID: ", gatewayUUID)

	var q ZoomableQuery
	// TODO: Only supports hour now
	if err := c.BindQuery(&q); err != nil || q.Resolution != "hour" {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	var responseData interface{}
	var err error
	switch zoomableType {
	case PowerState:
		responseData = w.Services.Devices.GetPowerState(gatewayUUID, q.StartTime, q.EndTime)
	case SolarPowerState:
		responseData, err = w.Services.Devices.GetSolarPowerState(gatewayUUID, q.StartTime, q.EndTime)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ErrSolarPowerStateGen, err.Error())
			return
		}
	case BatteryPowerState:
		responseData, err = w.Services.Battery.GetBatteryPowerState(gatewayUUID, q.StartTime, q.EndTime)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ErrBatteryPowerStateGen, err.Error())
			return
		}
	case BatteryChargeVoltageState:
		responseData, err = w.Services.Battery.GetBatteryChargeVoltageState(gatewayUUID, q.StartTime, q.EndTime)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ErrBatteryChargeVoltageStateGen, err.Error())
			return
		}
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}
