package money

import "C"
import (
	"math"
	"money-calculator/calendar/shift"
	"money-calculator/utils"
)

type Money struct {
	amount float64
}

func Calc(workshift *shift.WorkShift, hodinovka float64) Income {
	avarage := 0.0
	avarageHours := 0.0
	avaragePerHour := hodinovka

	for _, prevWorkshift := range workshift.PrevMonths {
		p := Calc(prevWorkshift, hodinovka)
		avarage += p.GrossSalary
		avarageHours += p.Hours
	}

	if avarage > 0 {
		avarage /= float64(len(workshift.PrevMonths))
		avarageHours /= float64(len(workshift.PrevMonths))
		avaragePerHour = avarage / avarageHours
	}

	hours := workshift.Hours

	grossSalary := hodinovka * hours
	calc := hodinovka * hours

	nights := avaragePerHour * workshift.NighHours * 0.1
	afternoons := workshift.AfternoonHours * 9
	weekends := avaragePerHour * workshift.WeekendHours * 0.1
	holidays := avaragePerHour * workshift.HolidayHours

	grossSalary += nights + afternoons + weekends + holidays

	danZaklad := math.RoundToEven(grossSalary)

	socialni := math.Round(danZaklad * 0.071)
	zdravotni := math.Round(danZaklad * 0.045)
	danPrijem := utils.RoundTens(danZaklad*0.15) - 2570

	cista := danZaklad - socialni
	cista -= zdravotni
	cista -= danPrijem

	return Income{
		Hours:          hours,
		Nights:         nights,
		Afternoons:     afternoons,
		Weekends:       weekends,
		Holidays:       holidays,
		GrossSalary:    danZaklad,
		Basic:          calc,
		Soc:            socialni,
		Zdrav:          zdravotni,
		Tax:            danPrijem,
		Money:          cista,
		Avarage:        avarage,
		AvaragePerHour: avaragePerHour,
	}
}
