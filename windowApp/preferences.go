package windowApp

import (
	"fyne.io/fyne/v2"
	"log"
)

func getPreferences() fyne.Preferences {
	app := fyne.CurrentApp()
	return app.Preferences()
}

func setWindowSize(width float32, height float32) {
	p := getPreferences()
	log.Println("Change", width, "x", height)
	p.SetFloat("app-width", float64(width))
	p.SetFloat("app-height", float64(height))

	log.Println("Act:", p.Float("app-width"), "x", p.Float("app-height"))
}
