package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// GetBatteryUsageInfo godoc
// @Summary     Show current usage state about battery
// @Description get battery by token, gateway UUID and startTime
// @Tags        time of use
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     app.StartTimeQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.BatteryUsageInfoResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Router      /{gwid}/devices/battery/usage-info [get]
func (w *APIWorker) GetBatteryUsageInfo(c *gin.Context) {
	appG := app.Gin{c}
	param := &app.StartTimeParam{}
	if err := param.Validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetBatteryUsageInfo(param)
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetTimeOfUseInfo godoc
// @Summary     Show the distribution of energy sources for peak types of the day
// @Description get energy source distribution by token, gateway UUID, startTime
// @Tags        time of use
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     app.StartTimeQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.TimeOfUseInfoResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /{gwid}/devices/tou/info [get]
func (w *APIWorker) GetTimeOfUseInfo(c *gin.Context) {
	appG := app.Gin{c}
	param := &app.StartTimeParam{}
	if err := param.Validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData, err := w.Services.Devices.GetTimeOfUseInfo(param)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrTimeOfUseInfoGen, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetSolarEnergyUsage godoc
// @Summary     Show the day's hourly energy usage of solar
// @Description get solar by token, gateway UUID, resolution, startTime and endTime
// @Tags        time of use
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     app.ZoomableQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.SolarEnergyUsageResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Router      /{gwid}/devices/solar/energy-usage [get]
func (w *APIWorker) GetSolarEnergyUsage(c *gin.Context) {
	appG := app.Gin{c}
	param := &app.ZoomableParam{}
	if err := param.Validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetSolarEnergyUsage(param)
	appG.Response(http.StatusOK, e.Success, responseData)
}
