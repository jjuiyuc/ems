package routers

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

var _ = Describe("User", func() {
	var (
		db     *sql.DB
		repo   *repository.Repository
		router *gin.Engine
		token  string
		err    error
	)

	BeforeEach(func() {
		config.Init(testutils.GetConfigDir(), "ut.yaml")
		cfg := config.GetConfig()
		models.Init(cfg)
		db = models.GetDB()

		repo = repository.NewRepository(db)
		w := &APIWorker{
			Services: services.NewServices(cfg, repo),
		}

		token, err = utils.GenerateToken(testutils.SeedUtClaims())
		Expect(err).Should(BeNil())

		router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), initPolicy(testutils.GetConfigDir()), w)
	})

	AfterEach(func() {
		err = testutils.SeedUtUser(db)
		Expect(err).Should(BeNil())
		models.Close()
	})

	Describe("UpdateName", func() {
		Context("success", func() {
			It("should be ok", func() {
				seedUtURL := "/api/users/name"
				seedUtArg := &PersonalName{
					Name: "test",
				}
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusOK,
					WantRv: app.Response{
						Code: e.Success,
						Msg:  "ok",
					},
				}
				payloadBuf, _ := json.Marshal(seedUtArg)
				testutils.GinkgoAssertRequest(tt, router, "PUT", bytes.NewBuffer(payloadBuf))
				user, err := repo.User.GetUserByUserID(nil, testdata.UtUser.ID)
				Expect(err).Should(BeNil())
				Expect(user.Name.String).To(Equal(seedUtArg.Name))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				seedUtURL := "/api/users/name"
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusBadRequest,
					WantRv: app.Response{
						Code: e.InvalidParams,
						Msg:  "invalid parameters",
					},
				}
				testutils.GinkgoAssertRequest(tt, router, "PUT", nil)
			})
		})
	})

	Describe("UpdatePassword", func() {
		Context("success", func() {
			It("should be ok", func() {
				seedUtURL := "/api/users/password"
				seedUtPassword := "abc123def456"
				seedUtArg := &PersonalPassword{
					CurrentPassword: testdata.UtUser.Password,
					NewPassword:     seedUtPassword,
				}
				seedUtAfterArg := &PersonalPassword{
					CurrentPassword: seedUtPassword,
					NewPassword:     testdata.UtUser.Password,
				}
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusOK,
					WantRv: app.Response{
						Code: e.Success,
						Msg:  "ok",
					},
				}
				payloadBuf, _ := json.Marshal(seedUtArg)
				testutils.GinkgoAssertRequest(tt, router, "PUT", bytes.NewBuffer(payloadBuf))
				user, err := repo.User.GetUserByUserID(nil, testdata.UtUser.ID)
				Expect(err).Should(BeNil())
				err = utils.ComparePassword(seedUtPassword, user.Password)
				Expect(err).Should(BeNil())

				payloadBuf, _ = json.Marshal(seedUtAfterArg)
				testutils.GinkgoAssertRequest(tt, router, "PUT", bytes.NewBuffer(payloadBuf))
			})
		})

		Context("fail", func() {
			It("should return fail", func() {
				seedUtURL := "/api/users/password"
				seedUtPassword := "abc123def456"
				seedUtArg := &PersonalPassword{
					CurrentPassword: seedUtPassword,
					NewPassword:     seedUtPassword,
				}
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusUnauthorized,
					WantRv: app.Response{
						Code: e.ErrAuthPasswordNotMatch,
						Msg:  "fail",
					},
				}
				payloadBuf, _ := json.Marshal(seedUtArg)
				testutils.GinkgoAssertRequest(tt, router, "PUT", bytes.NewBuffer(payloadBuf))
			})
		})

		Context("fail", func() {
			It("should return invalid parameters", func() {
				seedUtURL := "/api/users/password"
				tt := testutils.TestInfo{
					Token:      token,
					URL:        seedUtURL,
					WantStatus: http.StatusBadRequest,
					WantRv: app.Response{
						Code: e.InvalidParams,
						Msg:  "invalid parameters",
					},
				}
				testutils.GinkgoAssertRequest(tt, router, "PUT", nil)
			})
		})
	})
})
