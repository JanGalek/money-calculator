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
	"log"
	"strconv"
	"time"
)

type date struct {
	instruction *widget.Label
	dateChosen  *widget.Label
}

func (d *date) onSelected(t time.Time) {
	// use time object to set text on label with given format
	d.instruction.SetText("Date Selected:")
	d.dateChosen.SetText(t.Format("Mon 02 Jan 2006"))
}

func makeGUI(parentWindow fyne.Window) *fyne.Container {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {}),
	)

	result := widget.NewLabel("0")

	entry := widget.NewEntry()
	entry2 := widget.NewEntry()

	// Kontejnery s maximální šířkou
	//entryContainer := container.New(layout.NewHBoxLayout(), entry, layout.NewSpacer())
	//entry2Container := container.New(layout.NewHBoxLayout(), entry2, layout.NewSpacer())

	i := widget.NewLabel("Please Choose a Date")
	i.Alignment = fyne.TextAlignCenter
	l := widget.NewLabel("")
	l.Alignment = fyne.TextAlignCenter
	d := &date{instruction: i, dateChosen: l}
	startingDate := time.Now()

	calendar := xwidget.NewCalendar(startingDate, d.onSelected)

	c := container.NewVBox(i, l, calendar)

	years := widget.NewSelect([]string{"Actual", "2023", "2024", "2025"}, func(value string) {
		log.Println("Select set to", value)
	})

	months := widget.NewSelect([]string{"Actual", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}, func(value string) {
		log.Println("Select set to", value)
	})

	months.SetSelectedIndex(0)
	years.SetSelectedIndex(0)

	form := widget.NewForm(
		widget.NewFormItem("Prvni ranni smena", c),
		widget.NewFormItem("Rok", years),
		widget.NewFormItem("Mesic", months),
		widget.NewFormItem("Hodinovka", entry),
		widget.NewFormItem("Počet hodin", entry2),
	)

	form.OnSubmit = func() {
		hodinovka, err1 := strconv.ParseFloat(entry.Text, 64)
		hodiny, err2 := strconv.ParseFloat(entry2.Text, 64)

		if err1 != nil || err2 != nil {
			dialog.NewInformation("Error", "Prosím, zadejte platná čísla", parentWindow).Show()
			return
		}

		c := hodinovka * hodiny
		result.Text = fmt.Sprintf("%.2f", c)
		result.Refresh()
	}

	// Layouty pro zajištění širších vstupních polí
	cTop := container.New(layout.NewHBoxLayout(), toolbar, layout.NewSpacer())
	content := container.New(layout.NewVBoxLayout(), form) // Změna na VBoxLayout pro vertikální uspořádání
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), result, layout.NewSpacer())

	return container.New(layout.NewVBoxLayout(), cTop, content, centered)
}
