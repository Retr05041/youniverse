package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// Runner for UI
func Run() {
	a := app.New()
	w := a.NewWindow("Main")

	w.SetContent(widget.NewLabel("Hello, World."))
	w.ShowAndRun()
}

