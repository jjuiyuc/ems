package services

import (
	"time"

	log "github.com/sirupsen/logrus"

	"der-ems/internal/utils"
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
