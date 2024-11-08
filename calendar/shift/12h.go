package shift

import (
	"fmt"
	"money-calculator/calendar"
	"time"
)

func GetShiftDayType(datum time.Time, firstMorning time.Time) ShiftDayType {
	dayAmount := int(datum.Sub(firstMorning).Hours() / 24)

	if dayAmount < 0 {
		dayAmount *= -1
	}

	shiftIndex := dayAmount % 8

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
	daysInMonth := calendar.DaysInMonth(year, month)
	shiftDayTypes := []ShiftDayType{}

	fmt.Println(calendar.DaysInMonth(year, month))
	for i := 1; i <= calendar.DaysInMonth(year, month); i++ {
		date := calendar.NewDate(year, month, i)
		shiftDayType := GetShiftDayType(date.DateTime, firstMorning)
		fmt.Println(i)
		shiftDayTypes = append(shiftDayTypes, shiftDayType)

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

	prevs := Calc12HoursWorkShiftAverage(year, month, firstMorning, repeat)

	return &WorkShift{
		Hours:          float64(hours),
		NighHours:      float64(nighHours),
		AfternoonHours: float64(afternoonHours),
		WeekendHours:   float64(weekendHours),
		HolidayHours:   float64(holidayHours),
		DaysInMonth:    daysInMonth,
		DayTypes:       shiftDayTypes,
		Date:           date,
		PrevMonths:     prevs,
	}
}

func Calc12HoursWorkShiftAverage(year int, month int, firstMorning time.Time, repeat bool) []*WorkShift {

	y := year
	if month <= 3 {
		y -= 1
	}

	m1 := calendar.GetDate(y, MonthAverageGroups[month][0], 1)
	m2 := calendar.GetDate(y, MonthAverageGroups[month][1], 1)
	m3 := calendar.GetDate(y, MonthAverageGroups[month][2], 1)

	prevs := []*WorkShift{}
	if repeat {
		i1 := Get12HoursWorkShiftWithRepeat(int(m1.Month()), m1.Year(), firstMorning, false)
		i2 := Get12HoursWorkShiftWithRepeat(int(m2.Month()), m2.Year(), firstMorning, false)
		i3 := Get12HoursWorkShiftWithRepeat(int(m3.Month()), m3.Year(), firstMorning, false)
		prevs = append(prevs, i1, i2, i3)
	}

	return prevs
}
