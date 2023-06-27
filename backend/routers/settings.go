package routers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
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
	if err := w.Services.Settings.UpdateBatterySettings(userID.(int64), uri.GatewayID, body); err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrBatterySettingsUpdate, nil)
		return
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
