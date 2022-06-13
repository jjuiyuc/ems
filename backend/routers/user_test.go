package routers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
	"der-ems/services"
	"der-ems/testutils"
	"der-ems/testutils/fixtures"
)

type UserSuite struct {
	suite.Suite
	router *gin.Engine
	repo   *repository.Repository
	token  string
}

func Test_User(t *testing.T) {
	suite.Run(t, &UserSuite{})
}

func (s *UserSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()

	repo := repository.NewRepository(db)
	services := &services.Services{
		Email: &testutils.MockEmailService{},
		User:  services.NewUserService(repo),
	}
	w := &APIWorker{
		Services: services,
	}

	// Truncate & seed data
	err := testutils.SeedUtUser(db)
	s.Require().NoError(err)
	token, err := utils.GenerateToken(fixtures.UtUser.ID)
	s.Require().NoError(err)
	s.token = token

	s.repo = repo
	s.router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), w)
}

func (s *UserSuite) TearDownSuite() {
	models.Close()
}

func (s *UserSuite) Test_PasswordLostAndResetByToken() {
	type args1 struct {
		Username string `json:"username"`
	}

	type args2 struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	type response struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	tests1 := []struct {
		name       string
		args       args1
		wantStatus int
		wantRv     response
	}{
		{
			name: "passwordLost",
			args: args1{
				Username: fixtures.UtUser.Username,
			},
			wantStatus: http.StatusOK,
			wantRv: response{
				Code: e.Success,
				Msg:  "ok",
			},
		},
		{
			name:       "passwordLostInvalidParams",
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name: "passwordLostError",
			args: args1{
				Username: "xxx",
			},
			wantStatus: http.StatusUnauthorized,
			wantRv: response{
				Code: e.ErrPasswordLost,
				Msg:  "fail",
			},
		},
	}

	tests2 := []struct {
		name       string
		args       args2
		wantStatus int
		wantRv     response
	}{
		{
			name: "passwordResetByToken",
			args: args2{
				Password: fixtures.UtUser.Password,
			},
			wantStatus: http.StatusOK,
			wantRv: response{
				Code: e.Success,
				Msg:  "ok",
			},
		},
		{
			name: "passwordResetByTokenInvalidParams",
			args: args2{
				Password: fixtures.UtUser.Password,
			},
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name: "passwordResetByTokenError",
			args: args2{
				Token:    "xxx",
				Password: fixtures.UtUser.Password,
			},
			wantStatus: http.StatusUnauthorized,
			wantRv: response{
				Code: e.ErrPasswordToken,
				Msg:  "fail",
			},
		},
	}

	for _, tt := range tests1 {
		payloadBuf, err := json.Marshal(tt.args)
		s.Require().NoError(err)
		req, err := http.NewRequest("PUT", "/api/users/password/lost", bytes.NewBuffer(payloadBuf))
		s.Require().NoError(err)
		rv := httptest.NewRecorder()
		s.router.ServeHTTP(rv, req)
		s.Equal(tt.wantStatus, rv.Code)

		var res response
		err = json.Unmarshal([]byte(rv.Body.String()), &res)
		s.Require().NoError(err)
		s.Equal(tt.wantRv.Code, res.Code)
		s.Equal(tt.wantRv.Msg, res.Msg)

		if tt.name == "passwordLost" {
			dataMap := res.Data.(map[string]interface{})
			s.Equal(fixtures.UtUser.Username, dataMap["username"])
		}
	}

	user, err := s.repo.User.GetUserByUsername(fixtures.UtUser.Username)
	s.Require().NoError(err)

	for _, tt := range tests2 {
		if tt.name == "passwordResetByToken" {
			tt.args.Token = user.ResetPWDToken.String
		}

		payloadBuf, err := json.Marshal(tt.args)
		s.Require().NoError(err)
		req, err := http.NewRequest("PUT", "/api/users/password/reset-by-token", bytes.NewBuffer(payloadBuf))
		s.Require().NoError(err)
		rv := httptest.NewRecorder()
		s.router.ServeHTTP(rv, req)
		s.Equal(tt.wantStatus, rv.Code)

		var res response
		err = json.Unmarshal([]byte(rv.Body.String()), &res)
		s.Require().NoError(err)
		s.Equal(tt.wantRv.Code, res.Code)
		s.Equal(tt.wantRv.Msg, res.Msg)

		if tt.name == "passwordResetByToken" {
			dataMap := res.Data.(map[string]interface{})
			s.Equal("success", dataMap["status"])
		}
	}
}

func (s *UserSuite) Test_Authorize() {
	type response struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	tests := []struct {
		name       string
		token      string
		wantStatus int
		wantRv     response
	}{
		{
			name:       "authorizeNoHeader",
			token:      s.token,
			wantStatus: http.StatusUnauthorized,
			wantRv: response{
				Code: e.ErrAuthNoHeader,
				Msg:  "fail",
			},
		},
		{
			name:       "authorizeInvalidHeader",
			token:      "xxx xxx",
			wantStatus: http.StatusUnauthorized,
			wantRv: response{
				Code: e.ErrAuthInvalidHeader,
				Msg:  "fail",
			},
		},
		{
			name:       "authorizeWrongToken",
			token:      "xxx",
			wantStatus: http.StatusUnauthorized,
			wantRv: response{
				Code: e.ErrAuthTokenParse,
				Msg:  "fail",
			},
		},
	}

	for _, tt := range tests {
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/profile"), nil)
		s.Require().NoError(err)
		if tt.name != "authorizeNoHeader" {
			req.Header.Set("Authorization", testutils.GetAuthorization(tt.token))
		}
		rv := httptest.NewRecorder()
		s.router.ServeHTTP(rv, req)
		s.Equal(tt.wantStatus, rv.Code)

		var res response
		err = json.Unmarshal([]byte(rv.Body.String()), &res)
		s.Require().NoError(err)
		s.Equal(tt.wantRv.Code, res.Code)
		s.Equal(tt.wantRv.Msg, res.Msg)
	}
}

func (s *UserSuite) Test_GetProfile() {
	type response struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	tt := struct {
		name       string
		wantStatus int
		wantRv     response
	}{
		name:       "profile",
		wantStatus: http.StatusOK,
		wantRv: response{
			Code: e.Success,
			Msg:  "ok",
		},
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/profile"), nil)
	s.Require().NoError(err)
	req.Header.Set("Authorization", testutils.GetAuthorization(s.token))
	rv := httptest.NewRecorder()
	s.router.ServeHTTP(rv, req)
	s.Equal(tt.wantStatus, rv.Code)

	var res response
	err = json.Unmarshal([]byte(rv.Body.String()), &res)
	s.Require().NoError(err)
	s.Equal(tt.wantRv.Code, res.Code)
	s.Equal(tt.wantRv.Msg, res.Msg)

	dataMap := res.Data.(map[string]interface{})
	dataJSON, err := json.Marshal(dataMap)
	s.Require().NoError(err)
	var data deremsmodels.User
	err = json.Unmarshal(dataJSON, &data)
	s.Require().NoError(err)
	s.Equal(fixtures.UtUser.ID, data.ID)
	s.Equal(fixtures.UtUser.Username, data.Username)
	s.Equal(fixtures.UtUser.ExpirationDate, data.ExpirationDate)
}
