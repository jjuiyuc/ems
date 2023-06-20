package apps

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
	"der-ems/services"
	"der-ems/testutils"
	"der-ems/testutils/testdata"
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

	// Truncate & seed data
	err := testutils.SeedUtLocationAndGateway(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())

	// Mock seedUtTime
	loc, _ := time.LoadLocation("Asia/Taipei")
	s.seedUtTime = time.Date(2022, 8, 6, 0, 0, 0, 0, time.UTC).In(loc)
}

func (s *BillingWorkerSuite) TearDownSuite() {
	models.Close()
}

func (s *BillingWorkerSuite) Test_GetBillingTypeByLocationID() {
	type response struct {
		Gateway     *deremsmodels.Gateway
		BillingType *services.BillingType
	}

	testGateway := &deremsmodels.Gateway{
		ID:         testdata.UtGateway.ID,
		UUID:       testdata.UtGateway.UUID,
		LocationID: testdata.UtGateway.LocationID,
	}
	testBillingType := &services.BillingType{
		TOULocationID: testdata.UtLocation.TOULocationID.Int64,
		VoltageType:   testdata.UtLocation.VoltageType.String,
		TOUType:       testdata.UtLocation.TOUType.String,
	}

	tt := struct {
		name   string
		wantRv response
	}{
		name: "GetBillingTypeByLocationID",
		wantRv: response{
			Gateway:     testGateway,
			BillingType: testBillingType,
		},
	}

	gateways, err := getGateways(s.repo)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.Equalf(tt.wantRv.Gateway.ID, gateways[0].ID, e.ErrNewMessageNotEqual.Error())
	s.Equalf(tt.wantRv.Gateway.UUID, gateways[0].UUID, e.ErrNewMessageNotEqual.Error())
	s.Equalf(tt.wantRv.Gateway.LocationID, gateways[0].LocationID, e.ErrNewMessageNotEqual.Error())
	billingType, err := s.billing.GetBillingTypeByLocationID(gateways[0].LocationID.Int64)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.Equalf(tt.wantRv.BillingType.TOULocationID, billingType.TOULocationID, e.ErrNewMessageNotEqual.Error())
	s.Equalf(tt.wantRv.BillingType.VoltageType, billingType.VoltageType, e.ErrNewMessageNotEqual.Error())
	s.Equalf(tt.wantRv.BillingType.TOUType, billingType.TOUType, e.ErrNewMessageNotEqual.Error())
}

func (s *BillingWorkerSuite) Test_GetLocalTime() {
	type args struct {
		TOULocationID int64
		LocalTime     time.Time
	}

	seedUtTOULocationID := int64(1)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "GetLocalTime",
			args: args{
				TOULocationID: seedUtTOULocationID,
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
			s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
			s.Equalf(s.seedUtTime, localTime, e.ErrNewMessageNotEqual.Error())
		case "GetLocalTimeInvalidInput":
			_, err := s.billing.GetLocalTime(tt.args.TOULocationID, tt.args.LocalTime)
			s.Require().Errorf(err, e.ErrNewMessageGotNil.Error())
		}
	}
}

func (s *BillingWorkerSuite) Test_GetSundayOfBillingWeek() {
	type args struct {
		LocalTime time.Time
		SendNow   bool
	}

	type response struct {
		TimeOnSunday time.Time
	}

	loc, _ := time.LoadLocation("Asia/Taipei")

	tests := []struct {
		name   string
		args   args
		wantRv response
	}{
		{
			name: "GetSundayOfBillingWeek",
			args: args{
				LocalTime: s.seedUtTime,
				SendNow:   true,
			},
			wantRv: response{
				TimeOnSunday: time.Date(2022, 7, 31, 8, 0, 0, 0, loc),
			},
		},
		{
			name: "GetSundayOfBillingWeekNextWeek",
			args: args{
				LocalTime: s.seedUtTime,
				SendNow:   false,
			},
			wantRv: response{
				TimeOnSunday: time.Date(2022, 8, 7, 8, 0, 0, 0, loc),
			},
		},
	}

	for _, tt := range tests {
		timeOnSunday := s.billing.GetSundayOfBillingWeek(tt.args.LocalTime, tt.args.SendNow)
		s.Equalf(tt.wantRv.TimeOnSunday, timeOnSunday, e.ErrNewMessageNotEqual.Error())
	}
}

func (s *BillingWorkerSuite) Test_GetWeeklyBillingParamsByType() {
	type args struct {
		BillingType *services.BillingType
		LocalTime   time.Time
		SendNow     bool
	}

	type response struct {
		BillingParams services.BillingParams
	}

	seedUtBillingType := &services.BillingType{
		TOULocationID: testdata.UtLocation.TOULocationID.Int64,
		VoltageType:   testdata.UtLocation.VoltageType.String,
		TOUType:       testdata.UtLocation.TOUType.String,
	}

	var testBillingParams services.BillingParams
	testBillingParams.Timezone = "+0800"
	timeOnSunday := s.billing.GetSundayOfBillingWeek(s.seedUtTime, true)
	rate := services.RateInfo{
		Date:             timeOnSunday.Format(utils.YYYYMMDD),
		Interval:         "0000-2400",
		DemandChargeRate: 47.2,
		TOURate:          1.46,
	}
	if s.billing.IsSummer(seedUtBillingType.VoltageType, s.seedUtTime) {
		rate.DemandChargeRate = 47.2
		rate.TOURate = 1.46
	}
	testBillingParams.Rates = append(testBillingParams.Rates, rate)

	tt := struct {
		name   string
		args   args
		wantRv response
	}{
		name: "GetWeeklyBillingParamsByType",
		args: args{
			BillingType: seedUtBillingType,
			LocalTime:   s.seedUtTime,
			SendNow:     true,
		},
		wantRv: response{
			BillingParams: testBillingParams,
		},
	}

	billingParamsJSON, err := s.billing.GetWeeklyBillingParamsByType(*tt.args.BillingType, tt.args.LocalTime, tt.args.SendNow)
	var billingParams services.BillingParams
	err = json.Unmarshal(billingParamsJSON, &billingParams)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.Equalf(tt.wantRv.BillingParams.Timezone, billingParams.Timezone, e.ErrNewMessageNotEqual.Error())
	s.Equalf(tt.wantRv.BillingParams.Rates[0], billingParams.Rates[0], e.ErrNewMessageNotEqual.Error())
}

func (s *BillingWorkerSuite) Test_generateBillingParams() {
	gateways, _ := getGateways(s.repo)
	_, err := s.billing.GenerateBillingParams(gateways[0], true)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
}
