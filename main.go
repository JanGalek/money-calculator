package main

import (
	"fmt"
	"money-calculator/calendar"
	"money-calculator/calendar/shift"
	"money-calculator/windowApp"
	"time"
)

func main() {
	year := time.Now().Year()
	locale := "cs"
	hols := calendar.GetHolidays(year, locale)

	shift.GetShifts()

	fmt.Println(calendar.DaysInMonth(year, 2))
	calendar.GetCalendar()
	for key, holiday := range hols {
		fmt.Println(key, ":", year, ":", holiday.Format("2006-01-02"))
	}

	windowApp.GetApp()
}
