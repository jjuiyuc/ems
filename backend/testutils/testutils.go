package testutils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"der-ems/internal/app"
	"der-ems/internal/utils"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/testutils/testdata"
)

// TestInfo godoc
type TestInfo struct {
	Name       string
	Token      string
	URL        string
	WantStatus int
	WantRv     app.Response
}

// GetConfigDir godoc
func GetConfigDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "config")
}

// SeedUtUser godoc
func SeedUtUser(db *sql.DB) (err error) {
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return
	}
	_, err = db.Exec("truncate table user")
	if err != nil {
		return
	}
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		return
	}
	hashedPassword, err := utils.CreateHashedPassword(testdata.UtUser.Password)
	if err != nil {
		return
	}
	user := &deremsmodels.User{
		Username:       testdata.UtUser.Username,
		GroupID:        testdata.UtUser.GroupID,
		Password:       hashedPassword,
		ExpirationDate: testdata.UtUser.ExpirationDate,
		CreatedAt:      testdata.UtUser.CreatedAt,
		UpdatedAt:      testdata.UtUser.UpdatedAt,
	}
	err = user.Insert(db, boil.Infer())
	return
}

// SeedUtLocationAndGateway godoc
func SeedUtLocationAndGateway(db *sql.DB) (err error) {
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return
	}
	_, err = db.Exec("truncate table gateway")
	if err != nil {
		return
	}
	_, err = db.Exec("truncate table location")
	if err != nil {
		return
	}
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		return
	}
	location := testdata.UtLocation
	err = location.Insert(db, boil.Infer())
	if err != nil {
		return
	}
	gateway := testdata.UtGateway
	err = gateway.Insert(db, boil.Infer())
	return
}

// SeedUtClaims godoc
func SeedUtClaims() (claims utils.Claims) {
	claims = utils.Claims{
		UserID:  testdata.UtUser.ID,
		GroupType: testdata.UtGroupType,
	}
	return
}

// GetAuthorization godoc
func GetAuthorization(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

// CopyMap godoc
func CopyMap(originalMap map[string]interface{}) (targetMap map[string]interface{}) {
	targetMap = make(map[string]interface{})
	for key, value := range originalMap {
		targetMap[key] = value
	}
	return
}

// GetMockConsumerMessage godoc
func GetMockConsumerMessage(t *testing.T, seedUtTopic string, seedUtData []byte) (testMsg *sarama.ConsumerMessage, err error) {
	consumer := mocks.NewConsumer(t, mocks.NewTestConfig())
	defer func() {
		if err := consumer.Close(); err != nil {
			log.WithFields(log.Fields{
				"caused-by": "consumer.Close",
				"err":       err,
			}).Error()
		}
	}()

	var seedUtPartition int32 = 0
	consumer.ExpectConsumePartition(seedUtTopic, seedUtPartition, sarama.OffsetOldest).YieldMessage(&sarama.ConsumerMessage{Value: []byte(seedUtData)})

	test, err := consumer.ConsumePartition(seedUtTopic, seedUtPartition, sarama.OffsetOldest)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "consumer.ConsumePartition",
			"err":       err,
		}).Fatal()
	}
	testMsg = <-test.Messages()

	log.WithFields(log.Fields{
		"topic":     testMsg.Topic,
		"partition": testMsg.Partition,
		"offset":    testMsg.Offset,
		"value":     string(testMsg.Value),
	}).Info("consuming")
	return
}

// AssertRequest godoc
func AssertRequest(tt TestInfo, a *require.Assertions, router *gin.Engine, method string, body io.Reader) (rvData interface{}) {
	req, err := http.NewRequest(method, fmt.Sprintf(tt.URL), body)
	a.NoError(err)
	if tt.Token != "" {
		req.Header.Set("Authorization", GetAuthorization(tt.Token))
	}
	rv := httptest.NewRecorder()
	router.ServeHTTP(rv, req)
	a.Equal(tt.WantStatus, rv.Code)

	var res app.Response
	err = json.Unmarshal([]byte(rv.Body.String()), &res)
	a.NoError(err)
	a.Equal(tt.WantRv.Code, res.Code)
	a.Equal(tt.WantRv.Msg, res.Msg)
	return res.Data
}

// GinkgoAssertRequest godoc
func GinkgoAssertRequest(tt TestInfo, router *gin.Engine, method string, body io.Reader) (rvData interface{}) {
	req, err := http.NewRequest(method, fmt.Sprintf(tt.URL), body)
	Expect(err).Should(BeNil())
	if tt.Token != "" {
		req.Header.Set("Authorization", GetAuthorization(tt.Token))
	}
	rv := httptest.NewRecorder()
	router.ServeHTTP(rv, req)
	Expect(rv.Code).To(Equal(tt.WantStatus))

	var res app.Response
	err = json.Unmarshal([]byte(rv.Body.String()), &res)
	Expect(err).Should(BeNil())
	Expect(res.Code).To(Equal(tt.WantRv.Code))
	Expect(res.Msg).To(Equal(tt.WantRv.Msg))
	return res.Data
}
