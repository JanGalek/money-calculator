package utils

import (
	"math"
	"time"
)

func GetDate(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func InArray[T comparable](needle T, haystack []T) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func RoundHundreds(number float64) float64 {
	return math.Round(number/100) * 100
}

func RoundTens(number float64) float64 {
	return math.Round(number/10) * 10
}
