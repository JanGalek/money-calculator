package windowApp

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func getCoeficientContent(parentWindow fyne.Window) *fyne.Container {

	entry := widget.NewEntry()
	PayPerHour := GetHodinovka()

	entry.Text = fmt.Sprintf("%.2f", PayPerHour)
	entry.Refresh()

	form := widget.NewForm(
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

func GetCoeficientWindow() fyne.Window {
	window := App().NewWindow("Coeficient setting")
	window.SetContent(getCoeficientContent(window))
	return window
}
