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
// @Summary List users
// @Description list users based on token
// @Tags        account management user
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetUsersResponse}
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /account-management/users [get]
func (w *APIWorker) GetUsers(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	responseData, err := w.Services.AccountManagement.GetUsers(userID.(int64))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrAccountUsersGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// CreateUser godoc
// @Summary Create a user
// @Description create a user by token
// @Tags        account management user
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Accept      json
// @Param       username       body      string true "Username"
// @Param       password       body      string true "Password"
// @Param       name           body      string true "Name"
// @Param       groupID        body      int true "GroupID"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /account-management/users [post]
func (w *APIWorker) CreateUser(c *gin.Context, body *app.CreateUserBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if err := w.Services.AccountManagement.CreateUser(userID.(int64), body); err != nil {
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
	logrus.WithFields(logrus.Fields{
		"created-by":       userID,
		"username":         body.Username,
	}).Info("user-created")
}

// UpdateUser godoc
// @Summary Update a user
// @Description update a user by token and user id
// @Tags        account management user
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       userid         path      string true "User ID"
// @Accept      json
// @Param       password       body      string true "Password"
// @Param       name           body      string true "Name"
// @Param       groupID        body      int true "GroupID"
// @Param       unlock         body      bool true "Unlock"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /account-management/users/{user-id} [put]
func (w *APIWorker) UpdateUser(c *gin.Context, uri *app.UserURI, body *app.UpdateUserBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if err := w.Services.AccountManagement.UpdateUser(userID.(int64), uri.UserID, body); err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrAccountUserUpdate, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
	logrus.WithFields(logrus.Fields{
		"updated-by":       userID,
	}).Info("user-updated")
}

// DeleteUser godoc
// @Summary Delete a user
// @Description delete a user by token and user id
// @Tags        account management user
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       userid         path      string true "User ID"
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /account-management/users/{user-id} [delete]
func (w *APIWorker) DeleteUser(c *gin.Context, uri *app.UserURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if err := w.Services.AccountManagement.DeleteUser(userID.(int64), uri.UserID); err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}

		var code int
		if errors.Is(err, e.ErrNewOwnAccountDeletedNotAllow) {
			code = e.ErrOwnAccountDeletedNotAllow
		} else {
			code = e.ErrAccountUserDelete
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}