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

var _ = Describe("EnergyResources", func() {
	const (
		UtStartTime = "2022-08-03T16:00:00.000Z"
		UtEndTime   = "2022-08-03T20:15:00.000Z"
	)

	var expectedOnPeakTime = map[string]string{
		"timezone": "+0800",
		"start":    "07:30:00",
		"end":      "22:30:00",
	}

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
		err = testutils.SeedUtCustomerAndGateway(db)
		Expect(err).Should(BeNil())
		token, err = utils.GenerateToken(fixtures.UtUser.ID)
		Expect(err).Should(BeNil())
		// Mock user_gateway_right table
		_, err = db.Exec("TRUNCATE TABLE user_gateway_right")
		Expect(err).Should(BeNil())
		_, err = db.Exec(`
			INSERT INTO user_gateway_right (id,user_id,gw_id) VALUES
			(1,1,1);
		`)
		Expect(err).Should(BeNil())

		router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), w)
	})

	AfterEach(func() {
		models.Close()
	})

	Describe("GetSolarEnergyInfo", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/solar/energy-info", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?startTime=%s", prefixURL, UtStartTime)
				expectedResponseData := services.SolarEnergyInfoResponse{
					PvProducedLifetimeEnergyACDiff:        10,
					LoadPvConsumedEnergyPercentAC:         100,
					LoadPvConsumedLifetimeEnergyACDiff:    10,
					BatteryPvConsumedEnergyPercentAC:      0,
					BatteryPvConsumedLifetimeEnergyACDiff: 0,
					GridPvConsumedEnergyPercentAC:         150,
					GridPvConsumedLifetimeEnergyACDiff:    15,
					PvEnergyCostSavingsSum:                170,
					PvCo2SavingsSum:                       38,
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
				var data services.SolarEnergyInfoResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/solar/energy-info", fixtures.UtGateway.UUID)
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

	Describe("GetSolarPowerState", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/solar/power-state", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", UtStartTime, UtEndTime)
				expectedTimestamps := []int{1659543000, 1659549599, 1659553199, 1659556799, 1659557100}
				expectedPvAveragePowerACs := []float32{40, 0, 0, 0, 40}
				expectedResponseData := services.SolarPowerStateResponse{
					Timestamps:        expectedTimestamps,
					PvAveragePowerACs: expectedPvAveragePowerACs,
					OnPeakTime:        expectedOnPeakTime,
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
				var data services.SolarPowerStateResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/solar/power-state", fixtures.UtGateway.UUID)
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

	Describe("GetBatteryEnergyInfo", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/energy-info", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?startTime=%s", prefixURL, UtStartTime)
				expectedResponseData := services.BatteryEnergyInfoResponse{
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
					Voltage:                             51.2,
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
				var data services.BatteryEnergyInfoResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/energy-info", fixtures.UtGateway.UUID)
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

	Describe("GetBatteryPowerState", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/power-state", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", UtStartTime, UtEndTime)
				expectedTimestamps := []int{1659543000, 1659549599, 1659553199, 1659556799, 1659557100}
				expectedBatteryAveragePowerACs := []float32{-3.5, 0, 0, 0, -7}
				expectedResponseData := services.BatteryPowerStateResponse{
					Timestamps:             expectedTimestamps,
					BatteryAveragePowerACs: expectedBatteryAveragePowerACs,
					OnPeakTime:             expectedOnPeakTime,
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
				var data services.BatteryPowerStateResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail-invalid resolution param", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/power-state", fixtures.UtGateway.UUID)
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

		Context("fail-invalid startTime param", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/power-state", fixtures.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", "xxx", UtEndTime)
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

		Context("fail-invalid endTime param", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/power-state", fixtures.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", UtStartTime, UtStartTime)
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

		Context("fail-invalid period endTime param", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/power-state", fixtures.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", UtStartTime, UtStartTime)
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

		Context("fail-no resolution param", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/power-state", fixtures.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?startTime=%s&endTime=%s", prefixURL, UtStartTime, UtEndTime)
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

	Describe("GetBatteryChargeVoltageState", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/charge-voltage-state", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", UtStartTime, UtEndTime)
				expectedTimestamps := []int{1659543000, 1659549599, 1659553199, 1659556799, 1659557100}
				expectedBatterySoCs := []float32{80, 0, 0, 0, 160}
				expectedBatteryVoltages := []float32{28, 0, 0, 0, 56}
				expectedResponseData := services.BatteryChargeVoltageStateResponse{
					Timestamps:      expectedTimestamps,
					BatterySoCs:     expectedBatterySoCs,
					BatteryVoltages: expectedBatteryVoltages,
					OnPeakTime:      expectedOnPeakTime,
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
				var data services.BatteryChargeVoltageStateResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/battery/charge-voltage-state", fixtures.UtGateway.UUID)
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

	Describe("GetGridEnergyInfo", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/grid/energy-info", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?startTime=%s", prefixURL, UtStartTime)
				expectedResponseData := services.GridEnergyInfoResponse{
					GridConsumedLifetimeEnergyACDiff: 15,
					GridProducedLifetimeEnergyACDiff: 25,
					GridLifetimeEnergyACDiff:         10,
					GridLifetimeEnergyACDiffOfMonth:  10,
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
				var data services.GridEnergyInfoResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/grid/energy-info", fixtures.UtGateway.UUID)
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

	Describe("GetGridPowerState", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/grid/power-state", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?resolution=%s&startTime=%s&endTime=%s", prefixURL, "hour", UtStartTime, UtEndTime)
				expectedTimestamps := []int{1659543000, 1659549599, 1659553199, 1659556799, 1659557100}
				expectedGridAveragePowerACs := []float32{50, 0, 0, 0, 50}
				expectedResponseData := services.GridPowerStateResponse{
					Timestamps:          expectedTimestamps,
					GridAveragePowerACs: expectedGridAveragePowerACs,
					OnPeakTime:          expectedOnPeakTime,
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
				var data services.GridPowerStateResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/grid/power-state", fixtures.UtGateway.UUID)
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
