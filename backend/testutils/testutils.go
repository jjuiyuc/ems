package testutils

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"

	deremsmodels "der-ems/models/der-ems"
	"der-ems/testutils/fixtures"
)

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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(fixtures.UtUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user := &deremsmodels.User{
		Username:       fixtures.UtUser.Username,
		Password:       string(hashPassword[:]),
		ExpirationDate: fixtures.UtUser.ExpirationDate,
	}
	err = user.Insert(db, boil.Infer())
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
