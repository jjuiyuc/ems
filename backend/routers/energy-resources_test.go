package routers

import (
	"encoding/json"
	"fmt"
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
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	err = testutils.SeedUtCustomerAndGateway(db)
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

func (s *EnergyResourcesSuite) TearDownSuite() {
	models.Close()
}

func (s *EnergyResourcesSuite) Test_GetBatteryEnergyInfo() {
	prefixURL := fmt.Sprintf("/api/%s/devices/battery/energy-info", fixtures.UtGateway.UUID)
	seedUtURL := fmt.Sprintf("%s?startTime=%s", prefixURL, UtStartTime)
	seedUtInvalidParamsURL := fmt.Sprintf("%s?startTime=%s", prefixURL, "xxx")
	testResponseData := services.BatteryEnergyInfoResponse{
		BatteryLifetimeOperationCyclesDiff:  8,
		BatteryLifetimeOperationCycles:      16,
		BatterySoC:                          160,
		BatteryProducedLifetimeEnergyACDiff: 15,
		BatteryProducedLifetimeEnergyAC:     20,
		BatteryConsumedLifetimeEnergyACDiff: 0,
		BatteryConsumedLifetimeEnergyAC:     0,
		Model:                               "L051100-A UZ-Energy Battery",
		Capcity:                             30,
		PowerSources:                        "Solar + Grid",
		BatteryPower:                        24,
		Voltage:                             153.6,
	}

	tests := []testutils.TestInfo{
		{
			Name:       "batteryEnergyInfo",
			Token:      s.token,
			URL:        seedUtURL,
			WantStatus: http.StatusOK,
			WantRv: app.Response{
				Code: e.Success,
				Msg:  "ok",
				Data: testResponseData,
			},
		},
		{
			Name:       "batteryEnergyInfoInvalidParams",
			Token:      s.token,
			URL:        seedUtInvalidParamsURL,
			WantStatus: http.StatusBadRequest,
			WantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.Name)
		rvData := testutils.AssertRequest(tt, s.Require(), s.router, "GET", nil)
		if tt.Name == "batteryEnergyInfo" {
			dataMap := rvData.(map[string]interface{})
			dataJSON, err := json.Marshal(dataMap)
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			var data services.BatteryEnergyInfoResponse
			err = json.Unmarshal(dataJSON, &data)
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			s.Equalf(tt.WantRv.Data, data, e.ErrNewMessageNotEqual.Error())
		}
	}
}

func (s *EnergyResourcesSuite) Test_GetBatteryPowerState() {
	prefixURL := fmt.Sprintf("/api/%s/devices/battery/power-state", fixtures.UtGateway.UUID)
	seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, UtStartTime, UtEndTime)
	seedUtInvalidResolutionParamURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "xxx", UtStartTime, UtEndTime)
	seedUtInvalidStartTimeParamURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, "xxx", UtEndTime)
	seedUtInvalidEndTimeParamURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, UtStartTime, UtStartTime)
	seedUtInvalidPeriodEndTimeURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, UtStartTime, "2022-08-03T15:15:00.000Z")
	seedUtNoResolutionParamURL := fmt.Sprintf("%s?startTime=%s&endTime=%s", prefixURL, UtStartTime, UtEndTime)

	testTimestamps := []int{1659543000, 1659549599, 1659553199, 1659556799, 1659557100}
	testBatteryAveragePowerACs := []float32{-3.5, 0, 0, 0, -7}
	testResponseData := services.BatteryPowerStateResponse{
		Timestamps:             testTimestamps,
		BatteryAveragePowerACs: testBatteryAveragePowerACs,
		OnPeakTime:             testOnPeakTime,
	}

	tests := []testutils.TestInfo{
		{
			Name:       "batteryPowerState",
			Token:      s.token,
			URL:        seedUtURL,
			WantStatus: http.StatusOK,
			WantRv: app.Response{
				Code: e.Success,
				Msg:  "ok",
				Data: testResponseData,
			},
		},
		{
			Name:       "batteryPowerStateInvalidResolutionParam",
			Token:      s.token,
			URL:        seedUtInvalidResolutionParamURL,
			WantStatus: http.StatusBadRequest,
			WantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			Name:       "batteryPowerStateInvalidStartTimeParam",
			Token:      s.token,
			URL:        seedUtInvalidStartTimeParamURL,
			WantStatus: http.StatusBadRequest,
			WantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			Name:       "batteryPowerStateInvalidEndTimeParam",
			Token:      s.token,
			URL:        seedUtInvalidEndTimeParamURL,
			WantStatus: http.StatusBadRequest,
			WantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			Name:       "batteryPowerStateInvalidPeriodEndTime",
			Token:      s.token,
			URL:        seedUtInvalidPeriodEndTimeURL,
			WantStatus: http.StatusBadRequest,
			WantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
		{
			Name:       "batteryPowerStateNoResolutionParam",
			Token:      s.token,
			URL:        seedUtNoResolutionParamURL,
			WantStatus: http.StatusBadRequest,
			WantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.Name)
		rvData := testutils.AssertRequest(tt, s.Require(), s.router, "GET", nil)
		if tt.Name == "batteryPowerState" {
			dataMap := rvData.(map[string]interface{})
			dataJSON, err := json.Marshal(dataMap)
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			var data services.BatteryPowerStateResponse
			err = json.Unmarshal(dataJSON, &data)
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			s.Equalf(tt.WantRv.Data, data, e.ErrNewMessageNotEqual.Error())
		}
	}
}

func (s *EnergyResourcesSuite) Test_GetBatteryChargeVoltageState() {
	prefixURL := fmt.Sprintf("/api/%s/devices/battery/charge-voltage-state", fixtures.UtGateway.UUID)
	seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, UtResolution, UtStartTime, UtEndTime)
	seedUtInvalidParamsURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "xxx", UtStartTime, UtEndTime)

	testTimestamps := []int{1659543000, 1659549599, 1659553199, 1659556799, 1659557100}
	testBatterySoCs := []float32{80, 0, 0, 0, 160}
	testBatteryVoltages := []float32{28, 0, 0, 0, 56}
	testResponseData := services.BatteryChargeVoltageStateResponse{
		Timestamps:      testTimestamps,
		BatterySoCs:     testBatterySoCs,
		BatteryVoltages: testBatteryVoltages,
		OnPeakTime:      testOnPeakTime,
	}

	tests := []testutils.TestInfo{
		{
			Name:       "batteryChargeVoltageState",
			Token:      s.token,
			URL:        seedUtURL,
			WantStatus: http.StatusOK,
			WantRv: app.Response{
				Code: e.Success,
				Msg:  "ok",
				Data: testResponseData,
			},
		},
		{
			Name:       "batteryChargeVoltageStateInvalidParams",
			Token:      s.token,
			URL:        seedUtInvalidParamsURL,
			WantStatus: http.StatusBadRequest,
			WantRv: app.Response{
				Code: e.InvalidParams,
				Msg:  "invalid parameters",
			},
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.Name)
		rvData := testutils.AssertRequest(tt, s.Require(), s.router, "GET", nil)
		if tt.Name == "batteryChargeVoltageState" {
			dataMap := rvData.(map[string]interface{})
			dataJSON, err := json.Marshal(dataMap)
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			var data services.BatteryChargeVoltageStateResponse
			err = json.Unmarshal(dataJSON, &data)
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			s.Equalf(tt.WantRv.Data, data, e.ErrNewMessageNotEqual.Error())
		}
	}
}
