package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func startClicks(amounts, delays *widget.Entry, status *widget.Label) error {
	iterations, err := strconv.Atoi(amounts.Text)
	if err != nil {
		return errors.New("Amount must be a number")
	}
	if iterations <= 0 {
		return errors.New("Amount cannot be negative or zero")
	}

	delay, err := strconv.Atoi(delays.Text)
	if err != nil {
		return errors.New("Delay must be a number")
	}
	if iterations <= 0 {
		return errors.New("Delay cannot be negative or zero")
	}

	for i := iterations; i >= 1; i-- {
		status.SetText(fmt.Sprintf("%d clicks left out of %d", i, iterations))
		robotgo.Click()
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	return nil
}

func incrementInEntry(w *widget.Entry) error{
	number, err := strconv.Atoi(w.Text)
	if err != nil {
		return errors.New("Incremental must be a number")
	}
	if number <= 0 {
		return errors.New("Incremental cannot be negative or zero")
	}

	number++
	w.SetText(fmt.Sprint(number))
	return nil
}

func decrementInEntry(w *widget.Entry) error{
	number, err := strconv.Atoi(w.Text)
	if err != nil {
		return errors.New("Decremental must be a number")
	}
	if number <= 0 || number - 1 <= 0 {
		return errors.New("Decremental cannot be negative or zero")
	}

	number--
	w.SetText(fmt.Sprint(number))
	return nil
}

func main() {
	a := app.New()

	w := a.NewWindow("AutoClickerWGA")
	w.Resize(fyne.NewSize(300, 170))
	w.SetFixedSize(true)

	status := widget.NewLabel("")

	h := a.NewWindow("Help")
	h.SetFixedSize(true)
	help := widget.NewLabel("amount - number of repetitions\n" +
		"delay - delay between clicks(ms)\n" +
		"F1 - to show this help\n" +
		"F2 - to start\n" +
		"F3/F5 - to increment amount/delay\n" +
		"F4/F6 - to decrement amount/delay\n" +
		"start - to run with button")
	h.SetContent(help)
	showHelp := widget.NewButton("Help", func() {
		h.Show()
	})
	h.SetCloseIntercept(func() {
		h.Hide()
	})

	amounts := widget.NewEntry()
	amounts.SetText("1")

	delay := widget.NewEntry()
	delay.SetText("250")

	start := widget.NewButton("Start", func() {
		for i := 5; i >= 1; i-- {
			status.SetText(fmt.Sprintf("Clicks will start in %d seconds", i))
			time.Sleep(1 * time.Second)
		}
		err := startClicks(amounts, delay, status)
		if err != nil {
			status.SetText(err.Error())
		} else {
			status.SetText("Clicks ended")
		}
	})

	content := container.New(layout.NewGridLayout(2), widget.NewLabel("Amounts:"), amounts, widget.NewLabel("Delay:"), delay, start, showHelp)

	entryFields := container.NewVBox(content, widget.NewSeparator(), status)

	w.SetContent(entryFields)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyF6:
			err := decrementInEntry(delay)
			if err != nil {
				status.SetText(err.Error())
			} else {
				status.SetText("Delay decremented")
			}
		case fyne.KeyF5:
			err := incrementInEntry(delay)
			if err != nil {
				status.SetText(err.Error())
			} else {
				status.SetText("Delay incremented")
			}
		case fyne.KeyF4:
			err := decrementInEntry(amounts)
			if err != nil {
				status.SetText(err.Error())
			} else {
				status.SetText("Amount decremented")
			}
		case fyne.KeyF3:
			err := incrementInEntry(amounts)
			if err != nil {
				status.SetText(err.Error())
			} else {
				status.SetText("Amount incremented")
			}
		case fyne.KeyF2:
			err := startClicks(amounts, delay, status)
			if err != nil {
				status.SetText(err.Error())
			} else {
				status.SetText("Clicks ended")
			}
			w.RequestFocus()
		case fyne.KeyF1:
			h.Show()
		}
	})

	w.SetCloseIntercept(func (){
		h.Close()
		w.Close()
	})

	w.ShowAndRun()
}
