package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// StartTimeQuery godoc
type StartTimeQuery struct {
	StartTime time.Time `form:"startTime" binding:"required" example:"UTC time in ISO-8601" format:"date-time"`
}

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
	gatewayUUID := c.Param("gwid")
	log.Debug("gatewayUUID: ", gatewayUUID)

	var q StartTimeQuery
	if err := c.BindQuery(&q); err != nil {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	responseData := w.Services.Devices.GetChargeInfo(gatewayUUID, q.StartTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetDemandState godoc
func (w *APIWorker) GetDemandState(c *gin.Context) {
	w.getPeriodInfo(c)
}

func (w *APIWorker) getPeriodInfo(c *gin.Context) {
	appG := app.Gin{c}
	gatewayUUID := c.Param("gwid")
	log.Debug("gatewayUUID: ", gatewayUUID)

	var q PeriodQuery
	if err := c.BindQuery(&q); err != nil {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	responseData := w.Services.Devices.GetDemandState(gatewayUUID, q.StartTime, q.EndTime)
	appG.Response(http.StatusOK, e.Success, responseData)
}
