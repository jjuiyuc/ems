package apps

import (
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

type LocalCCWorkerSuite struct {
	suite.Suite
	seedUtTopic       string
	seedUtLocalCCData map[string]interface{}
	repo              *repository.Repository
	handler           localCCConsumerHandler
}

func Test_LocalCCWorker(t *testing.T) {
	suite.Run(t, &LocalCCWorkerSuite{})
}

func (s *LocalCCWorkerSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	repo := repository.NewRepository(db)
	handler := localCCConsumerHandler{
		cfg:  cfg,
		repo: repo,
	}

	s.seedUtTopic = kafka.ReceiveLocalCCData
	s.repo = repo
	s.handler = handler

	// Truncate & seed data
	err := testutils.SeedUtCustomerAndGateway(db)
	s.Require().NoError(err)
	// Mock seedUtLocalCCData data
	s.seedUtLocalCCData = map[string]interface{}{
		"gwID":                            fixtures.UtGateway.UUID,
		"timestamp":                       1653964322,
		"timestampDelta":                  5,
		"gridInstantaneousPowerAC":        20.15,
		"gridAveragePowerAC":              18.20,
		"gridProducedEnergyAC":            5.73,
		"gridConsumedEnergyAC":            0.000,
		"gridLifetimeEnergyAC":            76947.0098,
		"gridDeltaLifetimeEnergyAC":       12.10,
		"loadInstantaneousPowerAC":        -12.0928,
		"loadAveragePowerAC":              -2.000,
		"loadBatteryAveragePowerAC":       5.931,
		"loadPvAveragePowerAC":            3.3,
		"loadBatteryConsumedEnergyAC":     56,
		"loadPvConsumedEnergyAC":          1,
		"loadConsumedEnergyAC":            2,
		"loadLifetimeEnergyAC":            421,
		"batteryInstantaneousPowerDC":     10,
		"batteryInstantaneousPowerAC":     11,
		"batteryAveragePowerAC":           10.3,
		"batteryProducedEnergyDC":         2,
		"batteryProducedEnergyAC":         2.1,
		"batteryConsumedEnergyDC":         0.0,
		"batteryConsumedEnergyAC":         0.0,
		"batteryLifetimeEnergyDC":         382.2,
		"batteryDeltaLifetimeEnergyAC":    2.1,
		"batteryDeltaLifetimeEnergyDC":    2.1,
		"batteryLifetimeEnergyAC":         383,
		"batteryCurrentStoredEnergy":      31.3,
		"batteryDeltaCurrentStoredEnergy": 2.2,
		"batteryNameplateEnergy":          100,
		"batteryChargeEfficiency":         79.34,
		"batteryDischargeEfficiency":      75.53,
		"batteryInverterEfficiencyDCAC":   95.3,
		"batteryInverterEfficiencyACDC":   97.7,
		"batteryVoltage":                  800.8,
		"batterySoC":                      87.1,
		"batterySoH":                      99.2,
		"pvInstantaneousPowerDC":          5.9,
		"pvInstantaneousPowerAC":          4.9,
		"pvAveragePowerAC":                4.7,
		"pvProducedEnergyDC":              6.7,
		"pvProducedEnergyAC":              5.6,
		"pvLifetimeEnergyDC":              67099,
		"pvLifetimeEnergyAC":              67001,
		"pvDeltaLifetimeEnergyAC":         6.97,
		"pvInverterEfficiencyDCAC":        94,
		"allProducedEnergyAC":             83772.28,
		"allConsumedEnergyAC":             28726.28,
	}
}

func (s *LocalCCWorkerSuite) TearDownSuite() {
	models.Close()
}

func (s *LocalCCWorkerSuite) Test_SaveLocalCCData() {
	const (
		gwID      = "gwID"
		timestamp = "timestamp"
	)
	type args struct {
		Msg map[string]interface{}
	}

	// Modify seedUtLocalCCData data
	seedUtDataNewGW := testutils.CopyMap(s.seedUtLocalCCData)
	seedUtDataNewGW[gwID] = "U00000"
	seedUtDataNoGWID := testutils.CopyMap(s.seedUtLocalCCData)
	delete(seedUtDataNoGWID, gwID)
	seedUtDataNoTimestamp := testutils.CopyMap(s.seedUtLocalCCData)
	delete(seedUtDataNoTimestamp, timestamp)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "saveLocalCCData",
			args: args{
				Msg: s.seedUtLocalCCData,
			},
		},
		{
			name: "saveLocalCCDataNewGW",
			args: args{
				Msg: seedUtDataNewGW,
			},
		},
		{
			name: "saveLocalCCDataNoGWID",
			args: args{
				Msg: seedUtDataNoGWID,
			},
		},
		{
			name: "saveLocalCCDataNoTimestamp",
			args: args{
				Msg: seedUtDataNoTimestamp,
			},
		},
		{
			name: "saveLocalCCDataEmptyInput",
		},
	}

	for _, tt := range tests {
		log.Info("test name: ", tt.name)
		if tt.name == "saveLocalCCDataEmptyInput" {
			err := s.handler.saveLocalCCData(nil)
			s.Require().Error(e.ErrNewUnexpectedJSONInput, err)
			continue
		}

		dataJSON, err := json.Marshal(tt.args.Msg)
		s.Require().NoError(err)
		msg, err := testutils.GetMockConsumerMessage(s.T(), s.seedUtTopic, dataJSON)
		s.Require().NoError(err)
		s.Equal(s.seedUtTopic, msg.Topic)

		currentCount, err := s.repo.CCData.GetCCDataCount()
		s.Require().NoError(err)
		err = s.handler.saveLocalCCData(msg.Value)

		switch tt.name {
		case "saveLocalCCData", "saveLocalCCDataNewGW":
			s.Require().NoError(err)
			updatedCount, err := s.repo.CCData.GetCCDataCount()
			s.Require().NoError(err)
			s.Equal(currentCount+1, updatedCount)
		case "saveLocalCCDataNoGWID":
			s.Equal(e.ErrNewKeyNotExist(gwID).Error(), err.Error())
		case "saveLocalCCDataNoTimestamp":
			s.Equal(e.ErrNewKeyNotExist(timestamp).Error(), err.Error())
		}
	}
}
