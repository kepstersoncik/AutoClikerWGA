package main

import (
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func startClicks(amounts, delays *widget.Entry) {
	iterations, err := strconv.Atoi(amounts.Text)
	if err != nil {
		return
	}
	delay, err := strconv.Atoi(delays.Text)
	if err != nil {
		return
	}

	for i := 0; i < iterations; i++ {
		robotgo.Click()
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func main() {
	a := app.New()

	w := a.NewWindow("AutoClickerWGA")
	w.Resize(fyne.NewSize(300, 140))
	w.SetFixedSize(true)

	h := a.NewWindow("Help")
	h.SetFixedSize(true)
	help := widget.NewLabel("amount - number of repetitions\n" +
		"delay - delay between clicks(ms)\n" +
		"F1 - to show this help\n" +
		"F2 - to start\n" +
		"start - to run with button")
	h.SetContent(help)
	showHelp := widget.NewButton("Help", func() { h.Show() })

	amounts := widget.NewEntry()
	amounts.SetText("1")

	delay := widget.NewEntry()
	delay.SetText("250")

	start := widget.NewButton("Start", func() { startClicks(amounts, delay) })

	content := container.New(layout.NewGridLayout(2), widget.NewLabel("Amounts:"), amounts, widget.NewLabel("Delay:"), delay, start, showHelp)

	entryFields := container.NewVBox(content)

	w.SetContent(entryFields)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyF2:
			startClicks(amounts, delay)
			w.RequestFocus()
		case fyne.KeyF1:
			h.Show()
		}
	})

	w.ShowAndRun()
}
