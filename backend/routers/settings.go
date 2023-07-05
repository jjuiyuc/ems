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

// @Summary Show battery settings
// @Description get battery settings by token and gateway UUID
// @Tags        settings
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetBatterySettingsResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid}/battery-settings [get]
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

// @Summary Update battery settings
// @Description update battery settings by token and gateway UUID
// @Tags        settings
// @Security    ApiKeyAuth
// @Param       Authorization                 header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid                          path      string true "Gateway UUID"
// @Accept      json
// @Param       chargingSources               body      string true "Charging sources"
// @Param       reservedForGridOutagePercent  body      int    true "Reserved for grid outage percent"
// @Produce     json
// @Success     200                           {object}  app.Response
// @Failure     400                           {object}  app.Response
// @Failure     401                           {object}  app.Response
// @Failure     403                           {object}  app.Response
// @Failure     500                           {object}  app.Response
// @Router      /device-management/gateways/{gwid}/battery-settings [put]
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

		var code int
		switch err {
		case e.ErrNewFieldIsDisabled:
			code = e.ErrFieldIsDisabled
		default:
			code = e.ErrBatterySettingsUpdate
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	if dlData != nil {
		w.sendAISystemParam(dlData, uri.GatewayID)
	}
	appG.Response(http.StatusOK, e.Success, nil)
}

// @Summary Show maximum demand capacity
// @Description get meter settings by token and gateway UUID
// @Tags        settings
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetMeterSettingsResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid}/meter-settings [get]
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

// @Summary Update maximum demand capacity
// @Description update meter settings by token and gateway UUID
// @Tags        settings
// @Security    ApiKeyAuth
// @Param       Authorization      header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid               path      string true "Gateway UUID"
// @Accept      json
// @Param       maxDemandCapacity  body      int    true "Maximum demand capacity"
// @Produce     json
// @Success     200                {object}  app.Response
// @Failure     400                {object}  app.Response
// @Failure     401                {object}  app.Response
// @Failure     403                {object}  app.Response
// @Failure     500                {object}  app.Response
// @Router      /device-management/gateways/{gwid}/meter-settings [put]
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

		var code int
		switch err {
		case e.ErrNewFieldIsDisabled:
			code = e.ErrFieldIsDisabled
		default:
			code = e.ErrMeterSettingsUpdate
		}
		appG.Response(http.StatusInternalServerError, code, nil)
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

// @Summary Show power outage periods
// @Description get power outage periods by token and gateway UUID
// @Tags        settings
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       gwid           path      string true "Gateway UUID"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetPowerOutagePeriodsResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid}/power-outage-periods [get]
func (w *APIWorker) GetPowerOutagePeriods(c *gin.Context, uri *app.FieldURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	responseData, err := w.Services.Settings.GetPowerOutagePeriods(userID.(int64), uri.GatewayID)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrPowerOutagePeriodsGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// @Summary Create power outage periods
// @Description create power outage periods by token and gateway UUID
// @Tags        settings
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Accept      json
// @Param       periods        body      array  true "Periods"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid}/power-outage-periods [post]
func (w *APIWorker) CreatePowerOutagePeriods(c *gin.Context, uri *app.FieldURI, body *app.CreatePowerOutagePeriodsBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	for _, period := range body.Periods {
		logrus.WithFields(logrus.Fields{
			"created-by": userID,
			"period":     period,
		}).Info("power-outage-created")
	}
	dlData, err := w.Services.Settings.CreatePowerOutagePeriods(userID.(int64), uri.GatewayID, body)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}

		var code int
		switch err {
		case e.ErrNewFieldIsDisabled:
			code = e.ErrFieldIsDisabled
		case e.ErrNewPowerOutagePeriodsMoreThanMaximum:
			code = e.ErrPowerOutagePeriodsMoreThanMaximum
		case e.ErrNewPowerOutagePeriodInvalid:
			code = e.ErrPowerOutagePeriodInvalid
		default:
			code = e.ErrPowerOutagePeriodsCreate
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	if dlData != nil {
		kafka.SendDataToGateways(w.Cfg, kafka.SendAINotificationToLocalGW, dlData, []string{uri.GatewayID})
	}
	appG.Response(http.StatusOK, e.Success, nil)
}

// @Summary Delete a power outage period
// @Description delete a power outage period by token, gateway UUID and period ID
// @Tags        settings
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       periodid       path      string true "Period ID"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /device-management/gateways/{gwid}/power-outage-periods/{periodid} [delete]
func (w *APIWorker) DeletePowerOutagePeriod(c *gin.Context, uri *app.GatewayAndPeriodURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	dlData, err := w.Services.Settings.DeletePowerOutagePeriod(userID.(int64), uri.GatewayID, uri.PeriodID)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}

		var code int
		switch err {
		case e.ErrNewFieldIsDisabled:
			code = e.ErrFieldIsDisabled
		case e.ErrNewPowerOutagePeriodOngoing:
			code = e.ErrPowerOutagePeriodOngoing
		default:
			code = e.ErrPowerOutagePeriodDelete
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	if dlData != nil {
		kafka.SendDataToGateways(w.Cfg, kafka.SendAINotificationToLocalGW, dlData, []string{uri.GatewayID})
	}
	appG.Response(http.StatusOK, e.Success, nil)
}
