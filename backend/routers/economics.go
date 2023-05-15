package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// GetTimeOfUseEnergyCost godoc
// @Summary     Show energy cost of Pre-Ubiik and Post-Ubiik
// @Description get energy cost by token, gateway UUID, startTime and endTime
// @Tags        economics
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Param       query          query     app.PeriodQuery true "Query"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.TimeOfUseEnergyCostResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Router      /{gwid}/devices/tou/energy-cost [get]
func (w *APIWorker) GetTimeOfUseEnergyCost(c *gin.Context) {
	appG := app.Gin{c}
	param := &app.PeriodParam{}
	if err := param.Validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	responseData := w.Services.Devices.GetTimeOfUseEnergyCost(param)
	appG.Response(http.StatusOK, e.Success, responseData)
}
