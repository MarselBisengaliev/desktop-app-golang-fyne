package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	label := widget.NewLabel("Some text here...")
	btn := widget.NewButton("Change visibility", func() {
		label.Hidden = !label.Hidden
	})

	w.SetContent(container.NewVBox(btn, label))
	w.ShowAndRun()
}
