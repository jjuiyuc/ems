package utils

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Diff godoc
func Diff(x, y float32) float32 {
	value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", x-y), 32)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "diff parse",
			"err":       err,
		}).Error()
		return 0
	}
	return float32(value)
}

// Percent godoc
func Percent(x, y float32) float32 {
	if y == 0 {
		log.WithFields(log.Fields{"caused-by": "denominator is zero"}).Error()
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
