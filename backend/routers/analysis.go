package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// PeriodQuery godoc
type PeriodQuery struct {
	StartTime time.Time `form:"startTime" binding:"required" example:"UTC time in ISO-8601" format:"date-time"`
	EndTime   time.Time `form:"endTime" binding:"required,gtfield=StartTime" example:"UTC time in ISO-8601" format:"date-time"`
}

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
	userID, _ := c.Get("userID")
	if userID == nil {
		log.WithFields(log.Fields{"caused-by": "error token"}).Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	gatewayUUID := c.Param("gwid")
	log.Debug("gatewayUUID: ", gatewayUUID)

	var q PeriodQuery
	if err := c.BindQuery(&q); err != nil {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	energyDistributionInfo := w.Services.Devices.GetEnergyDistributionInfo(gatewayUUID, q.StartTime, q.EndTime)
	appG.Response(http.StatusOK, e.Success, energyDistributionInfo)
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
	// TODO: Only supports hour now
	if err := c.BindQuery(&q); err != nil || q.Resolution != "hour" {
		log.WithFields(log.Fields{"caused-by": "invalid param"}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	powerState := w.Services.Devices.GetPowerState(gatewayUUID, q.StartTime, q.EndTime)
	appG.Response(http.StatusOK, e.Success, powerState)
}
