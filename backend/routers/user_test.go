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
	"der-ems/internal/utils"
	"der-ems/models"
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
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	err = testutils.SeedUtLocationAndGateway(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	token, err := utils.GenerateToken(fixtures.UtUser.ID)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.token = token
	// Mock user_gateway_right table
	_, err = db.Exec("TRUNCATE TABLE user_gateway_right")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	_, err = db.Exec(`
		INSERT INTO user_gateway_right (id,user_id,gw_id) VALUES
		(1,1,1);
	`)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())

	s.repo = repo
	s.router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), w)
}

func (s *UserSuite) TearDownSuite() {
	models.Close()
}

func (s *UserSuite) Test_PasswordLostAndResetByToken() {
	type passwordLostArgs struct {
		Username string `json:"username"`
	}

	type passwordResetArgs struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	type passwordLostTestInfo struct {
		testutils.TestInfo
		args passwordLostArgs
	}

	type passwordResetTestInfo struct {
		testutils.TestInfo
		args passwordResetArgs
	}

	seedUtPasswordLostURL := "/api/users/password/lost"
	passwordLostTest := []passwordLostTestInfo{
		{
			testutils.TestInfo{
				Name:       "passwordLost",
				URL:        seedUtPasswordLostURL,
				WantStatus: http.StatusOK,
				WantRv: app.Response{
					Code: e.Success,
					Msg:  "ok",
				},
			},
			passwordLostArgs{
				Username: fixtures.UtUser.Username,
			},
		},
		{
			testutils.TestInfo{
				Name:       "passwordLostInvalidParams",
				URL:        seedUtPasswordLostURL,
				WantStatus: http.StatusBadRequest,
				WantRv: app.Response{
					Code: e.InvalidParams,
					Msg:  "invalid parameters",
				},
			},
			passwordLostArgs{},
		},
		{
			testutils.TestInfo{
				Name:       "passwordLostError",
				URL:        seedUtPasswordLostURL,
				WantStatus: http.StatusUnauthorized,
				WantRv: app.Response{
					Code: e.ErrPasswordLost,
					Msg:  "fail",
				},
			},
			passwordLostArgs{
				Username: "xxx",
			},
		},
	}

	seedUtPasswordResetURL := "/api/users/password/reset-by-token"
	passwordResetTest := []passwordResetTestInfo{
		{
			testutils.TestInfo{
				Name:       "passwordResetByToken",
				URL:        seedUtPasswordResetURL,
				WantStatus: http.StatusOK,
				WantRv: app.Response{
					Code: e.Success,
					Msg:  "ok",
				},
			},
			passwordResetArgs{
				Password: fixtures.UtUser.Password,
			},
		},
		{
			testutils.TestInfo{
				Name:       "passwordResetByTokenInvalidParams",
				URL:        seedUtPasswordResetURL,
				WantStatus: http.StatusBadRequest,
				WantRv: app.Response{
					Code: e.InvalidParams,
					Msg:  "invalid parameters",
				},
			},
			passwordResetArgs{
				Password: fixtures.UtUser.Password,
			},
		},
		{
			testutils.TestInfo{
				Name:       "passwordResetByTokenError",
				URL:        seedUtPasswordResetURL,
				WantStatus: http.StatusUnauthorized,
				WantRv: app.Response{
					Code: e.ErrPasswordToken,
					Msg:  "fail",
				},
			},
			passwordResetArgs{
				Token:    "xxx",
				Password: fixtures.UtUser.Password,
			},
		},
	}

	for _, tt := range passwordLostTest {
		log.Info("test name: ", tt.Name)
		payloadBuf, err := json.Marshal(tt.args)
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		rvData := testutils.AssertRequest(tt.TestInfo, s.Require(), s.router, "PUT", bytes.NewBuffer(payloadBuf))
		if tt.Name == "passwordLost" {
			dataMap := rvData.(map[string]interface{})
			s.Equalf(fixtures.UtUser.Username, dataMap["username"], e.ErrNewMessageNotEqual.Error())
		}
	}

	user, err := s.repo.User.GetUserByUsername(fixtures.UtUser.Username)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())

	for _, tt := range passwordResetTest {
		log.Info("test name: ", tt.Name)
		if tt.Name == "passwordResetByToken" {
			tt.args.Token = user.ResetPWDToken.String
		}
		payloadBuf, err := json.Marshal(tt.args)
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		testutils.AssertRequest(tt.TestInfo, s.Require(), s.router, "PUT", bytes.NewBuffer(payloadBuf))
	}
}

func (s *UserSuite) Test_Authorize() {
	seedUtURL := "/api/users/profile"
	tests := []testutils.TestInfo{
		{
			Name:       "authorizeNoHeader",
			URL:        seedUtURL,
			WantStatus: http.StatusUnauthorized,
			WantRv: app.Response{
				Code: e.ErrAuthNoHeader,
				Msg:  "fail",
			},
		},
		{
			Name:       "authorizeInvalidHeader",
			Token:      "xxx xxx",
			URL:        seedUtURL,
			WantStatus: http.StatusUnauthorized,
			WantRv: app.Response{
				Code: e.ErrAuthInvalidHeader,
				Msg:  "fail",
			},
		},
		{
			Name:       "authorizeWrongToken",
			Token:      "xxx",
			URL:        seedUtURL,
			WantStatus: http.StatusUnauthorized,
			WantRv: app.Response{
				Code: e.ErrAuthTokenParse,
				Msg:  "fail",
			},
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.Name)
		testutils.AssertRequest(tt, s.Require(), s.router, "GET", nil)
	}
}

func (s *UserSuite) Test_GetProfile() {
	seedUtURL := "/api/users/profile"
	tt := testutils.TestInfo{
		Name:       "profile",
		Token:      s.token,
		URL:        seedUtURL,
		WantStatus: http.StatusOK,
		WantRv: app.Response{
			Code: e.Success,
			Msg:  "ok",
		},
	}

	log.Info("test name: ", tt.Name)
	rvData := testutils.AssertRequest(tt, s.Require(), s.router, "GET", nil)
	dataMap := rvData.(map[string]interface{})
	dataJSON, err := json.Marshal(dataMap)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	var data services.ProfileResponse
	err = json.Unmarshal(dataJSON, &data)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.Equalf(fixtures.UtUser.ID, data.ID, e.ErrNewMessageNotEqual.Error())
	s.Equalf(fixtures.UtUser.Username, data.Username, e.ErrNewMessageNotEqual.Error())
	s.Equalf(fixtures.UtUser.ExpirationDate, data.ExpirationDate, e.ErrNewMessageNotEqual.Error())
	s.Equalf(fixtures.UtGateway.UUID, data.Gateways[0].GatewayID, e.ErrNewMessageNotEqual.Error())
	s.Equalf(fixtures.UtLocation.Address.String, data.Gateways[0].Address, e.ErrNewMessageNotEqual.Error())
}
