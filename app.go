package main

import (
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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
	w := a.NewWindow("Shopping List")

	clock := widget.NewLabel("")

	baseContainer := container.New(layout.NewCenterLayout())
	createItemButton := widget.NewButtonWithIcon("Add Item", theme.ContentAddIcon(), func() {
		baseContainer.RemoveAll()
		baseContainer.Add(widget.NewLabel("oh hello"))
		backButton := widget.NewButtonWithIcon("return", theme.NavigateBackIcon(), func() {
			// TODO current problem: I have no idea how to fix the problem of cyclical functional dependencies
			// this code adds the clock but no longer is the same as the previous call because I need to get hold of the
			// clock container containing this button ... ad infinitem lol
			baseContainer.RemoveAll()
			baseContainer.Add(clock)
		})
		baseContainer.Add(backButton)
	})

	clockContent := container.New(layout.NewVBoxLayout(), createItemButton, layout.NewSpacer(), clock)
	baseContainer.Add(clockContent)

	updateTime(clock)

	w.SetContent(baseContainer)

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
