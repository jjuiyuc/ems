package routers

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// PersonalName godoc
type PersonalName struct {
	Name string `form:"name" binding:"required,max=20"`
}

// PersonalPassword godoc
type PersonalPassword struct {
	CurrentPassword string `form:"currentPassword" binding:"required,max=50"`
	NewPassword     string `form:"newPassword" binding:"required,max=50"`
}

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
		log.WithField("caused-by", "error token").Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	profile, err := w.Services.User.GetProfile(userID.(int64))
	if err != nil {
		log.WithField("caused-by", "get profile").Error()
		appG.Response(http.StatusInternalServerError, e.ErrUserProfileGen, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, profile)
}

// UpdateName godoc
// @Summary Update the display name about an individual user
// @Description update user's name by token
// @Tags        user
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Accept      json
// @Param       name           body      string true "Name"
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /users/name [put]
func (w *APIWorker) UpdateName(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	var json PersonalName
	if err := c.BindJSON(&json); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	if err := w.Services.User.UpdateName(userID.(int64), json.Name); err != nil {
		log.WithField("caused-by", "update name").Error()
		appG.Response(http.StatusInternalServerError, e.ErrNameUpdate, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}

// UpdatePassword godoc
// @Summary Update the password about an individual user
// @Description update user's password by token
// @Tags        user
// @Security    ApiKeyAuth
// @Param       Authorization  header    string true "Input user's access token" default(Bearer <Add access token here>)
// @Accept      json
// @Param       password       body      string true "Password"
// @Success     200            {object}  app.Response
// @Failure     400            {object}  app.Response
// @Failure     401            {object}  app.Response
// @Failure     500            {object}  app.Response
// @Router      /users/password [put]
func (w *APIWorker) UpdatePassword(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	var json PersonalPassword
	if err := c.BindJSON(&json); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	errCode, err := w.Services.User.UpdatePassword(userID.(int64), json.CurrentPassword, json.NewPassword)
	if err != nil {
		log.WithField("caused-by", "update password").Error()
		if errCode == e.ErrAuthPasswordNotMatch {
			appG.Response(http.StatusUnauthorized, e.ErrAuthPasswordNotMatch, nil)
		} else {
			appG.Response(http.StatusInternalServerError, e.ErrPasswordUpdate, nil)
		}
		return
	}
	appG.Response(http.StatusOK, e.Success, nil)
}
