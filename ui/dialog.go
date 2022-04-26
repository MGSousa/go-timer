package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type Dialog struct {
	dialog.Dialog
}

func NewDialog(title, dismiss string, container *fyne.Container, window fyne.Window) *Dialog {
	d := Dialog{dialog.NewCustom(title, dismiss, container, window)}
	d.SetOnClosed(func() {
		fmt.Println("closed")
	})
	return &d
}

func (d *Dialog) ResizeAndShow(w, h float32) {
	d.Resize(fyne.NewSize(w, h))
	d.Show()
}
