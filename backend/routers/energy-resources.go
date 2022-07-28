package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// GetBatteryEnergyInfo provides the detailed information and current state about a battery
// @Summary Provide detailed information and current state about a battery
// @Tags Energy Resources
// @Security ApiKeyAuth
// @Param Authorization header string true "Input user's access token" default(Bearer <Add access token here>)
// @Produce application/json
// @Success 200 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /{gwid}/devices/battery/energy-info [get]
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

	batteryEnergyInfo := w.Services.Battery.GetBatteryEnergyInfo(gatewayUUID)
	appG.Response(http.StatusOK, e.Success, batteryEnergyInfo)
}
