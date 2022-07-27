package utils

import "time"

const (
	// ZHHMM ±hhmm
	ZHHMM = "Z0700"
	// YYYY YYYY
	YYYY = "2006"
	// YYYYMMDD YYYY-MM-DD
	YYYYMMDD = "2006-01-02"
	// HHMMSS24h HH:MM:SS
	HHMMSS24h = "15:04:05"
	// HHMM24h HHMM
	HHMM24h = "1504"
)

// IsDateFormat checks the value is a valid time or not
func IsDateFormat(layout, value string) (dateValue time.Time, ok bool) {
	dateValue, err := time.Parse(layout, value)
	if err == nil {
		ok = true
	}
	return
}
