package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"time"
)

type Canvas struct {
	T *canvas.Text
	C *canvas.Circle
}

func NewCanvas() Canvas {
	return Canvas{}
}

func (c *Canvas) Text(title string) {
	c.T = canvas.NewText(title, color.Black)
	c.T.Move(fyne.Position{
		X: 200,
		Y: 200,
	})
	c.T.TextSize = 45
	c.T.TextStyle.Bold = true
}

func (c *Canvas) Circle(x, y float32) {
	c.C = canvas.NewCircle(color.Black)
	c.C.Resize(fyne.NewSize(250, 250))

	// c.circle.Move(fyne.Position{
	//		X: 130,
	//		Y: 110,
	//	})

	c.C.Move(fyne.Position{
		X: x,
		Y: y,
	})
}

func (c *Canvas) NewAnimation() {
	if c.C == nil {
		return
	}
	canvas.NewColorRGBAAnimation(color.White, color.NRGBA{B: 0xff, A: 0xff}, time.Second*100, func(color color.Color) {
		c.C.FillColor = color
		canvas.Refresh(c.C)
	}).Start()
}