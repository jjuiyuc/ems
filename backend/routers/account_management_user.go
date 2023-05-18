package routers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// GetUsers godoc
func (w *APIWorker) GetUsers(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if userID == nil {
		logrus.WithField("caused-by", "error token").Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	responseData, err := w.Services.AccountManagement.GetUsers(userID.(int64))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrAccountUsersGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// CreateUser godoc
func (w *APIWorker) CreateUser(c *gin.Context, body *app.CreateUserBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	err := w.Services.AccountManagement.CreateUser(userID.(int64), body)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}

		var code int
		switch err {
		case e.ErrNewAccountUsernameExist:
			code = e.ErrAccountUsernameExist
		default:
			code = e.ErrAccountUserCreate
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}

// UpdateUser godoc
func (w *APIWorker) UpdateUser(c *gin.Context, uri *app.UserURI, body *app.UpdateUserBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	err := w.Services.AccountManagement.UpdateUser(userID.(int64), uri.UserID, body)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrAccountUserUpdate, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}
