package routers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/kafka"
)

func (w *APIWorker) GetBatterySettings(c *gin.Context, uri *app.FieldURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	responseData, err := w.Services.Settings.GetBatterySettings(userID.(int64), uri.GatewayID)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrBatterySettingsGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

func (w *APIWorker) UpdateBatterySettings(c *gin.Context, uri *app.FieldURI, body *app.UpdateBatterySettingsBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	logrus.WithFields(logrus.Fields{
		"updated-by":                       userID,
		"charging-sources":                 body.ChargingSources,
		"reserved-for-grid-outage-percent": body.ReservedForGridOutagePercent,
	}).Info("battery-settings-updated")
	dlData, err := w.Services.Settings.UpdateBatterySettings(userID.(int64), uri.GatewayID, body)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrBatterySettingsUpdate, nil)
		return
	}
	if dlData != nil {
		w.sendAISystemParam(dlData, uri.GatewayID)
	}
	appG.Response(http.StatusOK, e.Success, nil)
}

func (w *APIWorker) GetMeterSettings(c *gin.Context, uri *app.FieldURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	responseData, err := w.Services.Settings.GetMeterSettings(userID.(int64), uri.GatewayID)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrMeterSettingsGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

func (w *APIWorker) UpdateMeterSettings(c *gin.Context, uri *app.FieldURI, body *app.UpdateMeterSettingsBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	logrus.WithFields(logrus.Fields{
		"updated-by":          userID,
		"max-demand-capacity": body.MaxDemandCapacity,
	}).Info("meter-settings-updated")
	dlData, err := w.Services.Settings.UpdateMeterSettings(userID.(int64), uri.GatewayID, body)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrMeterSettingsUpdate, nil)
		return
	}
	if dlData != nil {
		w.sendAISystemParam(dlData, uri.GatewayID)
	}
	appG.Response(http.StatusOK, e.Success, nil)
}

func (w *APIWorker) sendAISystemParam(data []byte, gatewayUUID string) {
	kafka.SendDataToGateways(w.Cfg, kafka.SendAISystemParamToLocalGW, data, []string{gatewayUUID})
}
