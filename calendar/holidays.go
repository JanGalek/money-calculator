package calendar

import (
	"sort"
	"time"
)

type Holiday struct {
	Date  string
	Day   int
	Month int
	Name  string
}

var listHolidays = map[string][]Holiday{
	"cs": {
		{Date: "01-01", Day: 1, Month: 1, Name: "New Year"},
		{Date: "01-05", Day: 1, Month: 5, Name: "New Year"},
		{Date: "08-05", Day: 8, Month: 5, Name: "New Year"},
		{Date: "05-07", Day: 5, Month: 7, Name: "New Year"},
		{Date: "06-07", Day: 6, Month: 7, Name: "New Year"},
		{Date: "28-10", Day: 28, Month: 10, Name: "New Year"},
		{Date: "17-11", Day: 17, Month: 11, Name: "New Year"},
		{Date: "24-12", Day: 24, Month: 12, Name: "New Year"},
		{Date: "25-12", Day: 25, Month: 12, Name: "New Year"},
		{Date: "26-12", Day: 26, Month: 12, Name: "New Year"},
	},
}

func GetHolidays(year int, locale string) []time.Time {
	yearHolidays := []time.Time{
		GetGoodFriday(year),
		GetMonday(year),
	}

	for _, holiday := range listHolidays[locale] {
		date := GetDate(year, holiday.Month, holiday.Day)
		yearHolidays = append(yearHolidays, date)
	}

	sort.Slice(yearHolidays, func(i, j int) bool {
		return yearHolidays[i].Format("2006-01-02") < yearHolidays[j].Format("2006-01-02")
	})

	return yearHolidays
}
