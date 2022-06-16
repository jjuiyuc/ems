package apps

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/kafka"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/testutils"
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

	// Truncate data
	_, err := db.Exec("TRUNCATE TABLE cc_data")
	s.Require().NoError(err)
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.Require().NoError(err)
	_, err = db.Exec("TRUNCATE TABLE gateway")
	s.Require().NoError(err)
	_, err = db.Exec("TRUNCATE TABLE customer")
	s.Require().NoError(err)
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	s.Require().NoError(err)
	// Mock seedUtLocalCCData data
	s.seedUtLocalCCData = map[string]interface{}{
		"gwID":                            "U00001",
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

	// Mock customer table
	_, err = db.Exec(`
		INSERT INTO customer (id,customer_number,field_number,weather_lat,weather_lng) VALUES
		(1,'A00001','00001',24.75,121);
	`)
	s.Require().NoError(err)

	// Mock gateway table
	_, err = db.Exec(`
		INSERT INTO gateway (id,uuid,customer_id) VALUES
		(1,'U00001',1);
	`)
	s.Require().NoError(err)
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
	testDataNewGW := testutils.CopyMap(s.seedUtLocalCCData)
	testDataNewGW[gwID] = "U00000"
	testDataNoGWID := testutils.CopyMap(s.seedUtLocalCCData)
	delete(testDataNoGWID, gwID)
	testDataNoTimestamp := testutils.CopyMap(s.seedUtLocalCCData)
	delete(testDataNoTimestamp, timestamp)

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
				Msg: testDataNewGW,
			},
		},
		{
			name: "saveLocalCCDataNoGWID",
			args: args{
				Msg: testDataNoGWID,
			},
		},
		{
			name: "saveLocalCCDataNoGWID",
			args: args{
				Msg: testDataNoGWID,
			},
		},
		{
			name: "saveLocalCCDataNoTimestamp",
			args: args{
				Msg: testDataNoTimestamp,
			},
		},
	}
	for _, tt := range tests {
		testDataJson, err := json.Marshal(tt.args.Msg)
		s.Require().NoError(err)
		testMsg, err := testutils.GetMockConsumerMessage(s.T(), s.seedUtTopic, testDataJson)
		s.Require().NoError(err)
		s.Equal(s.seedUtTopic, testMsg.Topic)

		currentCount, err := s.repo.CCData.GetCCDataCount()
		s.Require().NoError(err)
		err = s.handler.SaveLocalCCData(testMsg.Value)

		if tt.name == "saveLocalCCData" || tt.name == "saveLocalCCDataNewGW" {
			s.Require().NoError(err)
			updatedCount, err := s.repo.CCData.GetCCDataCount()
			s.Require().NoError(err)
			s.Equal(currentCount+1, updatedCount)
		} else if tt.name == "saveLocalCCDataNoGWID" {
			s.Equal(e.NewKeyNotExistError(gwID).Error(), err.Error())
		} else if tt.name == "saveLocalCCDataNoTimestamp" {
			s.Equal(e.NewKeyNotExistError(timestamp).Error(), err.Error())
		}
	}
}
