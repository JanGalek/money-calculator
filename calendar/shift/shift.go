package shift

type Shift struct {
	Name           string
	Hours          float32
	PayHours       float32
	WorkInHolidays bool
	Func           func()
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
