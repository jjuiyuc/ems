package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// GetEnergyDistributionInfo godoc
// @Summary     Show the distribution of energy sources and distinations
// @Description get energy distribution by token, gateway UUID, startTime and endTime
// @Tags        analysis
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     app.PeriodQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.EnergyDistributionInfoResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/energy-distribution-info [get]
func (w *APIWorker) GetEnergyDistributionInfo(c *gin.Context) {
	appG := app.Gin{c}
	param := &app.PeriodParam{}
	if err := param.Validate(c); err != nil {
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
// @Param       query          query     app.ZoomableQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.PowerStateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/power-state [get]
func (w *APIWorker) GetPowerState(c *gin.Context) {
	appG := app.Gin{c}
	param := &app.ZoomableParam{}
	if err := param.Validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetPowerState(param.GatewayUUID, param.Query.Resolution, param.Query.StartTime, param.Query.EndTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetAccumulatedPowerState godoc
// @Summary     Show daily/monthly accumulated power state of Load/Solar/Battery/Grid
// @Description get power state by token, gateway UUID, resolution, startTime and endTime
// @Tags        analysis
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     app.ResolutionWithPeriodQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.AccumulatedPowerStateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/accumulated-power-state [get]
func (w *APIWorker) GetAccumulatedPowerState(c *gin.Context) {
	appG := app.Gin{c}
	param := &app.ResolutionWithPeriodParam{}
	if err := param.Validate(c); err != nil {
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
// @Param       query          query     app.ResolutionWithPeriodQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.PowerSelfSupplyRateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/power-self-supply-rate [get]
func (w *APIWorker) GetPowerSelfSupplyRate(c *gin.Context) {
	appG := app.Gin{c}
	param := &app.ResolutionWithPeriodParam{}
	if err := param.Validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetPowerSelfSupplyRate(param.GatewayUUID, param.Query.Resolution, param.Query.StartTime, param.Query.EndTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}
