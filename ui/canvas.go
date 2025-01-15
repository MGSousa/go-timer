package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"strconv"
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
		X: 110,
		Y: 200,
	})
	c.T.TextSize = 45
	c.T.TextStyle.Bold = true
}

func (c *Canvas) Circle(x, y float32) {
	c.C = canvas.NewCircle(color.Black)
	c.C.Resize(fyne.NewSize(250, 250))
	c.C.Move(fyne.Position{
		X: x,
		Y: y,
	})
}

func (c *Canvas) NewAnimation(duration string) error {
	if c.C == nil {
		return nil
	}
	d, err := strconv.Atoi(duration)
	if err != nil {
		return err
	}

	red := color.NRGBA{R: 0xff, A: 0xff}

	canvas.NewColorRGBAAnimation(
		color.White, red, time.Second*time.Duration(d*60), func(color color.Color) {
			c.C.FillColor = color
			canvas.Refresh(c.C)
		},
	).Start()

	return nil
}
