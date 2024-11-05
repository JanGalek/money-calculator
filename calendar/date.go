package calendar

import (
	"money-calculator/utils"
	"time"
)

type Date struct {
	Year     int
	Month    int `validate:"min=1,max=12"`
	Day      int `validate:"min=1,max=31"`
	DateTime time.Time
}

func NewDate(year int, month int, day int) Date {
	return Date{Year: year, Month: month, Day: day, DateTime: GetDate(year, month, day)}
}

func (d *Date) isWeekend() bool {
	weekendDays := []time.Weekday{time.Sunday, time.Saturday}
	return utils.InArray(d.DateTime.Weekday(), weekendDays)
}

func (d *Date) time() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
}

// Compare compares the date instant d with u. If d is before u, it returns -1;
// if d is after u, it returns +1; if they're the same, it returns 0.
func (d Date) Compare(u Date) int {
	return d.time().Compare(u.time())
}

func (d Date) Equal(u Date) bool {
	return d.time().Equal(u.time())
}
