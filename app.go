package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel APP")

	label := widget.NewLabel("Label")
	label2 := widget.NewLabel("Label 2")

	w.SetContent(container.NewVBox(
		label,
		label2,
	))
	w.ShowAndRun()
}
