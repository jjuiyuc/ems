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
	s.Require().NoError(err)
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
		"gwID":                            fixtures.UtGateway.UUID,
		"timestamp":                       1659340800,
		"gridIsPeakShaving":               0,
		"loadGridAveragePowerAC":          10,
		"batteryGridAveragePowerAC":       0,
		"gridContractPowerAC":             15,
		"loadPvAveragePowerAC":            20,
		"loadBatteryAveragePowerAC":       0,
		"batterySoC":                      80,
		"batteryProducedAveragePowerAC":   20,
		"batteryConsumedAveragePowerAC":   0,
		"batteryChargingFrom":             "Solar",
		"batteryDischargingTo":            "",
		"pvAveragePowerAC":                40,
		"loadAveragePowerAC":              30,
		"loadLinks":                       seedUtLoadLinks,
		"gridLinks":                       seedUtGridLinks,
		"pvLinks":                         seedUtPvLinks,
		"batteryLinks":                    seedUtBatteryLinks,
		"batteryPvAveragePowerAC":         20,
		"gridPvAveragePowerAC":            0,
		"gridProducedAveragePowerAC":      10,
		"gridConsumedAveragePowerAC":      0,
		"batteryLifetimeOperationCycles":  8,
		"batteryProducedLifetimeEnergyAC": 250,
		"batteryConsumedLifetimeEnergyAC": 250,
		"batteryAveragePowerAC":           -3.5,
		"batteryVoltage":                  28,
	}
}

func (s *LocalCCWorkerSuite) TearDownSuite() {
	// Delete test data in cc_data_log table
	_, err := s.db.Exec(`
		DELETE FROM cc_data_log WHERE id = 3 OR id = 4;
	`)
	s.Require().NoError(err)
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
		currentLogCount, err := s.repo.CCData.GetCCDataLogCount()
		s.Require().NoError(err)
		saveErr := s.handler.saveLocalCCData(msg.Value)
		saveLogErr := s.handler.saveLocalCCDataLog(msg.Value)

		switch tt.name {
		case "saveLocalCCData", "saveLocalCCDataNewGW":
			s.Require().NoError(saveErr)
			s.Require().NoError(saveLogErr)
			updatedCount, err := s.repo.CCData.GetCCDataCount()
			s.Require().NoError(err)
			s.Equal(currentCount+1, updatedCount)
			updatedCount, err = s.repo.CCData.GetCCDataLogCount()
			s.Require().NoError(err)
			s.Equal(currentLogCount+1, updatedCount)
		case "saveLocalCCDataNoGWID":
			s.Equal(e.ErrNewKeyNotExist(gwID).Error(), saveErr.Error())
			s.Equal(e.ErrNewKeyNotExist(gwID).Error(), saveLogErr.Error())
		case "saveLocalCCDataNoTimestamp":
			s.Equal(e.ErrNewKeyNotExist(timestamp).Error(), saveErr.Error())
			s.Equal(e.ErrNewKeyNotExist(timestamp).Error(), saveLogErr.Error())
		}
	}
}
