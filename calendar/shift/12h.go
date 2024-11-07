package shift

import (
	"money-calculator/calendar"
	"time"
)

type WorkShift struct {
	Hours        float64
	NighHours    float64
	HolidayHours float64
	WeekendHours float64
	Date         *calendar.Date
	PrevMonths   []*WorkShift
}

type ShiftDayType int

const (
	Morning ShiftDayType = 0
	Night   ShiftDayType = 1
	Free    ShiftDayType = 2
)

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
	holidayHours := 0
	weekendHours := 0

	for i := range calendar.DaysInMonth(year, month) {
		date := calendar.NewDate(year, month, i)
		shiftDayType := GetShiftDayType(date.DateTime, firstMorning)

		if shiftDayType != Free {
			hours += 11

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
	m1 := date.DateTime.AddDate(0, -1, 0)
	m2 := date.DateTime.AddDate(0, -2, 0)
	m3 := date.DateTime.AddDate(0, -3, 0)

	prevs := []*WorkShift{}
	if repeat {
		i1 := Get12HoursWorkShiftWithRepeat(int(m1.Month()), m1.Year(), firstMorning, false)
		i2 := Get12HoursWorkShiftWithRepeat(int(m2.Month()), m2.Year(), firstMorning, false)
		i3 := Get12HoursWorkShiftWithRepeat(int(m3.Month()), m3.Year(), firstMorning, false)
		prevs = append(prevs, i1, i2, i3)
	}

	return &WorkShift{
		Hours:        float64(hours),
		NighHours:    float64(nighHours),
		WeekendHours: float64(weekendHours),
		HolidayHours: float64(holidayHours),
		Date:         date,
		PrevMonths:   prevs,
	}
}
