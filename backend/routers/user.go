package routers

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// PasswordLost godoc
// @Summary     Send an email for reset the password
// @Description get email by username
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       username  body      string true "Username"
// @Success     200       {object}  app.Response
// @Failure     400       {object}  app.Response
// @Failure     401       {object}  app.Response
// @Router      /user/passwordLost [put]
func (w *APIWorker) PasswordLost(c *gin.Context) {
	appG := app.Gin{c}
	valid := validation.Validation{}

	var a struct {
		Username string `valid:"Required; MaxSize(50)"`
	}
	c.BindJSON(&a)
	if ok, err := valid.Valid(&a); !ok {
		log.WithFields(log.Fields{
			"caused-by": "valid.Valid",
			"err":       err,
		}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	name, token, err := w.Services.User.CreatePasswordToken(a.Username)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ErrPasswordLost, nil)
		return
	}
	err = w.Services.Email.SendResetEmail(c, name, a.Username, token)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ErrPasswordLost, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, map[string]string{
		"username": a.Username,
	})
}

// PasswordResetByToken godoc
// @Summary Reset the password
// @Description set a new password by having the token from email
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       token     body      string true "Token"
// @Param       password  body      string true "New password"
// @Success     200       {object}  app.Response
// @Failure     400       {object}  app.Response
// @Failure     401       {object}  app.Response
// @Router /user/PasswordResetByToken [put]
func (w *APIWorker) PasswordResetByToken(c *gin.Context) {
	appG := app.Gin{c}
	valid := validation.Validation{}

	var a struct {
		Token    string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(50)"`
	}
	c.BindJSON(&a)
	if ok, err := valid.Valid(&a); !ok {
		log.WithFields(log.Fields{
			"caused-by": "valid.Valid",
			"err":       err,
		}).Error()
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	if err := w.Services.User.PasswordResetByPasswordToken(a.Token, a.Password); err != nil {
		appG.Response(http.StatusUnauthorized, e.ErrPasswordToken, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
}

// GetProfile godoc
// @Summary Show the detailed information about an individual user
// @Description get user by token
// @Tags        user
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Produce     json
// @Success     200            {object}  app.Response{data=services.ProfileResponse}
// @Failure     401            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /user/profile [get]
func (w *APIWorker) GetProfile(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if userID == nil {
		log.WithFields(log.Fields{"caused-by": "error token"}).Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	profile, err := w.Services.User.GetProfile(userID.(int))
	if err != nil {
		log.WithFields(log.Fields{"caused-by": "get profile"}).Error()
		appG.Response(http.StatusInternalServerError, e.ErrUserProfileGen, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.Success, profile)
}
