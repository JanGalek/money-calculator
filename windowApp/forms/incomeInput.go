package forms

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type IncomeInputWidget struct {
	widget.BaseWidget
	Title   *widget.Label
	Comment *widget.Label
}

func NewIncomeInputWidget(title, comment string) *IncomeInputWidget {
	item := &IncomeInputWidget{
		Title:   widget.NewLabel(title),
		Comment: widget.NewLabel(comment),
	}
	item.Title.Truncation = fyne.TextTruncateEllipsis
	item.Theme().
		item.ExtendBaseWidget(item)

	return item
}

func (item *IncomeInputWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, nil, item.Comment, item.Title)
	return widget.NewSimpleRenderer(c)
}
