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
	repo       *repository.Repository
	billing    services.BillingService
	seedUtTime time.Time
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

	// Mock seedUtTime
	loc, _ := time.LoadLocation("Asia/Taipei")
	s.seedUtTime = time.Date(2022, 8, 6, 0, 0, 0, 0, time.UTC).In(loc)
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

	tests := []struct {
		name string
		args args
	}{
		{
			name: "GetLocalTime",
			args: args{
				TOULocationID: testTOULocationID,
				LocalTime:     s.seedUtTime,
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
			s.Equal(s.seedUtTime, localTime)
		case "GetLocalTimeInvalidInput":
			_, err := s.billing.GetLocalTime(tt.args.TOULocationID, tt.args.LocalTime)
			s.Require().Error(err)
		}
	}
}

func (s *BillingWorkerSuite) Test_getSundayOfBillingWeek() {
	type args struct {
		LocalTime    time.Time
		TimeOnSunday time.Time
	}

	loc, _ := time.LoadLocation("Asia/Taipei")

	tests := []struct {
		name string
		args args
	}{
		{
			name: "getSundayOfBillingWeek",
			args: args{
				LocalTime:    s.seedUtTime,
				TimeOnSunday: time.Date(2022, 7, 31, 8, 0, 0, 0, loc),
			},
		},
		{
			name: "getSundayOfBillingWeekNextWeek",
			args: args{
				LocalTime:    s.seedUtTime,
				TimeOnSunday: time.Date(2022, 8, 7, 8, 0, 0, 0, loc),
			},
		},
	}

	for _, tt := range tests {
		switch tt.name {
		case "getSundayOfBillingWeek":
			timeOnSunday := getSundayOfBillingWeek(s.seedUtTime, true)
			s.Equal(tt.args.TimeOnSunday, timeOnSunday)
		case "getSundayOfBillingWeekNextWeek":
			timeOnSunday := getSundayOfBillingWeek(s.seedUtTime, false)
			s.Equal(tt.args.TimeOnSunday, timeOnSunday)
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

	var testBillingParams BillingParams
	testBillingParams.Timezone = "+0800"
	timeOnSunday := getSundayOfBillingWeek(s.seedUtTime, true)
	rate := RateInfo{
		Date:             timeOnSunday.Format(utils.YYYYMMDD),
		Interval:         "0000-2400",
		DemandChargeRate: 47.2,
		TOURate:          1.46,
	}
	if s.billing.IsSummer(s.seedUtTime) {
		rate.DemandChargeRate = 47.2
		rate.TOURate = 1.46
	}
	testBillingParams.Rates = append(testBillingParams.Rates, rate)

	tt := struct {
		name string
		args args
	}{
		name: "getWeeklyBillingParamsByType",
		args: args{
			BillingType:   testBillingType,
			LocalTime:     s.seedUtTime,
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
