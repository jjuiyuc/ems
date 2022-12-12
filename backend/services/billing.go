package services

import (
	"time"

	log "github.com/sirupsen/logrus"

	"der-ems/internal/e"
	"der-ems/internal/utils"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// BillingType godoc
type BillingType struct {
	TOULocationID int64
	VoltageType   string
	TOUType       string
}

// BillingService godoc
type BillingService interface {
	GetBillingTypeByCustomerID(customerID int64) (billingType BillingType, err error)
	GetLocalTime(touLocationID int64, t time.Time) (localTime time.Time, err error)
	GetPeriodTypeOfDay(touLocationID int64, t time.Time) (periodType string)
	IsSummer(t time.Time) bool
	GetTOUsOfLocalTime(gwUUID string, t time.Time) (localTime time.Time, tous []*deremsmodels.Tou, err error)
	GetPeakType(localTime time.Time, tous []*deremsmodels.Tou) (peakType string, err error)
}

type defaultBillingService struct {
	repo *repository.Repository
}

// NewBillingService godoc
func NewBillingService(repo *repository.Repository) BillingService {
	return &defaultBillingService{repo}
}

// GetBillingTypeByCustomerID godoc
func (s defaultBillingService) GetBillingTypeByCustomerID(customerID int64) (billingType BillingType, err error) {
	customer, err := s.repo.Customer.GetCustomerByCustomerID(customerID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.Customer.GetCustomerByCustomerID",
			"err":       err,
		}).Error()
		return
	}
	billingType = BillingType{
		TOULocationID: customer.TOULocationID.Int64,
		VoltageType:   customer.VoltageType.String,
		TOUType:       customer.TOUType.String,
	}
	return
}

// GetLocalTime godoc
func (s defaultBillingService) GetLocalTime(touLocationID int64, t time.Time) (localTime time.Time, err error) {
	touLocation, err := s.repo.TOU.GetTOULocationByTOULocationID(touLocationID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.TOU.GetTOULocationByTOULocationID",
			"err":       err,
		}).Error()
		return
	}
	localLocation, err := time.LoadLocation(touLocation.Location.String)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "time.LoadLocation",
			"err":       err,
		}).Error()
		return
	}
	localTime = t.In(localLocation)
	return
}

// GetPeriodTypeOfDay godoc
func (s defaultBillingService) GetPeriodTypeOfDay(touLocationID int64, t time.Time) (periodType string) {
	// The day is holiday or not
	count, _ := s.repo.TOU.CountHolidayByDay(touLocationID, t.Format(utils.YYYY), t.Format(utils.YYYYMMDD))

	switch {
	case count > 0 || t.Weekday() == time.Sunday:
		periodType = "Sunday & Holiday"
	case t.Weekday() == time.Saturday:
		periodType = "Saturday"
	default:
		periodType = "Weekdays"
	}
	return
}

// IsSummer godoc
func (s defaultBillingService) IsSummer(t time.Time) bool {
	// XXX: Hardcode TPC summer is 06/30~09/30
	switch t.Month() {
	case time.June, time.July, time.August, time.September:
		return true
	}
	return false

}

// GetTOUsOfLocalTime godoc
func (s defaultBillingService) GetTOUsOfLocalTime(gwUUID string, t time.Time) (localTime time.Time, tous []*deremsmodels.Tou, err error) {
	gateway, err := s.repo.Gateway.GetGatewayByGatewayUUID(gwUUID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.Gateway.GetGatewayByGatewayUUID",
			"err":       err,
		}).Error()
		return
	}
	billingType, err := s.GetBillingTypeByCustomerID(gateway.CustomerID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.GetBillingTypeByCustomerID",
			"err":       err,
		}).Error()
		return
	}
	localTime, err = s.GetLocalTime(billingType.TOULocationID, t)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.GetLocalTime",
			"err":       err,
		}).Error()
		return
	}
	periodType := s.GetPeriodTypeOfDay(billingType.TOULocationID, localTime)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.GetPeriodTypeOfDay",
			"err":       err,
		}).Error()
		return
	}
	isSummer := s.IsSummer(localTime)
	tous, err = s.repo.TOU.GetTOUsByTOUInfo(billingType.TOULocationID, billingType.VoltageType, billingType.TOUType, periodType, isSummer, localTime.Format(utils.YYYYMMDD))
	if err == nil && len(tous) == 0 {
		err = e.ErrNewBillingsNotExist
	}
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.TOU.GetTOUsByTOUInfo",
			"err":       err,
		}).Error()
	}
	return
}

// GetPeakType godoc
func (s defaultBillingService) GetPeakType(localTime time.Time, tous []*deremsmodels.Tou) (peakType string, err error) {
	loc := time.FixedZone(localTime.Zone())
	localTime, err = time.ParseInLocation(utils.HHMMSS24h, localTime.Format(utils.HHMMSS24h), loc)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "time.ParseInLocation",
			"err":       err,
		}).Error()
		return
	}

	for _, tou := range tous {
		startTime, err := time.ParseInLocation(utils.HHMMSS24h, tou.PeriodStime.String, loc)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "time.ParseInLocation",
				"err":       err,
			}).Error()
			break
		}
		endTime, err := time.ParseInLocation(utils.HHMMSS24h, tou.PeriodEtime.String, loc)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "time.ParseInLocation",
				"err":       err,
			}).Error()
			break
		}
		if tou.PeriodEtime.String == "00:00:00" {
			endTime = endTime.AddDate(0, 0, 1)
		}
		if localTime.After(startTime) && localTime.Before(endTime) {
			peakType = tou.PeakType.String
			break
		}
	}
	return
}
