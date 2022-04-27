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
	canvas.Circle(50, 110)
	canvas.Text("Timer")

	// init widget
	widget := ui.NewWidget()

	// set widgets
	timeSelected := widget.Select(TIMES)
	timeTxt := widget.Text("Which minutes do you want to set the timer?")
	timeBtn := widget.Button("Confirm", func() {
		if dialog != nil {
			tp, err := time.ParseDuration(fmt.Sprintf("%sm", timeSelected.Selected))
			if err != nil {
				a.Notify(err.Error(), true)
				return
			}

			// add animation for color gradient on circle
			if err = canvas.NewAnimation(timeSelected.Selected); err != nil {
				a.Notify(err.Error(), true)
			}

			go func() {
				i, err := strconv.Atoi(timeSelected.Selected)
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

	// set initial dialog
	wdg := ui.NewContainer(timeTxt, timeSelected, timeBtn)
	dialog = ui.NewDialog("", "Cancel", wdg, window)
	dialog.ResizeAndShow(250, 300)

	window.SetContent(ui.NewContainerWithoutLayout(canvas.C, canvas.T))

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
