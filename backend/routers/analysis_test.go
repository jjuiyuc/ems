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

var _ = Describe("Analysis", func() {
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

	Describe("GetEnergyDistributionInfo", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/energy-distribution-info", fixtures.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?startTime=%s&endTime=%s", prefixURL, UtStartTime, UtEndTime)
				expectedResponseData := services.EnergyDistributionInfoResponse{
					AllProducedLifetimeEnergyACDiff:     0,
					PvProducedEnergyPercentAC:           0,
					GridProducedEnergyPercentAC:         0,
					BatteryProducedEnergyPercentAC:      0,
					PvProducedLifetimeEnergyACDiff:      0,
					GridProducedLifetimeEnergyACDiff:    0,
					BatteryProducedLifetimeEnergyACDiff: 250,
					AllConsumedLifetimeEnergyACDiff:     0,
					LoadConsumedEnergyPercentAC:         0,
					GridConsumedEnergyPercentAC:         0,
					BatteryConsumedEnergyPercentAC:      0,
					LoadConsumedLifetimeEnergyACDiff:    0,
					GridConsumedLifetimeEnergyACDiff:    0,
					BatteryConsumedLifetimeEnergyACDiff: 250,
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
				prefixURL := fmt.Sprintf("/api/%s/devices/energy-distribution-info", fixtures.UtGateway.UUID)
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
})
