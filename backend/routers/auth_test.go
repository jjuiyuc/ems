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
	models.GetDB().Exec("truncate table login_log")
	testutils.SeedUtUser()

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
		Code int                    `json:"code"`
		Msg  string                 `json:"msg"`
		Data map[string]interface{} `json:"data"`
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
		payloadBuf, _ := json.Marshal(tt.args)
		req, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(payloadBuf))
		rv := httptest.NewRecorder()
		s.router.ServeHTTP(rv, req)
		s.Equal(tt.wantStatus, rv.Code)

		var res response
		json.Unmarshal([]byte(rv.Body.String()), &res)
		s.Equal(tt.wantRv.Code, res.Code)
		s.Equal(tt.wantRv.Msg, res.Msg)

		if tt.name == "login" {
			s.NotEmpty(res.Data["token"])
			count, _ := deremsmodels.LoginLogs().Count(models.GetDB())
			s.Equal(1, int(count))
		}
	}
}
