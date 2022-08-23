package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/services"
	"der-ems/testutils"
	"der-ems/testutils/fixtures"
)

type AuthorizationSuite struct {
	suite.Suite
	router *gin.Engine
	repo   *repository.Repository
}

func Test_Authorization(t *testing.T) {
	suite.Run(t, &AuthorizationSuite{})
}

func (s *AuthorizationSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()

	repo := repository.NewRepository(db)
	services := services.NewServices(cfg, repo)
	worker := &APIWorker{
		Services: services,
	}

	// Truncate & seed data
	_, err := db.Exec("truncate table login_log")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	err = testutils.SeedUtUser(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())

	s.repo = repo
	s.router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), worker)
}

func (s *AuthorizationSuite) TearDownSuite() {
	models.Close()
}

func (s *AuthorizationSuite) Test_GetAuth() {
	type args struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type authTestInfo struct {
		testutils.TestInfo
		args
	}

	seedUtURL := "/api/auth"
	tests := []authTestInfo{
		{
			testutils.TestInfo{
				Name:       "login",
				URL:        seedUtURL,
				WantStatus: http.StatusOK,
				WantRv: app.Response{
					Code: e.Success,
					Msg:  "ok",
				},
			},
			args{
				Username: fixtures.UtUser.Username,
				Password: fixtures.UtUser.Password,
			},
		},
		{
			testutils.TestInfo{
				Name:       "loginInvalidParams",
				URL:        seedUtURL,
				WantStatus: http.StatusBadRequest,
				WantRv: app.Response{
					Code: e.InvalidParams,
					Msg:  "invalid parameters",
				},
			},
			args{
				Username: fixtures.UtUser.Username,
			},
		},
		{
			testutils.TestInfo{
				Name:       "loginUserNotExist",
				URL:        seedUtURL,
				WantStatus: http.StatusUnauthorized,
				WantRv: app.Response{
					Code: e.ErrAuthUserNotExist,
					Msg:  "fail",
				},
			},
			args{
				Username: "xxx",
				Password: fixtures.UtUser.Password,
			},
		},
		{
			testutils.TestInfo{
				Name:       "loginPasswordNotMatch",
				URL:        seedUtURL,
				WantStatus: http.StatusUnauthorized,
				WantRv: app.Response{
					Code: e.ErrAuthPasswordNotMatch,
					Msg:  "fail",
				},
			},
			args{
				Username: fixtures.UtUser.Username,
				Password: "xxx",
			},
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.Name)
		payloadBuf, err := json.Marshal(tt.args)
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		rvData := testutils.AssertRequest(tt.TestInfo, s.Require(), s.router, "POST", bytes.NewBuffer(payloadBuf))
		if tt.Name == "login" {
			dataMap := rvData.(map[string]interface{})
			s.NotEmpty(dataMap["token"])
			count, err := s.repo.User.GetLoginLogCount()
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			s.Equalf(1, int(count), e.ErrNewMessageNotEqual.Error())
		}
	}
}
