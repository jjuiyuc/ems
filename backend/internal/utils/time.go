package utils

import (
	"time"
)

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

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func isEndDayOfMonth(t time.Time) bool {
	return t.AddDate(0, 0, 1).Month() != t.Month()
}

// AddDate godoc
func AddDate(t time.Time, years, months, days int) time.Time {
	yearOfTime := t.Year()
	monthOfTime := t.Month()
	dayOfTime := t.Day()

	// 1. Year
	yearOfTime += years

	// 2. Month
	monthOfTime += time.Month(months)
	if monthOfTime > 12 {
		monthOfTime -= 12
		yearOfTime++
	} else if monthOfTime < 1 {
		monthOfTime += 12
		yearOfTime--
	}
	// Adjust day of month
	if (dayOfTime == 28 || dayOfTime == 29 || dayOfTime == 30 || dayOfTime == 31) && monthOfTime == 2 {
		if isLeapYear(yearOfTime) {
			dayOfTime = 29
		} else {
			dayOfTime = 28
		}
	} else if dayOfTime == 31 || isEndDayOfMonth(t) {
		switch monthOfTime {
		case 1, 3, 5, 7, 8, 10, 12:
			dayOfTime = 31
		case 4, 6, 9, 11:
			dayOfTime = 30
		}
	}

	// 3. Day
	dayOfTime += days

	return time.Date(yearOfTime, monthOfTime, dayOfTime, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}
