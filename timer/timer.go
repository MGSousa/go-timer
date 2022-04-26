package timer

import (
	"fmt"
	"fyne.io/fyne/v2"
	"go-timer/ui"
	"strconv"
	"time"
)

type Countdown struct {
	total  int
	hour   int
	minute int
	second int
}

var (
	dialog *ui.Dialog
	TIMES  = []string{
		"1",
		"2",
		"5",
		"10",
	}
)

func New() {
	a := ui.NewApp("go-timer")
	window := a.NewWindow("Timer")

	canvas := ui.NewCanvas()
	canvas.Circle(130, 110)
	canvas.Text("Start")

	// set widgets
	widget := ui.NewWidget()
	widget.Select(TIMES)
	widget.Text("Which minutes do you want to set the timer?")
	widget.Button("Confirm", func() {
		if dialog != nil {
			tp, err := time.ParseDuration(fmt.Sprintf("%sm", widget.Sl.Selected))
			if err != nil {
				a.Notify(err.Error(), true)
				return
			}

			go func() {
				i, err := strconv.Atoi(widget.Sl.Selected)
				if err != nil {
					a.Notify(err.Error(), true)
					return
				}
				canvas.T.Text = fmt.Sprintf("%02d:59", i-1)
				canvas.T.Refresh()

				diff := time.Now().Add(tp)
				for range time.Tick(1 * time.Second) {
					t := getTimeRemaining(diff)
					canvas.T.Text = fmt.Sprintf("%02d:%02d", t.minute, t.second)
					canvas.T.Refresh()

					if t.total <= 0 {
						a.Notify("Countdown over!", false)
						break
					}
				}
			}()

			dialog.Hide()
		}
	})

	wdg := ui.NewContainer(widget.Txt, widget.Sl, widget.Btn)

	// set dialog
	dialog = ui.NewDialog("", "Cancel", wdg, window)
	dialog.ResizeAndShow(500, 200)

	window.SetContent(ui.NewContainerWithoutLayout(canvas.C, canvas.T))

	// add animation for color gradient on circle
	canvas.NewAnimation()

	window.Resize(fyne.NewSize(500, 500))
	window.SetPadded(false)
	window.ShowAndRun()
}

func getTimeRemaining(t time.Time) Countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	hours := total / (60 * 60) % 24
	minutes := int(total/60) % 60
	seconds := total % 60

	return Countdown{
		total:  total,
		hour:   hours,
		minute: minutes,
		second: seconds,
	}
}
