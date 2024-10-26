package main

import (
	"fmt"
	"time"
)

func main() {
	year := time.Now().Year()
	locale := "cs"
	hols := getHolidays(year, locale)

	for key, holiday := range hols {
		fmt.Println(key, ":", year, ":", holiday.Format("2006-01-02"))
	}
}
