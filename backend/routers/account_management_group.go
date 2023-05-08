package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// GetGroups godoc
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
		appG.Response(http.StatusInternalServerError, errCode, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}
