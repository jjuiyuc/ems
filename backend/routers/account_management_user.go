package routers

import (
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
