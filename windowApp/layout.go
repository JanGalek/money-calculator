package windowApp

import (
	"fyne.io/fyne/v2"
	"log"
)

const sideWith = 2

type AppLayout struct {
	top, left, right, content fyne.CanvasObject
}

func newAppLayout(top, left, right, content fyne.CanvasObject) fyne.Layout {
	return &AppLayout{top: top, left: left, right: right, content: content}
}

func (l *AppLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	height := size.Height - topHeight
	l.top.Resize(fyne.NewSize(size.Width, topHeight))

	l.left.Move(fyne.NewPos(0, topHeight))
	l.left.Resize(fyne.NewSize(sideWith, height))

	l.right.Move(fyne.NewPos(size.Width-sideWith, topHeight))
	l.right.Resize(fyne.NewSize(sideWith, height))

	l.content.Move(fyne.NewPos(sideWith, topHeight))
	l.content.Resize(fyne.NewSize(size.Width-sideWith*2, height))

	log.Println("Size", size)
	//setWindowSize(size.Width, size.Height)

}

func (l *AppLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(50, 50)
}
