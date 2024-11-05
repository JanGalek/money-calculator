package windowApp

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
	"math"
	"money-calculator/calendar/shift"
	"money-calculator/money"
	"strconv"
	"time"
)

type date struct {
	instruction *widget.Label
	dateChosen  *widget.Label
	dateTime    time.Time
}

func (d *date) onSelected(t time.Time) {
	// use time object to set text on label with given format
	d.instruction.SetText("Date Selected:")
	d.dateChosen.SetText(t.Format("Mon 02 Jan 2006"))
	d.dateTime = t
	getPreferences().SetString("firstMorning", d.dateTime.String())
}

func getTimeFromString(value string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 -0700 MST", value)
}

func makeGUI(parentWindow fyne.Window) *fyne.Container {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {}),
	)

	result := widget.NewLabel("0")
	result2 := widget.NewLabel("0")
	result3 := widget.NewLabel("0")
	result4 := widget.NewLabel("0")
	result5 := widget.NewLabel("0")

	entry := widget.NewEntry()
	hodinovka := getPreferences().Float("hodinovka")

	entry.Text = fmt.Sprintf("%.2f", hodinovka)
	entry.Refresh()

	firstMorning := getPreferences().StringWithFallback("firstMorning", "")

	i := widget.NewLabel("Please Choose a Date")
	i.Alignment = fyne.TextAlignCenter
	l := widget.NewLabel(firstMorning)
	l.Alignment = fyne.TextAlignCenter
	d := &date{instruction: i, dateChosen: l}
	startingDate := time.Now()

	if firstMorning != "" {
		startingDate, _ = getTimeFromString(firstMorning)
		d.onSelected(startingDate)
	}

	calendar := xwidget.NewCalendar(startingDate, d.onSelected)
	c := container.NewVBox(i, l, calendar)

	year := int64(startingDate.Year())
	years := widget.NewSelect([]string{"Actual", "2023", "2024", "2025"}, func(value string) {

		if value != "Actual" {
			year, _ = strconv.ParseInt(value, 10, 64)
		}
	})

	month := int64(time.Now().Month())
	months := widget.NewSelect([]string{"Actual", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}, func(value string) {

		if value != "Actual" {
			month, _ = strconv.ParseInt(value, 10, 64)
		} else {
			month = int64(time.Now().Month())
		}
	})

	months.SetSelectedIndex(0)
	years.SetSelectedIndex(0)

	form := widget.NewForm(
		widget.NewFormItem("Prvni ranni smena", c),
		widget.NewFormItem("Rok", years),
		widget.NewFormItem("Mesic", months),
		widget.NewFormItem("Hodinovka", entry),
	)

	form.OnSubmit = func() {
		hodinovka, err1 := strconv.ParseFloat(entry.Text, 64)
		getPreferences().SetFloat("hodinovka", hodinovka)

		workshift := shift.Get12HoursWorkShift(int(month), int(year), d.dateTime)

		hodiny := workshift.Hours

		if err1 != nil {
			dialog.NewInformation("Error", "Prosím, zadejte platná čísla", parentWindow).Show()
			return
		}

		calc := hodinovka * hodiny
		nights := hodinovka * workshift.NighHours * 0.1
		weekends := hodinovka * workshift.WeekendHours * 0.25
		holidays := hodinovka * workshift.HolidayHours

		zaklad := hodinovka * hodiny

		calc += nights + weekends + holidays

		danZaklad := money.RoundHundreds(calc)

		socialni := math.Round(danZaklad * 0.071)
		zdravotni := math.Round(danZaklad * 0.045)
		danPrijem := math.Round(danZaklad*0.15) - 2570

		cista := danZaklad - socialni
		cista -= zdravotni
		cista -= danPrijem

		result.Text = fmt.Sprintf("Hoding: %.2f, Zaklad: %.2f , Hrubá: %.2f , Čistá: %.2f", hodiny, zaklad, calc, cista)
		result.Refresh()
		result5.Text = fmt.Sprintf("Soc: %.2f, Zdrav: %.2f, Dan: %.2f", socialni, zdravotni, danPrijem)
		result5.Refresh()

		result2.Text = fmt.Sprintf("Noční: %.2f", nights)
		result2.Refresh()
		result3.Text = fmt.Sprintf("Víkendy: %.2f", weekends)
		result3.Refresh()
		result4.Text = fmt.Sprintf("Svátky: %.2f", holidays)
		result4.Refresh()

	}

	// Layouty pro zajištění širších vstupních polí
	cTop := container.New(layout.NewHBoxLayout(), toolbar, layout.NewSpacer())
	content := container.New(layout.NewVBoxLayout(), form) // Změna na VBoxLayout pro vertikální uspořádání
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), result, layout.NewSpacer())
	centered5 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), result5, layout.NewSpacer())
	centered2 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), result2, layout.NewSpacer())
	centered3 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), result3, layout.NewSpacer())
	centered4 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), result4, layout.NewSpacer())

	return container.New(layout.NewVBoxLayout(), cTop, content, centered, centered5, centered2, centered3, centered4)
}
