package apps

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/utils"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
	"der-ems/services"
	"der-ems/testutils"
)

type BillingWorkerSuite struct {
	suite.Suite
	repo    *repository.Repository
	billing services.BillingService
}

func Test_BillingWorker(t *testing.T) {
	suite.Run(t, &BillingWorkerSuite{})
}

func (s *BillingWorkerSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	repo := repository.NewRepository(db)
	s.repo = repo
	s.billing = services.NewBillingService(repo)

	_, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.Require().NoError(err)
	_, err = db.Exec("TRUNCATE TABLE gateway")
	s.Require().NoError(err)
	_, err = db.Exec("TRUNCATE TABLE customer")
	s.Require().NoError(err)
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	s.Require().NoError(err)

	// Mock customer table
	_, err = db.Exec(`
		INSERT INTO customer (id,customer_number,field_number,address,lat,lng,weather_lat,weather_lng,timezone,tou_location_id,voltage_type,tou_type) VALUES
		(1,'00001','001','宜蘭縣五結鄉大吉五路157巷68號',24.70155508690467,121.7973398847259,24.75,121.75,'+0800',1,'Low voltage','Two-section');
	`)
	s.Require().NoError(err)

	// Mock gateway table
	_, err = db.Exec(`
		INSERT INTO gateway (id,uuid,customer_id) VALUES
		(1,'04F1FD6D9C6F64C3352285CCEAF59EE1',1);
	`)
	s.Require().NoError(err)
}

func (s *BillingWorkerSuite) TearDownSuite() {
	models.Close()
}

func (s *BillingWorkerSuite) Test_GetBillingTypeByCustomerID() {
	type args struct {
		Gateway     *deremsmodels.Gateway
		BillingType *services.BillingType
	}

	testGateway := &deremsmodels.Gateway{
		ID:         1,
		UUID:       "04F1FD6D9C6F64C3352285CCEAF59EE1",
		CustomerID: 1,
	}
	testBillingType := &services.BillingType{
		TOULocationID: 1,
		VoltageType:   "Low voltage",
		TOUType:       "Two-section",
	}

	tt := struct {
		name string
		args args
	}{
		name: "GetBillingTypeByCustomerID",
		args: args{
			Gateway:     testGateway,
			BillingType: testBillingType,
		},
	}

	gateways, err := getGateways(s.repo)
	s.Require().NoError(err)
	s.Equal(tt.args.Gateway.ID, gateways[0].ID)
	s.Equal(tt.args.Gateway.UUID, gateways[0].UUID)
	s.Equal(tt.args.Gateway.CustomerID, gateways[0].CustomerID)
	billingType, err := s.billing.GetBillingTypeByCustomerID(gateways[0].CustomerID)
	s.Require().NoError(err)
	s.Equal(tt.args.BillingType.TOULocationID, billingType.TOULocationID)
	s.Equal(tt.args.BillingType.VoltageType, billingType.VoltageType)
	s.Equal(tt.args.BillingType.TOUType, billingType.TOUType)
}

func (s *BillingWorkerSuite) Test_GetLocalTime() {
	type args struct {
		TOULocationID int
		LocalTime     time.Time
	}

	testTOULocationID := 1
	loc, _ := time.LoadLocation("Asia/Taipei")
	t := time.Now().UTC()
	testLocalTime := t.In(loc)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "GetLocalTime",
			args: args{
				TOULocationID: testTOULocationID,
				LocalTime:     testLocalTime,
			},
		},
		{
			name: "GetLocalTimeInvalidInput",
		},
	}

	for _, tt := range tests {
		switch tt.name {
		case "GetLocalTime":
			localTime, err := s.billing.GetLocalTime(tt.args.TOULocationID, tt.args.LocalTime)
			s.Require().NoError(err)
			s.Equal(testLocalTime, localTime)
		case "GetLocalTimeInvalidInput":
			_, err := s.billing.GetLocalTime(tt.args.TOULocationID, tt.args.LocalTime)
			s.Require().Error(err)
		}
	}
}

func (s *BillingWorkerSuite) Test_getWeeklyBillingParamsByType() {
	type args struct {
		BillingType   *services.BillingType
		LocalTime     time.Time
		BillingParams BillingParams
	}

	testBillingType := &services.BillingType{
		TOULocationID: 1,
		VoltageType:   "Low voltage",
		TOUType:       "Two-section",
	}

	testLocalTime, _ := s.billing.GetLocalTime(testBillingType.TOULocationID, time.Now().UTC())

	var testBillingParams BillingParams
	testBillingParams.Timezone = "+0800"
	timeOnSunday := getSundayOfBillingWeek(testLocalTime, true)
	rate := RateInfo{
		Date:             timeOnSunday.Format(utils.YYYYMMDD),
		Interval:         "0000-2400",
		DemandChargeRate: 47.2,
		TOURate:          1.46,
	}
	if s.billing.IsSummer(testLocalTime) {
		rate.DemandChargeRate = 47.2
		rate.TOURate = 1.46
	} else {
		rate.DemandChargeRate = 34.6
		rate.TOURate = 1.36
	}
	testBillingParams.Rates = append(testBillingParams.Rates, rate)

	tt := struct {
		name string
		args args
	}{
		name: "getWeeklyBillingParamsByType",
		args: args{
			BillingType:   testBillingType,
			LocalTime:     testLocalTime,
			BillingParams: testBillingParams,
		},
	}

	billingParamsJSON, err := getWeeklyBillingParamsByType(s.repo, s.billing, *tt.args.BillingType, tt.args.LocalTime, true)
	var billingParams BillingParams
	err = json.Unmarshal(billingParamsJSON, &billingParams)
	s.Require().NoError(err)
	s.Equal(tt.args.BillingParams.Timezone, billingParams.Timezone)
	s.Equal(tt.args.BillingParams.Rates[0], billingParams.Rates[0])
}

func (s *BillingWorkerSuite) Test_generateBillingParams() {
	testGateways, _ := getGateways(s.repo)
	_, err := generateBillingParams(s.repo, s.billing, testGateways[0], true)
	s.Require().NoError(err)
}
