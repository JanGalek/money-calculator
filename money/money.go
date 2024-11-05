package money

import (
	"money-calculator/calendar/shift"
	"time"
)

type Money struct {
	amount float64
}

type Income struct {
	Money
}

func Calc(hourRate float32, hours float32, shift shift.Shift, time time.Time) float32 {
	return hourRate * hours
}

func BasicCalc(hourRate float32, hours float32) float32 {
	return hourRate * hours
}
