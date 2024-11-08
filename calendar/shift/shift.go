package shift

import "money-calculator/calendar"

type WorkShift struct {
	Hours          float64
	NighHours      float64
	HolidayHours   float64
	WeekendHours   float64
	AfternoonHours float64
	DaysInMonth    int
	DayTypes       []ShiftDayType
	Date           *calendar.Date
	PrevMonths     []*WorkShift
}

type ShiftDayType int

const (
	Morning ShiftDayType = 0
	Night   ShiftDayType = 1
	Free    ShiftDayType = 2
)

type Shift struct {
	Name           string
	Hours          float32
	PayHours       float32
	WorkInHolidays bool
	Func           func()
}

var MonthAverageGroups = [][]int{
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

func NewWorkShift(name string, hours float32, payHours float32, workInHolidays bool, Func func()) *Shift {
	s := &Shift{
		Name:           name,
		Hours:          hours,
		PayHours:       payHours,
		WorkInHolidays: workInHolidays,
		Func:           Func,
	}

	return s
}

var Shifts = []*Shift{
	NewWorkShift("TDK", 12, 11, true, GetWorkday),
	{Name: "8h klasik", Hours: 8, PayHours: 7.5, WorkInHolidays: false, Func: GetWorkday},
}

func GetShifts() []*Shift {
	return Shifts
}
