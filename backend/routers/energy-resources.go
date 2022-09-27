package routers

import (
	"github.com/gin-gonic/gin"
)

// GetSolarEnergyInfo godoc
// @Summary     Show the detailed information and current state about solar
// @Description get solar by token, gateway UUID and startTime
// @Tags        energy resources
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     StartTimeQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.SolarEnergyInfoResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/solar/energy-info [get]
func (w *APIWorker) GetSolarEnergyInfo(c *gin.Context) {
	w.getResponseByStartTimeAPIType(c, SolarEnergyInfo)
}

// GetSolarPowerState godoc
// @Summary     Show today's hourly power state of solar
// @Description get solar by token, gateway UUID, resolution, startTime and endTime
// @Tags        energy resources
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     ZoomableQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.SolarPowerStateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /{gwid}/devices/solar/power-state [get]
func (w *APIWorker) GetSolarPowerState(c *gin.Context) {
	w.getZoomableInfo(c, SolarPowerState)
}

// GetBatteryEnergyInfo godoc
// @Summary     Show the detailed information and current state about battery
// @Description get battery by token, gateway UUID and startTime
// @Tags        energy resources
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     StartTimeQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.BatteryEnergyInfoResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/battery/energy-info [get]
func (w *APIWorker) GetBatteryEnergyInfo(c *gin.Context) {
	appG := app.Gin{c}
	param := &StartTimeParam{}
	if err := param.validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Battery.GetBatteryEnergyInfo(param.GatewayUUID, param.Query.StartTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetBatteryPowerState godoc
// @Summary     Show today's hourly power state of battery
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
	w.getZoomableInfo(c, BatteryPowerState)
}

// GetBatteryChargeVoltageState godoc
// @Summary     Show today's hourly charge and voltage state of battery
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
	w.getZoomableInfo(c, BatteryChargeVoltageState)
}
