package money

import (
	"math"
	"money-calculator/calendar/shift"
)

type Money struct {
	amount float64
}

type Income struct {
	Hours          float64
	Nights         float64
	Weekends       float64
	Holidays       float64
	GrossSalary    float64
	Basic          float64
	Soc            float64
	Zdrav          float64
	Tax            float64
	Money          float64
	Avarage        float64
	AvaragePerHour float64
}

func Calc(workshift *shift.WorkShift, hodinovka float64) Income {
	avarage := 0.0
	avaragePerHour := hodinovka

	for _, prevWorkshift := range workshift.PrevMonths {
		avarage += Calc(prevWorkshift, hodinovka).GrossSalary
	}

	if avarage > 0 {
		avarage /= float64(len(workshift.PrevMonths))
		avaragePerHour = avarage / 168
	}

	hours := workshift.Hours

	grossSalary := hodinovka * hours
	calc := hodinovka * hours

	nights := avaragePerHour * workshift.NighHours * 0.1
	weekends := avaragePerHour * workshift.WeekendHours * 0.25
	holidays := avaragePerHour * workshift.HolidayHours

	grossSalary += nights + weekends + holidays

	danZaklad := RoundHundreds(grossSalary)

	socialni := math.Round(danZaklad * 0.071)
	zdravotni := math.Round(danZaklad * 0.045)
	danPrijem := math.Round(danZaklad*0.15) - 2570

	cista := danZaklad - socialni
	cista -= zdravotni
	cista -= danPrijem

	return Income{
		Hours:          hours,
		Nights:         nights,
		Weekends:       weekends,
		Holidays:       holidays,
		GrossSalary:    grossSalary,
		Basic:          calc,
		Soc:            socialni,
		Zdrav:          zdravotni,
		Tax:            danPrijem,
		Money:          cista,
		Avarage:        avarage,
		AvaragePerHour: avaragePerHour,
	}
}

func RoundHundreds(number float64) float64 {
	return math.Round(number/100) * 100
}
