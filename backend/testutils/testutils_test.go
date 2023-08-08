package testutils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/kafka"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/testutils/testdata"
)

type TestutilsSuite struct {
	suite.Suite
	repo       *repository.Repository
	seedUtData map[string]interface{}
}

func Test_Testutils(t *testing.T) {
	suite.Run(t, &TestutilsSuite{})
}

func (s *TestutilsSuite) SetupSuite() {
	config.Init(GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()

	// Truncate & seed data
	err := SeedUtGroupAndUser(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.seedUtData = map[string]interface{}{
		"gwID":      "U00001",
		"timestamp": 1653964322,
	}

	s.repo = repository.NewRepository(db)
}

func (s *TestutilsSuite) Test_SeedUtGroupAndUser() {
	_, err := s.repo.User.GetUserByUsername(testdata.UtUser.Username)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
}

func (s *TestutilsSuite) Test_GetAuthorization() {
	seedUtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNjU1MzcxNjU2LCJpc3MiOiJkZXJlbXMifQ.VLBUMzihKZBJQ5zw845bSyokIEy-gQV1kS5w0g_dDdo"

	data := GetAuthorization(seedUtToken)
	s.Equalf(fmt.Sprintf("Bearer %s", seedUtToken), data, e.ErrNewMessageNotEqual.Error())
}

func (s *TestutilsSuite) Test_CopyMap() {
	data := CopyMap(s.seedUtData)
	s.Equalf(s.seedUtData, data, e.ErrNewMessageNotEqual.Error())
}

func (s *TestutilsSuite) Test_GetMockConsumerMessage() {
	seedUtTopic := kafka.ReceiveLocalCCData

	dataJSON, err := json.Marshal(s.seedUtData)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	msg, err := GetMockConsumerMessage(s.T(), seedUtTopic, dataJSON)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.Equalf(seedUtTopic, msg.Topic, e.ErrNewMessageNotEqual.Error())
}
