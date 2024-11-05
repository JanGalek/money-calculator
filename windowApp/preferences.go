package windowApp

import (
	"fyne.io/fyne/v2"
)

func getPreferences() fyne.Preferences {
	app := fyne.CurrentApp()
	return app.Preferences()
}
