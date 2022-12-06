package utils

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// ThreeDecimalPlaces godoc
func ThreeDecimalPlaces(x float32) float32 {
	value, err := strconv.ParseFloat(fmt.Sprintf("%.3f", x), 32)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "three decimal places parse",
			"err":       err,
		}).Error()
		return 0
	}
	return float32(value)
}

// GetZeroForNegativeValue godoc
func GetZeroForNegativeValue(x float32) float32 {
	if x < 0 {
		log.WithFields(log.Fields{
			"caused-by": "value is negative",
			"x":         x,
		}).Warn()
		return 0
	}
	return x
}

// Diff godoc
func Diff(x, y float32) float32 {
	value, err := strconv.ParseFloat(fmt.Sprintf("%.3f", x-y), 32)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "diff parse",
			"err":       err,
		}).Error()
		return 0
	}
	return float32(value)
}

// Division godoc
func Division(x, y float32) float32 {
	if x == 0 || y == 0 {
		log.WithFields(log.Fields{"caused-by": "numerator/denominator is zero"}).Error()
		return 0
	}
	value, err := strconv.ParseFloat(fmt.Sprintf("%.3f", x/y), 32)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "division parse",
			"err":       err,
		}).Error()
		return 0
	}
	return float32(value)
}

// Percent godoc
func Percent(x, y float32) float32 {
	if x == 0 || y == 0 {
		log.WithFields(log.Fields{"caused-by": "numerator/denominator is zero"}).Error()
		return 0
	}
	value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (x/y)*100), 32)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "percent parse",
			"err":       err,
		}).Error()
		return 0
	}
	return float32(value)
}
