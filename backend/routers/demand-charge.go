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
