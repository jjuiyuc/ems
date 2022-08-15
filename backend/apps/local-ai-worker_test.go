package apps

import (
	"database/sql"
	"encoding/json"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/kafka"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/testutils"
	"der-ems/testutils/fixtures"
)

type LocalAIWorkerSuite struct {
	suite.Suite
	seedUtTopic       string
	seedUtLocalAIData map[string]interface{}
	db                *sql.DB
	repo              *repository.Repository
	handler           localAIConsumerHandler
}

func Test_LocalAIWorker(t *testing.T) {
	suite.Run(t, &LocalAIWorkerSuite{})
}

func (s *LocalAIWorkerSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	repo := repository.NewRepository(db)
	handler := localAIConsumerHandler{
		cfg:  cfg,
		repo: repo,
	}

	s.seedUtTopic = kafka.ReceiveLocalAIData
	s.db = db
	s.repo = repo
	s.handler = handler

	// Truncate & seed data
	err := testutils.SeedUtCustomerAndGateway(db)
	s.Require().NoError(err)
	// Mock seedUtLocalAIData data
	s.seedUtLocalAIData = map[string]interface{}{
		"gwID":      fixtures.UtGateway.UUID,
		"timestamp": 1659340800,
		"value":     16,
		"type":      "batterySchedule",
	}
}

func (s *LocalAIWorkerSuite) TearDownSuite() {
	models.Close()
}

func (s *LocalAIWorkerSuite) Test_SaveLocalAIData() {
	const (
		gwID      = "gwID"
		timestamp = "timestamp"
	)
	type args struct {
		Msg map[string]interface{}
	}

	// Modify seedUtLocalAIData data
	seedUtDataNewGW := testutils.CopyMap(s.seedUtLocalAIData)
	seedUtDataNewGW[gwID] = "U00000"
	seedUtDataNoGWID := testutils.CopyMap(s.seedUtLocalAIData)
	delete(seedUtDataNoGWID, gwID)
	seedUtDataNoTimestamp := testutils.CopyMap(s.seedUtLocalAIData)
	delete(seedUtDataNoTimestamp, timestamp)
	seedUtDataGWIDUnexpectedValue := testutils.CopyMap(s.seedUtLocalAIData)
	seedUtDataGWIDUnexpectedValue[gwID] = nil
	seedUtDataTimestampUnexpectedValue := testutils.CopyMap(s.seedUtLocalAIData)
	seedUtDataTimestampUnexpectedValue[timestamp] = nil

	tests := []struct {
		name string
		args args
	}{
		{
			name: "saveLocalAIData",
			args: args{
				Msg: s.seedUtLocalAIData,
			},
		},
		{
			name: "saveLocalAIDataNewGW",
			args: args{
				Msg: seedUtDataNewGW,
			},
		},
		{
			name: "saveLocalAIDataNoGWID",
			args: args{
				Msg: seedUtDataNoGWID,
			},
		},
		{
			name: "saveLocalAIDataNoTimestamp",
			args: args{
				Msg: seedUtDataNoTimestamp,
			},
		},
		{
			name: "saveLocalAIDataGWIDUnexpectedValue",
			args: args{
				Msg: seedUtDataGWIDUnexpectedValue,
			},
		},
		{
			name: "saveLocalAIDataTimestampUnexpectedValue",
			args: args{
				Msg: seedUtDataTimestampUnexpectedValue,
			},
		},
		{
			name: "saveLocalAIDataEmptyInput",
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.name)
		if tt.name == "saveLocalAIDataEmptyInput" {
			err := s.handler.saveLocalAIData(nil)
			s.Require().Error(e.ErrNewUnexpectedJSONInput, err)
			continue
		}

		dataJSON, err := json.Marshal(tt.args.Msg)
		s.Require().NoError(err)
		msg, err := testutils.GetMockConsumerMessage(s.T(), s.seedUtTopic, dataJSON)
		s.Require().NoError(err)
		s.Equal(s.seedUtTopic, msg.Topic)

		currentCount, err := s.repo.AIData.GetAIDataCount()
		s.Require().NoError(err)
		saveErr := s.handler.saveLocalAIData(msg.Value)

		switch tt.name {
		case "saveLocalAIData", "saveLocalAIDataNewGW":
			s.Require().NoError(saveErr)
			updatedCount, err := s.repo.AIData.GetAIDataCount()
			s.Require().NoError(err)
			s.Equal(currentCount+1, updatedCount)
		case "saveLocalAIDataNoGWID":
			s.Equal(e.ErrNewKeyNotExist(gwID).Error(), saveErr.Error())
		case "saveLocalAIDataNoTimestamp":
			s.Equal(e.ErrNewKeyNotExist(timestamp).Error(), saveErr.Error())
		case "saveLocalAIDataGWIDUnexpectedValue":
			s.Equal(e.ErrNewKeyUnexpectedValue(gwID).Error(), saveErr.Error())
		case "saveLocalAIDataTimestampUnexpectedValue":
			s.Equal(e.ErrNewKeyUnexpectedValue(timestamp).Error(), saveErr.Error())
		}
	}
}
