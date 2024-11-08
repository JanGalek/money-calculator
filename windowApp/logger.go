package windowApp

import "fyne.io/fyne/v2/widget"

var data = []string{"a", "string", "list"}

func getDebugLog() {
	t := widget.NewRichTextWithText("")
	t.AppendMarkdown("")
}
