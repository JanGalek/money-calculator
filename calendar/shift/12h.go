package shift

import (
	"money-calculator/calendar"
	"time"
)

type WorkShift struct {
	Hours          float64
	NighHours      float64
	HolidayHours   float64
	WeekendHours   float64
	AfternoonHours float64
	Date           *calendar.Date
	PrevMonths     []*WorkShift
}

type ShiftDayType int

const (
	Morning ShiftDayType = 0
	Night   ShiftDayType = 1
	Free    ShiftDayType = 2
)

var MonthAvageGroups = [][]int{
	1:  {10, 11, 12},
	2:  {10, 11, 12},
	3:  {10, 11, 12},
	4:  {1, 2, 3},
	5:  {1, 2, 3},
	6:  {1, 2, 3},
	7:  {4, 5, 6},
	8:  {4, 5, 6},
	9:  {4, 5, 6},
	10: {7, 8, 9},
	11: {7, 8, 9},
	12: {7, 8, 9},
}

// GetShiftDayType vypočítá směnu pro dané datum na základě zadaného prvního ranního dne
func GetShiftDayType(datum time.Time, firstMorning time.Time) ShiftDayType {
	// Vypočítáme počet dní od prvního ranního dne do zadaného data
	dayAmount := int(datum.Sub(firstMorning).Hours() / 24)

	if dayAmount < 0 {
		dayAmount *= -1
	}

	// Vypočítáme index směny v cyklu (0-7)
	shiftIndex := dayAmount % 8

	// Podle indexu směny vrátíme odpovídající typ směny
	switch shiftIndex {
	case 0, 1:
		return Morning
	case 2, 3:
		return Night
	default:
		return Free
	}
}
func Get12HoursWorkShift(month int, year int, firstMorning time.Time) *WorkShift {
	return Get12HoursWorkShiftWithRepeat(month, year, firstMorning, true)
}

func Get12HoursWorkShiftWithRepeat(month int, year int, firstMorning time.Time, repeat bool) *WorkShift {
	hours := 0
	nighHours := 0
	afternoonHours := 0
	holidayHours := 0
	weekendHours := 0

	for i := range calendar.DaysInMonth(year, month) {
		date := calendar.NewDate(year, month, i)
		shiftDayType := GetShiftDayType(date.DateTime, firstMorning)

		if shiftDayType != Free {
			hours += 11
			afternoonHours += 4

			if shiftDayType == Night {
				nighHours += 8
			}

			if date.IsWeekend() {
				weekendHours += 11
			}

			if date.IsHoliday() {
				holidayHours += 11
			}
		}
	}
	date := calendar.NewDate(year, month, 1)

	y := year
	if month <= 3 {
		y -= 1
	}

	m1 := calendar.GetDate(y, MonthAvageGroups[month][0], 1)
	m2 := calendar.GetDate(y, MonthAvageGroups[month][1], 1)
	m3 := calendar.GetDate(y, MonthAvageGroups[month][2], 1)

	prevs := []*WorkShift{}
	if repeat {
		i1 := Get12HoursWorkShiftWithRepeat(int(m1.Month()), m1.Year(), firstMorning, false)
		i2 := Get12HoursWorkShiftWithRepeat(int(m2.Month()), m2.Year(), firstMorning, false)
		i3 := Get12HoursWorkShiftWithRepeat(int(m3.Month()), m3.Year(), firstMorning, false)
		prevs = append(prevs, i1, i2, i3)
	}

	return &WorkShift{
		Hours:          float64(hours),
		NighHours:      float64(nighHours),
		AfternoonHours: float64(afternoonHours),
		WeekendHours:   float64(weekendHours),
		HolidayHours:   float64(holidayHours),
		Date:           date,
		PrevMonths:     prevs,
	}
}
