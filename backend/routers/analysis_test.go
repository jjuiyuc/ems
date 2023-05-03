package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"der-ems/config"
	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/services"
	"der-ems/testutils"
	"der-ems/testutils/testdata"
)

var _ = Describe("Analysis", func() {
	const (
		UtStartTime        = "2022-08-03T16:00:00.000Z"
		UtEndTime          = "2022-08-03T20:15:00.000Z"
		UtStartTimeForWeek = "2022-07-30T16:00:00.000Z"
		UtEndTimeForWeek   = "2022-08-02T16:00:00.000Z"
	)

	var (
		router *gin.Engine
		token  string
		err    error
	)

	BeforeEach(func() {
		config.Init(testutils.GetConfigDir(), "ut.yaml")
		cfg := config.GetConfig()
		models.Init(cfg)
		db := models.GetDB()

		repo := repository.NewRepository(db)
		w := &APIWorker{
			Services: services.NewServices(cfg, repo),
		}

		// Truncate & seed data
		err = testutils.SeedUtUser(db)
		Expect(err).Should(BeNil())
		err = testutils.SeedUtLocationAndGateway(db)
		Expect(err).Should(BeNil())
		token, err = utils.GenerateToken(testutils.SeedUtClaims())
		Expect(err).Should(BeNil())
		// Mock group_gateway_right table
		_, err = db.Exec("TRUNCATE TABLE group_gateway_right")
		Expect(err).Should(BeNil())
		_, err = db.Exec(`
			INSERT INTO group_gateway_right (id,group_id,gw_id,location_id,enabled_at) VALUES
			(1,1,1,1,'2022-07-01 00:00:00');
		`)
		Expect(err).Should(BeNil())

		router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), initPolicy(testutils.GetConfigDir()), w)
	})

	AfterEach(func() {
		models.Close()
	})

	Describe("GetEnergyDistributionInfo", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/energy-distribution-info", testdata.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?startTime=%s&endTime=%s", prefixURL, UtStartTime, UtEndTime)
				expectedResponseData := services.EnergyDistributionInfoResponse{
					AllProducedLifetimeEnergyACDiff:     50,
					PvProducedEnergyPercentAC:           20,
					GridProducedEnergyPercentAC:         50,
					BatteryProducedEnergyPercentAC:      30,
					PvProducedLifetimeEnergyACDiff:      10,
					GridProducedLifetimeEnergyACDiff:    25,
					BatteryProducedLifetimeEnergyACDiff: 15,
					AllConsumedLifetimeEnergyACDiff:     50,
					LoadConsumedEnergyPercentAC:         70,
					GridConsumedEnergyPercentAC:         30,
					BatteryConsumedEnergyPercentAC:      0,
					LoadConsumedLifetimeEnergyACDiff:    35,
					GridConsumedLifetimeEnergyACDiff:    15,
					BatteryConsumedLifetimeEnergyACDiff: 0,
				}
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusOK,
					WantRv: app.Response{
						Code: e.Success,
						Msg:  "ok",
						Data: expectedResponseData,
					},
				}
				rvData := testutils.GinkgoAssertRequest(tt, router, "GET", nil)
				dataMap := rvData.(map[string]interface{})
				dataJSON, err := json.Marshal(dataMap)
				Expect(err).Should(BeNil())
				var data services.EnergyDistributionInfoResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/energy-distribution-info", testdata.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?startTime=%s&endTime=%s", prefixURL, UtStartTime, "xxx")
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtInvalidParamsURL,
					WantStatus: http.StatusBadRequest,
					WantRv: app.Response{
						Code: e.InvalidParams,
						Msg:  "invalid parameters",
					},
				}
				testutils.GinkgoAssertRequest(tt, router, "GET", nil)
			})
		})
	})

	Describe("GetPowerState", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/power-state", testdata.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", UtStartTime, UtEndTime)
				expectedTimestamps := []int{1659543000, 1659557100}
				expectedLoadAveragePowerACs := []float32{30, 30}
				expectedPvAveragePowerACs := []float32{40, 40}
				expectedBatteryAveragePowerACs := []float32{-3.5, -7}
				expectedGridAveragePowerACs := []float32{50, 50}
				expectedResponseData := services.PowerStateResponse{
					Timestamps:             expectedTimestamps,
					LoadAveragePowerACs:    expectedLoadAveragePowerACs,
					PvAveragePowerACs:      expectedPvAveragePowerACs,
					BatteryAveragePowerACs: expectedBatteryAveragePowerACs,
					GridAveragePowerACs:    expectedGridAveragePowerACs,
				}
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusOK,
					WantRv: app.Response{
						Code: e.Success,
						Msg:  "ok",
						Data: expectedResponseData,
					},
				}
				rvData := testutils.GinkgoAssertRequest(tt, router, "GET", nil)
				dataMap := rvData.(map[string]interface{})
				dataJSON, err := json.Marshal(dataMap)
				Expect(err).Should(BeNil())
				var data services.PowerStateResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/power-state", testdata.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "xxx", UtStartTime, UtEndTime)
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtInvalidParamsURL,
					WantStatus: http.StatusBadRequest,
					WantRv: app.Response{
						Code: e.InvalidParams,
						Msg:  "invalid parameters",
					},
				}
				testutils.GinkgoAssertRequest(tt, router, "GET", nil)
			})
		})
	})

	Describe("GetAccumulatedPowerState", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/accumulated-power-state", testdata.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "day", UtStartTimeForWeek, UtEndTimeForWeek)
				expectedTimestamps := []int{1659283140, 1659369555, 1659455970}
				expectedLoadConsumedLifetimeEnergyACDiffs := []float32{10, 20, 30}
				expectedPvProducedLifetimeEnergyACDiffs := []float32{5, 10, 15}
				expectedBatteryLifetimeEnergyACDiffs := []float32{15, 30, 45}
				expectedGridLifetimeEnergyACDiffs := []float32{5, 15, 25}
				expectedResponseData := services.AccumulatedPowerStateResponse{
					Timestamps:                        expectedTimestamps,
					LoadConsumedLifetimeEnergyACDiffs: expectedLoadConsumedLifetimeEnergyACDiffs,
					PvProducedLifetimeEnergyACDiffs:   expectedPvProducedLifetimeEnergyACDiffs,
					BatteryLifetimeEnergyACDiffs:      expectedBatteryLifetimeEnergyACDiffs,
					GridLifetimeEnergyACDiffs:         expectedGridLifetimeEnergyACDiffs,
				}
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusOK,
					WantRv: app.Response{
						Code: e.Success,
						Msg:  "ok",
						Data: expectedResponseData,
					},
				}
				rvData := testutils.GinkgoAssertRequest(tt, router, "GET", nil)
				dataMap := rvData.(map[string]interface{})
				dataJSON, err := json.Marshal(dataMap)
				Expect(err).Should(BeNil())
				var data services.AccumulatedPowerStateResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/accumulated-power-state", testdata.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "xxx", UtStartTimeForWeek, UtEndTimeForWeek)
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtInvalidParamsURL,
					WantStatus: http.StatusBadRequest,
					WantRv: app.Response{
						Code: e.InvalidParams,
						Msg:  "invalid parameters",
					},
				}
				testutils.GinkgoAssertRequest(tt, router, "GET", nil)
			})
		})
	})

	Describe("GetPowerSelfSupplyRate", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/power-self-supply-rate", testdata.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "day", UtStartTimeForWeek, UtEndTimeForWeek)
				expectedTimestamps := []int{1659283140, 1659369555, 1659455970}
				expectedLoadSelfConsumedEnergyPercentACs := []float32{10, 15, 20}
				expectedResponseData := services.PowerSelfSupplyRateResponse{
					Timestamps:                       expectedTimestamps,
					LoadSelfConsumedEnergyPercentACs: expectedLoadSelfConsumedEnergyPercentACs,
				}
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusOK,
					WantRv: app.Response{
						Code: e.Success,
						Msg:  "ok",
						Data: expectedResponseData,
					},
				}
				rvData := testutils.GinkgoAssertRequest(tt, router, "GET", nil)
				dataMap := rvData.(map[string]interface{})
				dataJSON, err := json.Marshal(dataMap)
				Expect(err).Should(BeNil())
				var data services.PowerSelfSupplyRateResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/power-self-supply-rate", testdata.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "xxx", UtStartTimeForWeek, UtEndTimeForWeek)
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtInvalidParamsURL,
					WantStatus: http.StatusBadRequest,
					WantRv: app.Response{
						Code: e.InvalidParams,
						Msg:  "invalid parameters",
					},
				}
				testutils.GinkgoAssertRequest(tt, router, "GET", nil)
			})
		})
	})
})
