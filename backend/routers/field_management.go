package routers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/kafka"
	"der-ems/services"
)

// GetFields godoc
// @Summary List fields
// @Description list fields based on token
// @Tags        field management
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetFieldsResponse}
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways [get]
func (w *APIWorker) GetFields(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	responseData, err := w.Services.FieldManagement.GetFields(userID.(int64))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrFieldsGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetDeviceModels godoc
// @Summary List device models
// @Description list device models based on token
// @Tags        field management
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetDeviceModelsResponse}
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/devices/models [get]
func (w *APIWorker) GetDeviceModels(c *gin.Context) {
	appG := app.Gin{c}
	responseData, err := w.Services.FieldManagement.GetDeviceModels()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrDeviceModelsGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// GetField godoc
// @Summary Show a field
// @Description get a field by token and gateway UUID
// @Tags        field management
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetFieldResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid} [get]
func (w *APIWorker) GetField(c *gin.Context, uri *app.FieldURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	responseData, err := w.Services.FieldManagement.GetField(userID.(int64), uri.GatewayID)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrFieldGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// EnableField godoc
// @Summary Enable/Disable a field
// @Description enable/disable a field by token and gateway UUID
// @Tags        field management
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Accept      json
// @Param       enable         body      bool true "Enable"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid}/field-state [put]
func (w *APIWorker) EnableField(c *gin.Context, uri *app.FieldURI, body *app.EnableFieldBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if err := w.Services.FieldManagement.EnableField(userID.(int64), uri.GatewayID, *body.Enable); err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrFieldEnable, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
	logrus.WithFields(logrus.Fields{"enabled-by": userID}).Info("field-enabled")
}

// SyncDeviceSettings godoc
// @Summary Sync device settings
// @Description sync a field device settings by token and gateway id
// @Tags        field management
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid}/sync-device-settings [get]
func (w *APIWorker) SyncDeviceSettings(c *gin.Context, uri *app.FieldURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	deviceSettings, err := w.Services.FieldManagement.GenerateDeviceSettings(userID.(int64), uri.GatewayID)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}

		var code int
		switch err {
		case e.ErrNewFieldIsDisabled:
			code = e.ErrFieldIsDisabled
		default:
			code = e.ErrDeviceSettingsSync
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	w.sendDeviceSettings(deviceSettings)
	appG.Response(http.StatusOK, e.Success, nil)
	logrus.WithFields(logrus.Fields{"sync-by": userID}).Info("device-settings-sync")
}

func (w *APIWorker) sendDeviceSettings(deviceSettings *services.DeviceSettingsData) {
	kafka.SendDataToGateways(w.Cfg, kafka.SendWeatherDataToLocalGW, deviceSettings.WeatherData, []string{deviceSettings.GWUUID})
	kafka.SendDataToGateways(w.Cfg, kafka.SendAIBillingParamsToLocalGW, deviceSettings.BillingData, []string{deviceSettings.GWUUID})
	kafka.SendDataToGateways(w.Cfg, kafka.SendDeviceMappingToLocalGW, deviceSettings.DeviceMappingData, []string{deviceSettings.GWUUID})
	kafka.SendDataToAIServer(w.Cfg, kafka.SendGPSLocation, deviceSettings.LocationData)
}

// UpdateFieldGroups godoc
// @Summary Update a field groups
// @Description update a field groups by token and gateway UUID
// @Tags        field management
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Accept      json
// @Param       groups         body      array  true "Groups"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid}/account-groups [put]
func (w *APIWorker) UpdateFieldGroups(c *gin.Context, uri *app.FieldURI, body *app.UpdateFieldGroupsBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	for _, group := range body.Groups {
		logrus.WithFields(logrus.Fields{
			"updated-by":    userID,
			"body-group-id": group.ID,
			"body-check":    *group.Check,
		}).Info("field-groups-updated")
	}
	if err := w.Services.FieldManagement.UpdateFieldGroups(userID.(int64), uri.GatewayID, body.Groups); err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}

		var code int
		switch err {
		case e.ErrNewOwnAccountGroupModifiedNotAllow:
			code = e.ErrOwnAccountGroupModifiedNotAllow
		default:
			code = e.ErrFieldGroupsUpdate
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}
