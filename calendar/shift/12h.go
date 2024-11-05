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

	return &WorkShift{Hours: float64(hours), NighHours: float64(nighHours), WeekendHours: float64(weekendHours), HolidayHours: float64(holidayHours)}
}
