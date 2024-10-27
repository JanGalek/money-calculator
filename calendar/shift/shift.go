package shift

type Shift struct {
	Name           string
	Hours          int
	PayHours       float32
	WorkInHolidays bool
	Func           func()
}

var Shifts = []Shift{
	{Name: "8h klasik", Hours: 8, PayHours: 7.5, WorkInHolidays: false, Func: GetWorkday},
}

func GetShifts() []Shift {
	return Shifts
}
