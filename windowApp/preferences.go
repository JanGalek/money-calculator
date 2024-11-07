package windowApp

import (
	"fyne.io/fyne/v2"
	"time"
)

func App() fyne.App {
	return fyne.CurrentApp()
}

func GetPreferences() fyne.Preferences {
	app := App()
	return app.Preferences()
}

func SetFirstMorning(firstMorning string) {
	GetPreferences().SetString("firstMorning", firstMorning)
}

func GetFirstMorning() string {
	return GetPreferences().StringWithFallback("firstMorning", "")
}

func GetFirstMorningDate() time.Time {
	firstMorning := GetFirstMorning()
	startingDate := time.Now()

	if firstMorning != "" {
		startingDate, _ = GetTimeFromString(firstMorning)
	}

	return startingDate
}

func SetHodinovka(hodinovka float64) {
	GetPreferences().SetFloat("hodinovka", hodinovka)
}

func GetHodinovka() float64 {
	return GetPreferences().FloatWithFallback("hodinovka", 0.0)
}

func SetCoeficientSoc(value float64) {

}

func GetCoeficientSoc() {

}

func SetCoeficientZdrav(value float64) {

}
func GetCoeficientZdrav() {

}

func SetCoeficientTax(value float64) {

}
func GetCoeficientTax() {

}
