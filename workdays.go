package main

import "time"

type RepeatType struct {
}

type Shift struct {
	Name       string
	RepeatType RepeatType
}

// 12 hours, 2 days mornings, 2 days nights and 4 days free
func ShiftForMonth_12h2m2n4f(firstMorning time.Time, year int, month int) ([]time.Time) {
	days := []time.Time

	for k, v := range ListDaysInMonth(getDate(year, month, 1)) {
		println()
	}
}

func ListDaysInMonth(t time.Time) []int {
	days := make([]int, DaysInMonth(t))
	for i := range days {
		days[i] = i + 1
	}
	return days
}

func DaysInMonth(t time.Time) int {
	y, m, _ := t.Date()
	return time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}