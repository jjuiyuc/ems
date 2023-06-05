package routers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

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

func (w *APIWorker) GetDeviceModels(c *gin.Context) {
	appG := app.Gin{c}
	responseData, err := w.Services.FieldManagement.GetDeviceModels()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrDeviceModelsGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

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
