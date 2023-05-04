package routers

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
)

// GetAuth godoc
// @Summary     Show a token
// @Description get token by username and password
// @Tags        authorization
// @Accept      json
// @Produce     json
// @Param       username  body      string true "Username"
// @Param       password  body      string true "Password"
// @Success     200       {object}  app.Response
// @Failure     400       {object}  app.Response
// @Failure     401       {object}  app.Response
// @Failure     500       {object}  app.Response
// @Router      /auth [post]
func (w *APIWorker) GetAuth(c *gin.Context) {
	appG := app.Gin{c}
	valid := validation.Validation{}

	var a struct {
		Username string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(50)"`
	}
	c.BindJSON(&a)
	if ok, err := valid.Valid(&a); !ok {
		log.WithFields(log.Fields{
			"caused-by": "valid.Valid",
			"err":       err,
		}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, valid.Errors)
		return
	}

	user, groupType, errCode, err := w.Services.Auth.Login(a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusUnauthorized, errCode, map[string]string{
			"msg": err.Error(),
		})
		return
	}

	claims := utils.Claims{
		UserID:    user.ID,
		GroupType: groupType,
	}
	token, err := utils.GenerateToken(claims)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "utils.GenerateToken",
			"err":       err,
		}).Error()
		appG.Response(http.StatusInternalServerError, e.ErrAuthTokenGen, nil)
		return
	}

	w.Services.Auth.CreateLoginLog(user, token)

	appG.Response(http.StatusOK, e.Success, map[string]string{
		"token": token,
	})
}
