package ui

import (
	"fyne.io/fyne/v2/widget"
)

type Widget struct {
}

func NewWidget() *Widget {
	return &Widget{}
}

func (w *Widget) Text(label string) *widget.Label {
	return widget.NewLabel(label)
}

func (w *Widget) Button(label string, tapped func()) *widget.Button {
	return widget.NewButton(label, tapped)
}

func (w *Widget) Select(choices []string) *widget.Select {
	return widget.NewSelect(choices, func(s string) {})
}
