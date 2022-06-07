package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/repository"
	"der-ems/services"
)

// PasswordLost sends an email for reset the password
// @Summary Send an email for reset the password
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param username path string true "Username"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Router /user/passwordLost [put]
func PasswordLost(c *gin.Context) {
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

	repo := repository.NewUserRepository()
	s := services.NewUserService(repo)
	err := s.PasswordLost(c, a.Username)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ErrPasswordLost, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, map[string]string{
		"username": a.Username,
	})
}

// GetProfile provides the detailed information about an individual user
// @Summary Provide detailed information about an individual user
// @Tags User
// @Security ApiKeyAuth
// @Param Authorization header string true "Input user's access token" default(Bearer <Add access token here>)
// @Produce application/json
// @Success 200 {object} app.Response
// @Failure 401 {object} app.Response
// @Router /user/profile [get]
func GetProfile(c *gin.Context) {
	appG := app.Gin{c}
	userID, _ := c.Get("userID")
	if userID == nil {
		log.WithFields(log.Fields{"caused-by": "error token"}).Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}

	repo := repository.NewUserRepository()
	s := services.NewUserService(repo)
	user, err := s.GetProfile(userID.(int))
	if err != nil {
		log.WithFields(log.Fields{"caused-by": "error token"}).Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, user)
}
