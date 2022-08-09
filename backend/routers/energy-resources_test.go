package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/services"
	"der-ems/testutils"
	"der-ems/testutils/fixtures"
)

const (
	UtResolution = "hour"
	UtStartTime  = "2022-08-03T16:00:00.000Z"
	UtEndTime    = "2022-08-03T20:15:00.000Z"
)

var testOnPeakTime = map[string]string{
	"timezone": "+0800",
	"start":    "07:30:00",
	"end":      "22:30:00",
}

type EnergyResourcesSuite struct {
	suite.Suite
	router *gin.Engine
	repo   *repository.Repository
	token  string
}

func Test_EnergyResources(t *testing.T) {
	suite.Run(t, &EnergyResourcesSuite{})
}

func (s *EnergyResourcesSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()

	repo := repository.NewRepository(db)
	w := &APIWorker{
		Services: services.NewServices(cfg, repo),
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

func (s *EnergyResourcesSuite) TearDownSuite() {
	models.Close()
}

func (s *EnergyResourcesSuite) Test_GetBatteryEnergyInfo() {
	type response struct {
		Code int                                 `json:"code"`
		Msg  string                              `json:"msg"`
		Data *services.BatteryEnergyInfoResponse `json:"data"`
	}

	prefixURL := "/api/" + fixtures.UtGateway.UUID + "/devices/battery/energy-info"
	seedUtURL := prefixURL + "?startTime=2022-08-03T16:00:00.000Z"
	seedUtInvalidParamsURL := prefixURL + "?startTime=xxx"
	testResponseData := &services.BatteryEnergyInfoResponse{
		BatteryOperationCycles:          8,
		BatteryLifetimeOperationCycles:  16,
		BatterySoC:                      80,
		BatteryProducedEnergyAC:         250,
		BatteryProducedLifetimeEnergyAC: 500,
		BatteryConsumedEnergyAC:         250,
		BatteryConsumedLifetimeEnergyAC: 500,
		Model:                           "L051100-A UZ-Energy Battery",
		Capcity:                         30,
		PowerSources:                    "Solar + Grid",
		BatteryPower:                    24,
		Voltage:                         153.6,
	}

	tests := []struct {
		name       string
		token      string
		url        string
		wantStatus int
		wantRv     response
	}{
		{
			name:       "batteryEnergyInfo",
			token:      s.token,
			url:        seedUtURL,
			wantStatus: http.StatusOK,
			wantRv: response{
				Code: e.Success,
				Msg:  "ok",
				Data: testResponseData,
			},
		},
		{
			name:       "batteryEnergyInfoInvalidParams",
			token:      s.token,
			url:        seedUtInvalidParamsURL,
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.name)
		req, err := http.NewRequest("GET", fmt.Sprintf(tt.url), nil)
		s.Require().NoError(err)
		req.Header.Set("Authorization", testutils.GetAuthorization(tt.token))
		rv := httptest.NewRecorder()
		s.router.ServeHTTP(rv, req)
		s.Equal(tt.wantStatus, rv.Code)

		var res response
		err = json.Unmarshal([]byte(rv.Body.String()), &res)
		s.Require().NoError(err)
		s.Equal(tt.wantRv.Code, res.Code)
		s.Equal(tt.wantRv.Msg, res.Msg)
		s.Equal(tt.wantRv.Data, res.Data)
	}
}

func (s *EnergyResourcesSuite) Test_GetBatteryPowerState() {
	type response struct {
		Code int                                 `json:"code"`
		Msg  string                              `json:"msg"`
		Data *services.BatteryPowerStateResponse `json:"data"`
	}

	prefixURL := fmt.Sprintf("/api/%s/devices/battery/power-state", fixtures.UtGateway.UUID)
	seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, UtStartTime, UtEndTime)
	seedUtInvalidResolutionParamURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "xxx", UtStartTime, UtEndTime)
	seedUtInvalidStartTimeParamURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, "xxx", UtEndTime)
	seedUtInvalidEndTimeParamURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, UtStartTime, UtStartTime)
	seedUtInvalidPeriodEndTimeURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, UtStartTime, "2022-08-03T16:15:00.000Z")
	seedUtNoResolutionParamURL := fmt.Sprintf("%s?startTime=%s&endTime=%s", prefixURL, UtStartTime, UtEndTime)

	testTimestamps := []int{1659542400, 1659546000, 1659549600, 1659553200, 1659556800}
	testBatteryAveragePowerACs := []float32{-3.5, 0, 0, 0, 0}
	testResponseData := &services.BatteryPowerStateResponse{
		Timestamps:             testTimestamps,
		BatteryAveragePowerACs: testBatteryAveragePowerACs,
		OnPeakTime:             testOnPeakTime,
	}

	tests := []struct {
		name       string
		token      string
		url        string
		wantStatus int
		wantRv     response
	}{
		{
			name:       "batteryPowerState",
			token:      s.token,
			url:        seedUtURL,
			wantStatus: http.StatusOK,
			wantRv: response{
				Code: e.Success,
				Msg:  "ok",
				Data: testResponseData,
			},
		},
		{
			name:       "batteryPowerStateInvalidResolutionParam",
			token:      s.token,
			url:        seedUtInvalidResolutionParamURL,
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name:       "batteryPowerStateInvalidStartTimeParam",
			token:      s.token,
			url:        seedUtInvalidStartTimeParamURL,
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name:       "batteryPowerStateInvalidEndTimeParam",
			token:      s.token,
			url:        seedUtInvalidEndTimeParamURL,
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name:       "batteryPowerStateInvalidPeriodEndTime",
			token:      s.token,
			url:        seedUtInvalidPeriodEndTimeURL,
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			name:       "batteryPowerStateNoResolutionParam",
			token:      s.token,
			url:        seedUtNoResolutionParamURL,
			wantStatus: http.StatusBadRequest,
			wantRv: response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.name)
		req, err := http.NewRequest("GET", fmt.Sprintf(tt.url), nil)
		s.Require().NoError(err)
		req.Header.Set("Authorization", testutils.GetAuthorization(tt.token))
		rv := httptest.NewRecorder()
		s.router.ServeHTTP(rv, req)
		s.Equal(tt.wantStatus, rv.Code)

		var res response
		err = json.Unmarshal([]byte(rv.Body.String()), &res)
		s.Require().NoError(err)
		s.Equal(tt.wantRv.Code, res.Code)
		s.Equal(tt.wantRv.Msg, res.Msg)
		s.Equal(tt.wantRv.Data, res.Data)
	}
}
