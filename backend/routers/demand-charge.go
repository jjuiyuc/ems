package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// GetChargeInfo godoc
// @Summary     Show the demand charge information
// @Description get demand charge information by token, gateway UUID and startTime
// @Tags        demand charge
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     StartTimeQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.ChargeInfoResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/charge-info [get]
func (w *APIWorker) GetChargeInfo(c *gin.Context) {
	appG := app.Gin{c}
	param := &StartTimeParam{}
	if err := param.validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetChargeInfo(param.GatewayUUID, param.Query.StartTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetDemandState godoc
// @Summary     Show today's demand details
// @Description get demand details by token, gateway UUID, startTime and endTime
// @Tags        demand charge
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     PeriodQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.DemandStateResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/demand-state [get]
func (w *APIWorker) GetDemandState(c *gin.Context) {
	appG := app.Gin{c}
	param := &PeriodParam{}
	if err := param.validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetDemandState(param.GatewayUUID, param.Query.StartTime, param.Query.EndTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}
