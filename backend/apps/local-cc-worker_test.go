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

type LocalCCWorkerSuite struct {
	suite.Suite
	seedUtTopic       string
	seedUtLocalCCData map[string]interface{}
	db                *sql.DB
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
	s.db = db
	s.repo = repo
	s.handler = handler

	// Truncate & seed data
	err := testutils.SeedUtCustomerAndGateway(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	// Mock seedUtLocalCCData data
	seedUtLoadLinks := map[string]int{
		"grid":    0,
		"battery": 0,
		"pv":      0,
	}
	seedUtGridLinks := map[string]int{
		"load":    1,
		"battery": 0,
		"pv":      0,
	}
	seedUtPvLinks := map[string]int{
		"load":    1,
		"battery": 1,
		"grid":    0,
	}
	seedUtBatteryLinks := map[string]int{
		"load": 0,
		"pv":   0,
		"grid": 0,
	}
	s.seedUtLocalCCData = map[string]interface{}{
		"gwID":                              fixtures.UtGateway.UUID,
		"timestamp":                         1659340800,
		"gridIsPeakShaving":                 0,
		"loadGridAveragePowerAC":            10,
		"batteryGridAveragePowerAC":         0,
		"gridContractPowerAC":               15,
		"loadPvAveragePowerAC":              20,
		"loadBatteryAveragePowerAC":         0,
		"batterySoC":                        80,
		"batteryProducedAveragePowerAC":     20,
		"batteryConsumedAveragePowerAC":     0,
		"batteryChargingFrom":               "Solar",
		"batteryDischargingTo":              "",
		"pvAveragePowerAC":                  40,
		"loadAveragePowerAC":                30,
		"loadLinks":                         seedUtLoadLinks,
		"gridLinks":                         seedUtGridLinks,
		"pvLinks":                           seedUtPvLinks,
		"batteryLinks":                      seedUtBatteryLinks,
		"batteryPvAveragePowerAC":           20,
		"gridPvAveragePowerAC":              0,
		"gridProducedAveragePowerAC":        10,
		"gridConsumedAveragePowerAC":        0,
		"batteryLifetimeOperationCycles":    8,
		"batteryProducedLifetimeEnergyAC":   250,
		"batteryConsumedLifetimeEnergyAC":   250,
		"batteryAveragePowerAC":             -3.5,
		"batteryVoltage":                    28,
		"allProducedLifetimeEnergyAC":       80,
		"pvProducedLifetimeEnergyAC":        40,
		"gridProducedLifetimeEnergyAC":      10,
		"allConsumedLifetimeEnergyAC":       80,
		"loadConsumedLifetimeEnergyAC":      40,
		"gridConsumedLifetimeEnergyAC":      10,
		"gridAveragePowerAC":                3.7,
		"batteryLifetimeEnergyAC":           0.5,
		"gridLifetimeEnergyAC":              0.5,
		"loadSelfConsumedLifetimeEnergyAC":  60,
		"gridPowerCost":                     0,
		"gridPowerCostSavings":              0,
		"loadPvConsumedLifetimeEnergyAC":    40,
		"batteryPvConsumedLifetimeEnergyAC": 15,
		"gridPvConsumedLifetimeEnergyAC":    50,
		"pvEnergyCostSavings":               85,
		"pvCo2Savings":                      19,
	}
}

func (s *LocalCCWorkerSuite) TearDownSuite() {
	// Delete this time test data in cc_data_log table (Default is 3 records in DB, that's way delete index 4 and 5)
	_, err := s.db.Exec(`
		DELETE FROM cc_data_log WHERE id = 4 OR id = 5;
	`)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	models.Close()
}

func (s *LocalCCWorkerSuite) Test_SaveLocalCCDataAndLog() {
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
	seedUtDataGWIDUnexpectedValue := testutils.CopyMap(s.seedUtLocalCCData)
	seedUtDataGWIDUnexpectedValue[gwID] = nil
	seedUtDataTimestampUnexpectedValue := testutils.CopyMap(s.seedUtLocalCCData)
	seedUtDataTimestampUnexpectedValue[timestamp] = nil

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
			name: "saveLocalCCDataGWIDUnexpectedValue",
			args: args{
				Msg: seedUtDataGWIDUnexpectedValue,
			},
		},
		{
			name: "saveLocalCCDataTimestampUnexpectedValue",
			args: args{
				Msg: seedUtDataTimestampUnexpectedValue,
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
			s.Equalf(e.ErrNewUnexpectedJSONInput.Error(), err.Error(), e.ErrNewMessageNotEqual.Error())
			continue
		}

		dataJSON, err := json.Marshal(tt.args.Msg)
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		msg, err := testutils.GetMockConsumerMessage(s.T(), s.seedUtTopic, dataJSON)
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		s.Equalf(s.seedUtTopic, msg.Topic, e.ErrNewMessageNotEqual.Error())

		currentCount, err := s.repo.CCData.GetCCDataCount()
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		currentLogCount, err := s.repo.CCData.GetCCDataLogCount()
		s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
		saveErr := s.handler.saveLocalCCData(msg.Value)
		saveLogErr := s.handler.saveLocalCCDataLog(msg.Value)

		switch tt.name {
		case "saveLocalCCData", "saveLocalCCDataNewGW":
			s.Require().NoErrorf(saveErr, e.ErrNewMessageReceivedUnexpectedErr.Error())
			s.Require().NoErrorf(saveLogErr, e.ErrNewMessageReceivedUnexpectedErr.Error())
			updatedCount, err := s.repo.CCData.GetCCDataCount()
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			s.Equalf(currentCount+1, updatedCount, e.ErrNewMessageNotEqual.Error())
			updatedCount, err = s.repo.CCData.GetCCDataLogCount()
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			s.Equalf(currentLogCount+1, updatedCount, e.ErrNewMessageNotEqual.Error())
		case "saveLocalCCDataNoGWID":
			s.Equalf(e.ErrNewKeyNotExist(gwID).Error(), saveErr.Error(), e.ErrNewMessageNotEqual.Error())
			s.Equalf(e.ErrNewKeyNotExist(gwID).Error(), saveLogErr.Error(), e.ErrNewMessageNotEqual.Error())
		case "saveLocalCCDataNoTimestamp":
			s.Equalf(e.ErrNewKeyNotExist(timestamp).Error(), saveErr.Error(), e.ErrNewMessageNotEqual.Error())
			s.Equalf(e.ErrNewKeyNotExist(timestamp).Error(), saveLogErr.Error(), e.ErrNewMessageNotEqual.Error())
		case "saveLocalCCDataGWIDUnexpectedValue":
			s.Equalf(e.ErrNewKeyUnexpectedValue(gwID).Error(), saveErr.Error(), e.ErrNewMessageNotEqual.Error())
			s.Equalf(e.ErrNewKeyUnexpectedValue(gwID).Error(), saveLogErr.Error(), e.ErrNewMessageNotEqual.Error())
		case "saveLocalCCDataTimestampUnexpectedValue":
			s.Equalf(e.ErrNewKeyUnexpectedValue(timestamp).Error(), saveErr.Error(), e.ErrNewMessageNotEqual.Error())
			s.Equalf(e.ErrNewKeyUnexpectedValue(timestamp).Error(), saveLogErr.Error(), e.ErrNewMessageNotEqual.Error())
		}
	}
}
