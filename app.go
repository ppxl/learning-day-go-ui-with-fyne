package main

import (
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Boxed input Clock")
	ui, entry := makeUI()
	clock := widget.NewLabel("")
	content := container.New(layout.NewVBoxLayout(), ui, entry, layout.NewSpacer(), clock)

	updateTime(clock)

	w.SetContent(content)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()

	fmt.Println("Exited")
}

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}
