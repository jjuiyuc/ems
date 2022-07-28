package utils

import (
	"time"

	log "github.com/sirupsen/logrus"

	"der-ems/repository"
)

// BillingType godoc
type BillingType struct {
	TOULocationID int
	VoltageType   string
	TOUType       string
}

// GetBillingTypeByCustomerID godoc
func GetBillingTypeByCustomerID(repo *repository.Repository, customerID int) (billingType BillingType, err error) {
	customer, err := repo.Customer.GetCustomerByCustomerID(customerID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "repo.Customer.GetCustomerByCustomerID",
			"err":       err,
		}).Error()
		return
	}
	billingType = BillingType{
		TOULocationID: customer.TOULocationID.Int,
		VoltageType:   customer.VoltageType.String,
		TOUType:       customer.TOUType.String,
	}
	return
}

// GetLocalTime godoc
func GetLocalTime(repo *repository.Repository, touLocationID int, t time.Time) (localTime time.Time, err error) {
	touLocation, err := repo.TOU.GetTOULocationByTOULocationID(touLocationID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "repo.TOU.GetTOULocationByTOULocationID",
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
func GetPeriodTypeOfDay(repo *repository.Repository, touLocationID int, t time.Time) (periodType string) {
	// The day is holiday or not
	count, _ := repo.TOU.CountHolidayByDay(touLocationID, t.Format(YYYY), t.Format(YYYYMMDD))

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
func IsSummer(t time.Time) bool {
	// XXX: Hardcode TPC summer is 06/30~09/30
	switch t.Month() {
	case time.June, time.July, time.August, time.September:
		return true
	}
	return false

}
