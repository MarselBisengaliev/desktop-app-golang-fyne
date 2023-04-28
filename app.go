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

	label := widget.NewLabel("")

	sel := widget.NewSelectEntry(
		[]string{
			"Option 1",
			"Option 2",
			"Option 3",
			"Option 4",
			"Option 5",
		},
	)

	sel.PlaceHolder = "Choose one or enter custom option"

	btn := widget.NewButton("Get Data", func() {
		label.SetText("You chose " + sel.Text)
	})

	w.SetContent(
		container.NewVBox(
			sel,
			btn,
			label,
		),
	)
	w.ShowAndRun()
}
