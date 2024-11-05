package windowApp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func GetApp() {
	mApp := app.NewWithID("cz.jangalek.money-calculator")
	p := mApp.Preferences()
	wW := p.FloatWithFallback("app-width", 800)
	wH := p.FloatWithFallback("app-height", 600)

	w := mApp.NewWindow("Money Calculator")
	w.Resize(fyne.Size{Width: float32(wW), Height: float32(wH)})

	w.SetContent(makeGUI(w))
	w.ShowAndRun()
}
