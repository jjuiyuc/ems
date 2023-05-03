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

var _ = Describe("Economics", func() {
	const (
		UtStartTime = "2022-07-31T16:00:00.000Z"
		UtEndTime   = "2022-08-02T16:00:00.000Z"
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

	Describe("GetTimeOfUseEnergyCost", func() {
		Context("success", func() {
			It("should be ok", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/tou/energy-cost", testdata.UtGateway.UUID)
				seedUtURL := fmt.Sprintf("%s?startTime=%s&endTime=%s", prefixURL, UtStartTime, UtEndTime)
				expectedResponseData := services.TimeOfUseEnergyCostResponse{
					EnergyCosts: services.EnergyCostsInfo{
						PreUbiikThisMonth:             798,
						PostUbiikThisMonth:            636,
						PreUbiikLastMonth:             0,
						PostUbiikLastMonth:            0,
						PreUbiikTheSameMonthLastYear:  0,
						PostUbiikTheSameMonthLastYear: 0,
					},
					EnergyDailyCosts: services.EnergyCostsDailyInfo{
						Timestamps:                    []int{1659369555, 1659455970},
						PreUbiikThisMonth:             []int{411, 387},
						PostUbiikThisMonth:            []int{340, 296},
						PreUbiikLastMonth:             []int{0, 0},
						PostUbiikLastMonth:            []int{0, 0},
						PreUbiikTheSameMonthLastYear:  []int{0, 0},
						PostUbiikTheSameMonthLastYear: []int{0, 0},
						SavedThisMonth:                []int{71, 91},
						SavedLastMonth:                []int{0, 0},
						SavedTheSameMonthLastYear:     []int{0, 0},
					},
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
				var data services.TimeOfUseEnergyCostResponse
				err = json.Unmarshal(dataJSON, &data)
				Expect(err).Should(BeNil())
				Expect(data).To(Equal(expectedResponseData))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				prefixURL := fmt.Sprintf("/api/%s/devices/tou/energy-cost", testdata.UtGateway.UUID)
				seedUtInvalidParamsURL := fmt.Sprintf("%s?startTime=%s&endTime=%s", prefixURL, "xxx", UtEndTime)
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
