package ui

import (
	"fyne.io/fyne/v2/widget"
)

type Widget struct {
	Txt *widget.Label
	Btn *widget.Button
	Sl  *widget.Select
}

func NewWidget() *Widget {
	return &Widget{}
}

func (w *Widget) Text(label string) {
	w.Txt = widget.NewLabel(label)
}

func (w *Widget) Button(label string, tapped func()) {
	w.Btn = widget.NewButton(label, tapped)
}

func (w *Widget) Select(choices []string) {
	w.Sl = widget.NewSelect(choices, func(s string) {})
}
