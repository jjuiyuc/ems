package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"

	"der-ems/config"
	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
	"der-ems/services"
	"der-ems/testutils"
	"der-ems/testutils/testdata"
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
	err := testutils.SeedUtWebpageAndRight(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	err = testutils.SeedUtGroupAndUser(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	err = testutils.SeedUtLocationAndGateway(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	token, err := utils.GenerateToken(testutils.SeedUtClaims())
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.token = token
	// Mock group_gateway_right table
	_, err = db.Exec("TRUNCATE TABLE group_gateway_right")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	_, err = db.Exec(`
		INSERT INTO group_gateway_right (id,group_id,gw_id,location_id,enabled_at) VALUES
		(1,1,1,1,'2022-07-01 00:00:00');
	`)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())

	s.repo = repo
	s.router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), initPolicy(testutils.GetConfigDir()), w)
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
				Username: testdata.UtUser.Username,
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
				Password: testdata.UtUser.Password,
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
				Password: testdata.UtUser.Password,
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
				Password: testdata.UtUser.Password,
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
			s.Equalf(testdata.UtUser.Username, dataMap["username"], e.ErrNewMessageNotEqual.Error())
		}
	}

	user, err := s.repo.User.GetUserByUsername(testdata.UtUser.Username)
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
	expectedWebpagePermissionsInfoType1 := services.WebpagePermissionsInfo{Create: false, Read: true, Update: false, Delete: false}
	expectedWebpagePermissionsInfoType2 := services.WebpagePermissionsInfo{Create: true, Read: true, Update: true, Delete: false}
	expectedWebpagePermissionsInfoType3 := services.WebpagePermissionsInfo{Create: true, Read: true, Update: true, Delete: true}
	expectedResponseData := services.ProfileResponse{
		User: &deremsmodels.User{
			ID:                 1,
			Username:           "ut-user@gmail.com",
			GroupID:            1,
			PasswordRetryCount: null.IntFrom(0),
			ExpirationDate:     null.TimeFrom(time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)),
			CreatedAt:          time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:          time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		Group: services.GroupInfo{
			Group: &deremsmodels.Group{
				ID:        1,
				Name:      "Admin",
				TypeID:    1,
				CreatedAt: time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
			},
			Gateways: []services.GatewayInfo{
				{
					GatewayID: "0E0BA27A8175AF978C49396BDE9D7A1E",
					Permissions: []services.GatewayPermissionInfo{
						{
							EnabledAt: time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
							Location: services.LocationInfo{
								Name:    "Field A",
								Address: null.StringFrom("宜蘭縣五結鄉大吉五路157巷68號"),
							},
						},
					},
				},
			},
			Webpages: []services.WebpageInfo{
				{ID: 1, Name: "dashboard", Permissions: expectedWebpagePermissionsInfoType1},
				{ID: 2, Name: "analysis", Permissions: expectedWebpagePermissionsInfoType1},
				{ID: 3, Name: "timeOfUseEnergy", Permissions: expectedWebpagePermissionsInfoType1},
				{ID: 4, Name: "economics", Permissions: expectedWebpagePermissionsInfoType1},
				{ID: 5, Name: "demandCharge", Permissions: expectedWebpagePermissionsInfoType1},
				{ID: 6, Name: "energyResources", Permissions: expectedWebpagePermissionsInfoType1},
				{ID: 7, Name: "fieldManagement", Permissions: expectedWebpagePermissionsInfoType2},
				{ID: 8, Name: "accountManagementGroup", Permissions: expectedWebpagePermissionsInfoType3},
				{ID: 9, Name: "accountManagementUser", Permissions: expectedWebpagePermissionsInfoType3},
				{ID: 10, Name: "settings", Permissions: expectedWebpagePermissionsInfoType3},
				{ID: 11, Name: "advancedSettings", Permissions: expectedWebpagePermissionsInfoType2},
			},
		},
	}
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
	s.Equalf(expectedResponseData, data, e.ErrNewMessageNotEqual.Error())
}
