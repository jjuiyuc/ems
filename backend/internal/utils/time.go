package utils

import "time"

// IsDateFormat checks the value is a valid time or not
func IsDateFormat(layout, value string) (dateValue time.Time, ok bool) {
	dateValue, err := time.Parse(layout, value)
	if err == nil {
		ok = true
	}
	return
}
