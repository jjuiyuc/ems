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
		log.WithField("caused-by", "numerator/denominator is zero").Warn()
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
		log.WithField("caused-by", "numerator/denominator is zero").Warn()
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

// SumOfArray godoc
func SumOfArray(array []int) int {
	result := 0
	for _, value := range array {
		result += value
	}
	return result
}

// DiffTwoArrays godoc
func DiffTwoArrays(array1, array2 []int) (diff []int) {
	for i := range array1 {
		if i < len(array2) {
			diff = append(diff, array1[i]-array2[i])
		} else {
			log.WithField("caused-by", "the lengths of two arrays are different").Warn()
			diff = append(diff, 0)
		}
	}
	return
}

// UnorderedEqualTwoArrays godoc
func UnorderedEqualTwoArrays(array1, array2 []int64) bool {
	if len(array1) != len(array2) {
		return false
	}
	exists := make(map[int64]bool)
	for _, value := range array1 {
		exists[value] = true
	}
	for _, value := range array2 {
		if !exists[value] {
			return false
		}
	}
	return true
}
