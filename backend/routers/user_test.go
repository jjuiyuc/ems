package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	s.Require().NoError(err)
	err = testutils.SeedUtCustomerAndGateway(db)
	s.Require().NoError(err)
	token, err := utils.GenerateToken(fixtures.UtUser.ID)
	s.Require().NoError(err)
	s.token = token
	// Mock user_gateway_right table
	_, err = db.Exec("TRUNCATE TABLE user_gateway_right")
	s.Require().NoError(err)
	_, err = db.Exec(`
		INSERT INTO user_gateway_right (id,user_id,gw_id) VALUES
		(1,1,1);
	`)
	s.Require().NoError(err)

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

	passwordLostTest := []struct {
		name       string
		args       passwordLostArgs
		wantStatus int
		wantRv     app.Response
	}{
		{
			name: "passwordLost",
			args: passwordLostArgs{
				Username: fixtures.UtUser.Username,
			},
			wantStatus: http.StatusOK,
			wantRv: app.Response{
				Code: e.Success,
				Msg:  "ok",
			},
		},
		{
			name:       "passwordLostInvalidParams",
			wantStatus: http.StatusBadRequest,
			wantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name: "passwordLostError",
			args: passwordLostArgs{
				Username: "xxx",
			},
			wantStatus: http.StatusUnauthorized,
			wantRv: app.Response{
				Code: e.ErrPasswordLost,
				Msg:  "fail",
			},
		},
	}

	passwordResetTest := []struct {
		name       string
		args       passwordResetArgs
		wantStatus int
		wantRv     app.Response
	}{
		{
			name: "passwordResetByToken",
			args: passwordResetArgs{
				Password: fixtures.UtUser.Password,
			},
			wantStatus: http.StatusOK,
			wantRv: app.Response{
				Code: e.Success,
				Msg:  "ok",
			},
		},
		{
			name: "passwordResetByTokenInvalidParams",
			args: passwordResetArgs{
				Password: fixtures.UtUser.Password,
			},
			wantStatus: http.StatusBadRequest,
			wantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name: "passwordResetByTokenError",
			args: passwordResetArgs{
				Token:    "xxx",
				Password: fixtures.UtUser.Password,
			},
			wantStatus: http.StatusUnauthorized,
			wantRv: app.Response{
				Code: e.ErrPasswordToken,
				Msg:  "fail",
			},
		},
	}

	for _, tt := range passwordLostTest {
		log.Info("test name: ", tt.name)
		payloadBuf, err := json.Marshal(tt.args)
		s.Require().NoError(err)
		req, err := http.NewRequest("PUT", "/api/users/password/lost", bytes.NewBuffer(payloadBuf))
		s.Require().NoError(err)
		rv := httptest.NewRecorder()
		s.router.ServeHTTP(rv, req)
		s.Equal(tt.wantStatus, rv.Code)

		var res app.Response
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

	for _, tt := range passwordResetTest {
		log.Info("test name: ", tt.name)
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

		var res app.Response
		err = json.Unmarshal([]byte(rv.Body.String()), &res)
		s.Require().NoError(err)
		s.Equal(tt.wantRv.Code, res.Code)
		s.Equal(tt.wantRv.Msg, res.Msg)
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
		testutils.ValidateGetRequestStatusAndCode(tt, s.Require(), s.router)
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
	rvData := testutils.ValidateGetRequestStatusAndCode(tt, s.Require(), s.router)
	dataMap := rvData.(map[string]interface{})
	dataJSON, err := json.Marshal(dataMap)
	s.Require().NoError(err)
	var data services.ProfileResponse
	err = json.Unmarshal(dataJSON, &data)
	s.Require().NoError(err)
	s.Equal(fixtures.UtUser.ID, data.ID)
	s.Equal(fixtures.UtUser.Username, data.Username)
	s.Equal(fixtures.UtUser.ExpirationDate, data.ExpirationDate)
	s.Equal(fixtures.UtGateway.UUID, data.Gateways[0].GatewayID)
	s.Equal(fixtures.UtCustomer.Address.String, data.Gateways[0].Address)
}
