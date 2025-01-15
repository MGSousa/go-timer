package timer

import (
	"fmt"
	"go-timer/ui"

	"strconv"
	"time"

	"fyne.io/fyne/v2"
)

type Countdown struct {
	total  int
	hour   int
	minute int
	second int
}

var (
	dialog *ui.Dialog

	TIMES = []string{
		"1",
		"2",
		"5",
		"10",
		"Other times ...",
	}
	NOTIF = []string{
		"0",
		"10",
		"20",
		"30",
		"40",
		"50",
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
	timeTxt := widget.Text("Which time do you want to set the countdown?")
	timeSelected := widget.Select(TIMES)
	timeSelected.PlaceHolder = "(in minutes)"

	notifyTxt := widget.Text("In which time do you want to be notified?")
	notifySelected := widget.Select(NOTIF)
	notifySelected.PlaceHolder = "(in percentage)"

	timeBtn := widget.Button("Confirm", func() {
		countdownTime := timeSelected.Selected
		notifyTime := notifySelected.Selected

		if dialog != nil {
			tp, err := time.ParseDuration(fmt.Sprintf("%sm", countdownTime))
			if err != nil {
				a.Notify(err.Error(), true)
				return
			}

			// add animation for color gradient on circle
			if err = canvas.NewAnimation(countdownTime); err != nil {
				a.Notify(err.Error(), true)
			}

			i, err := strconv.Atoi(countdownTime)
			if err != nil {
				a.Notify(err.Error(), true)
				return
			}
			ntp, err := strconv.Atoi(notifyTime)
			if err != nil {
				a.Notify("Choose percentage of the time", true)
				return
			}

			go func(i, ntp int) {
				canvas.T.Text = fmt.Sprintf("%02d:59", i-1)
				canvas.T.Refresh()

				diff := time.Now().Add(tp)
				for range time.Tick(1 * time.Second) {
					t := getTimeRemaining(diff)
					canvas.T.Text = fmt.Sprintf("%02d:%02d", t.minute, t.second)
					canvas.T.Refresh()

					if int(float64(t.total)/float64(tp.Seconds())*100) == ntp && ntp != 0 {
						a.Notify(fmt.Sprintf("Remaining %d%%, passed %s", ntp, canvas.T.Text), false)
					}

					if t.total <= 0 {
						a.Notify(fmt.Sprintf("%sm Countdown finished!", countdownTime), false)
						break
					}
				}
			}(i, ntp)

			dialog.Hide()
		}
	})

	// set initial dialog
	wdg := ui.NewContainer(timeTxt, timeSelected, notifyTxt, notifySelected, timeBtn)
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
