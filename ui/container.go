package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// NewContainer initiate new VBOX container
func NewContainer(obj ...fyne.CanvasObject) *fyne.Container {
	return container.NewVBox(obj...)
}

// NewContainerWithoutLayout initiate new container without defined layout
func NewContainerWithoutLayout(obj ...fyne.CanvasObject) *fyne.Container {
	return container.NewWithoutLayout(obj...)
}
