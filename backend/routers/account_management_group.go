package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

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
	if userID == nil {
		log.WithField("caused-by", "error token").Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	responseData, err := w.Services.AccountManagement.GetGroups(userID.(int64))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrAccountGroupsGen, err.Error())
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
// @Produce     json
// @Param       name           body      string true "Name"
// @Param       typeID         body      int true "TypeID"
// @Param       parentID       body      int true "ParentID"
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     403            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /account-management/groups [post]
func (w *APIWorker) CreateGroup(c *gin.Context) {
	appG := app.Gin{c}
	body := &app.CreateGroupBody{}
	if err := body.Validate(c); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	errCode, err := w.Services.AccountManagement.CreateGroup(body)
	if err != nil {
		if errCode != e.ErrAccountGroupNameOnSameLevelExist {
			errCode = e.ErrorAccountGroupCreate
		}
		appG.Response(http.StatusInternalServerError, errCode, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}
