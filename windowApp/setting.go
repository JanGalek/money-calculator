package windowApp

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
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
	SetFirstMorning(d.dateTime.String())
}

func getContent(parentWindow fyne.Window) *fyne.Container {

	entry := widget.NewEntry()
	PayPerHour := GetHodinovka()

	entry.Text = fmt.Sprintf("%.2f", PayPerHour)
	entry.Refresh()

	firstMorning := GetFirstMorning()

	i := widget.NewLabel("Please Choose a Date")
	i.Alignment = fyne.TextAlignCenter
	l := widget.NewLabel(firstMorning)
	l.Alignment = fyne.TextAlignCenter
	d := &date{instruction: i, dateChosen: l}
	startingDate := time.Now()

	if firstMorning != "" {
		startingDate, _ = GetTimeFromString(firstMorning)
		d.onSelected(startingDate)
	}

	calendar := xwidget.NewCalendar(startingDate, d.onSelected)
	c := container.NewVBox(i, l, calendar)

	form := widget.NewForm(
		widget.NewFormItem("Prvni ranni smena", c),
		widget.NewFormItem("Hodinovka", entry),
	)

	form.OnSubmit = func() {
		PayPerHour, err1 := strconv.ParseFloat(entry.Text, 64)
		SetHodinovka(PayPerHour)

		if err1 != nil {
			dialog.NewInformation("Error", "Prosím, zadejte platná čísla", parentWindow).Show()
			return
		}
		parentWindow.Close()
	}

	// Layouty pro zajištění širších vstupních polí
	content := container.New(layout.NewVBoxLayout(), form) // Změna na VBoxLayout pro vertikální uspořá

	return container.New(layout.NewVBoxLayout(), content)
}

func GetSettingWindow() fyne.Window {
	window := App().NewWindow("Setting")
	window.SetContent(getContent(window))
	return window
}
