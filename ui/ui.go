package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type UI struct {
	app fyne.App
}

func NewApp(id string) UI {
	return UI{app: app.NewWithID(id)}
}

func (ui *UI) NewWindow(title string) *Window {
	return &Window{ui.app.NewWindow(title)}
}

func (ui *UI) Notify(content string, error bool) {
	var title = "Info"
	if error {
		title = "Error"
	}
	ui.app.SendNotification(fyne.NewNotification(title, content))
}
