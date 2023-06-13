package routers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// GetGroups godoc
// @Summary List groups
// @Description list groups based on token
// @Tags        account management group
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetGroupsResponse}
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /account-management/groups [get]
func (w *APIWorker) GetGroups(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	responseData, err := w.Services.AccountManagement.GetGroups(userID.(int64))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrAccountGroupsGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// CreateGroup godoc
// @Summary Create a group
// @Description create a group by token
// @Tags        account management group
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Accept      json
// @Param       name           body      string true "Name"
// @Param       typeID         body      int true "TypeID"
// @Param       parentID       body      int true "ParentID"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /account-management/groups [post]
func (w *APIWorker) CreateGroup(c *gin.Context, body *app.CreateGroupBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if err := w.Services.AccountManagement.CreateGroup(body); err != nil {
		var code int
		if errors.Is(err, e.ErrNewAccountGroupNameOnSameLevelExist) {
			code = e.ErrAccountGroupNameOnSameLevelExist
		} else {
			code = e.ErrAccountGroupCreate
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
	logrus.WithFields(logrus.Fields{
		"created-by": userID,
		"name":       body.Name,
	}).Info("group-created")
}

// GetGroup godoc
// @Summary Show a group
// @Description get a group by token and group id
// @Tags        account management group
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       groupid        path      string true "Group ID"
// @Produce     json
// @Success     200            {object}  app.Response{data=services.GetGroupResponse}
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /api/account-management/groups/{groupid} [get]
func (w *APIWorker) GetGroup(c *gin.Context, uri *app.GroupURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	responseData, err := w.Services.AccountManagement.GetGroup(userID.(int64), uri.GroupID)
	if err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}
		appG.Response(http.StatusInternalServerError, e.ErrAccountGroupGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, responseData)
}

// UpdateGroup godoc
// @Summary Update a group
// @Description update a group by token and group id
// @Tags        account management group
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       groupid        path      string true "Group ID"
// @Accept      json
// @Param       name           body      string true "Name"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /api/account-management/groups/{groupid} [put]
func (w *APIWorker) UpdateGroup(c *gin.Context, uri *app.GroupURI, body *app.UpdateGroupBody) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if err := w.Services.AccountManagement.UpdateGroup(userID.(int64), uri.GroupID, body); err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}

		var code int
		switch err {
		case e.ErrNewAccountGroupNameOnSameLevelExist:
			code = e.ErrAccountGroupNameOnSameLevelExist
		case e.ErrNewOwnAccountGroupModifiedNotAllow:
			code = e.ErrOwnAccountGroupModifiedNotAllow
		default:
			code = e.ErrAccountGroupUpdate
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
	logrus.WithFields(logrus.Fields{
		"updated-by":       userID,
	}).Info("group-updated")
}

// DeleteGroup godoc
// @Summary Delete a group
// @Description delete a group by token and group id
// @Tags        account management group
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Param       groupid        path      string true "Group ID"
// @Produce     json
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /api/account-management/groups/{groupid} [delete]
func (w *APIWorker) DeleteGroup(c *gin.Context, uri *app.GroupURI) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if err := w.Services.AccountManagement.DeleteGroup(userID.(int64), uri.GroupID); err != nil {
		if errors.Is(err, e.ErrNewAuthPermissionNotAllow) {
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			return
		}

		var code int
		switch err {
		case e.ErrNewOwnAccountGroupModifiedNotAllow:
			code = e.ErrOwnAccountGroupModifiedNotAllow
		case e.ErrNewAccountGroupHasSubGroup:
			code = e.ErrAccountGroupHasSubGroup
		case e.ErrNewAccountGroupHasUser:
			code = e.ErrAccountGroupHasUser
		default:
			code = e.ErrAccountGroupDelete
		}
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}
