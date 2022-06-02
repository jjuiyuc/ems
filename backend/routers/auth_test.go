package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/testutils"
	"der-ems/testutils/fixtures"
)

type AuthorizationSuite struct {
	suite.Suite
	router *gin.Engine
}

func Test_Authorization(t *testing.T) {
	suite.Run(t, &AuthorizationSuite{})
}

func (s *AuthorizationSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	models.Init()

	// Truncate & seed data
	_, err := models.GetDB().Exec("truncate table login_log")
	s.Require().NoError(err)
	err = testutils.SeedUtUser()
	s.Require().NoError(err)

	s.router = InitRouter(config.GetConfig().GetBool("server.cors"), config.GetConfig().GetString("server.ginMode"))
}

func (s *AuthorizationSuite) TearDownSuite() {
	models.Close()
}

func (s *AuthorizationSuite) Test_GetAuth() {
	type args struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type response struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantRv     response
	}{
		{
			name: "login",
			args: args{
				Username: fixtures.UtUser.Username,
				Password: fixtures.UtUser.Password,
			},
			wantStatus: http.StatusOK,
			wantRv: response{
				Code: e.Success,
				Msg:  "ok",
			},
		},
		{
			name: "login_invalidParams",
			args: args{
				Username: fixtures.UtUser.Username,
			},
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name: "login_notExistUser",
			args: args{
				Username: "xxx",
				Password: fixtures.UtUser.Password,
			},
			wantStatus: http.StatusUnauthorized,
			wantRv: response{
				Code: e.ErrAuthLogin,
				Msg:  "fail",
			},
		},
	}

	for _, tt := range tests {
		payloadBuf, err := json.Marshal(tt.args)
		s.Require().NoError(err)
		req, err := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(payloadBuf))
		s.Require().NoError(err)
		rv := httptest.NewRecorder()
		s.router.ServeHTTP(rv, req)
		s.Equal(tt.wantStatus, rv.Code)

		var res response
		err = json.Unmarshal([]byte(rv.Body.String()), &res)
		s.Require().NoError(err)
		s.Equal(tt.wantRv.Code, res.Code)
		s.Equal(tt.wantRv.Msg, res.Msg)

		if tt.name == "login" {
			dataMap := res.Data.(map[string]interface{})
			s.NotEmpty(dataMap["token"])
			count, err := deremsmodels.LoginLogs().Count(models.GetDB())
			s.Require().NoError(err)
			s.Equal(1, int(count))
		}
	}
}
