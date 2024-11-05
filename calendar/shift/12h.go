package shift

import (
	"money-calculator/calendar"
	"time"
)

func get4DaysShifts(firstMorning calendar.Date) {

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

/*
func GetWorkdayTDK(firstMorning calendar.Date, date calendar.Date) {
	days := calendar.DaysInMonth(date.Year, date.Month)

	for i := range days {

	}
}
*/
