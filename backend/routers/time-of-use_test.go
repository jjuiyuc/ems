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
	"der-ems/testutils/fixtures"
)

var _ = Describe("TimeOfUse", func() {
	const (
		UtStartTime = "2022-08-03T16:00:00.000Z"
		UtEndTime   = "2022-08-03T20:15:00.000Z"
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
		token, err = utils.GenerateToken(fixtures.UtUser.ID)
		Expect(err).Should(BeNil())
		// Mock group_gateway_right table
		_, err = db.Exec("TRUNCATE TABLE group_gateway_right")
		Expect(err).Should(BeNil())
		_, err = db.Exec(`
			INSERT INTO group_gateway_right (id,group_id,gw_id,enabled_at) VALUES
			(1,2,1,'2022-07-01 00:00:00');
		`)
		Expect(err).Should(BeNil())

		router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), w)
	})

	AfterEach(func() {
		models.Close()
	})

	Describe("GetBatteryUsageInfo", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/usage-info", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?startTime=%s", prefixURL, UtStartTime)
				expectedResponseData := services.BatteryUsageInfoResponse{
					BatterySoC:                    160,
					BatteryProducedAveragePowerAC: 20,
					BatteryConsumedAveragePowerAC: 0,
					BatteryChargingFrom:           "Solar",
					BatteryDischargingTo:          "",
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
				var data services.BatteryUsageInfoResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/usage-info", fixtures.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?startTime=%s", prefixURL, "xxx")
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

	Describe("GetTimeOfUseInfo", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/tou/info", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?startTime=%s", prefixURL, UtStartTime)
				expectedEnergySources := map[string]interface{}{
					"offPeak": map[string]interface{}{
						"allProducedLifetimeEnergyACDiff":     50.0,
						"gridProducedEnergyPercentAC":         50.0,
						"gridProducedLifetimeEnergyACDiff":    25.0,
						"pvProducedEnergyPercentAC":           20.0,
						"pvProducedLifetimeEnergyACDiff":      10.0,
						"batteryProducedEnergyPercentAC":      30.0,
						"batteryProducedLifetimeEnergyACDiff": 15.0,
					},
				}
				expectedTimeOfUse := map[string]interface{}{
					"timezone":        "+0800",
					"currentPeakType": "onPeak",
					"offPeak": []interface{}{
						map[string]interface{}{
							"end":     "07:30:00",
							"start":   "00:00:00",
							"touRate": 1.46,
						},
						map[string]interface{}{
							"end":     "24:00:00",
							"start":   "22:30:00",
							"touRate": 1.46,
						},
					},
					"onPeak": []interface{}{
						map[string]interface{}{
							"end":     "22:30:00",
							"start":   "07:30:00",
							"touRate": 3.42,
						},
					},
					"midPeak": nil,
				}
				expectedResponseData := services.TimeOfUseInfoResponse{
					EnergySources: expectedEnergySources,
					TimeOfUse:     expectedTimeOfUse,
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
				var data services.TimeOfUseInfoResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/tou/info", fixtures.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?startTime=%s", prefixURL, "xxx")
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

	Describe("GetSolarEnergyUsage", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/solar/energy-usage", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", UtStartTime, UtEndTime)
				expectedTimestamps := []int{1659543000, 1659557100}
				expectedLoadPvConsumedEnergyPercentACs := []float32{50, 50}
				expectedResponseData := services.SolarEnergyUsageResponse{
					Timestamps:                     expectedTimestamps,
					LoadPvConsumedEnergyPercentACs: expectedLoadPvConsumedEnergyPercentACs,
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
				var data services.SolarEnergyUsageResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/solar/energy-usage", fixtures.UtGateway.UUID)
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
})
