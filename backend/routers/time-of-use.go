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
// @Param       query          query     StartTimeQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.BatteryUsageInfoResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Router      /{gwid}/devices/battery/usage-info [get]
func (w *APIWorker) GetBatteryUsageInfo(c *gin.Context) {
	appG := app.Gin{c}
	param := &StartTimeParam{}
	if err := param.validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetBatteryUsageInfo(param.GatewayUUID, param.Query.StartTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}
