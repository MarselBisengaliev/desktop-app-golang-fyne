package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))

	title := widget.NewLabel("Rate the work of our application from 0 to 10")
	label := widget.NewLabel("Give your feedback about our application here!")
	entry := widget.NewEntry()
	entry.PlaceHolder = "I love..."

	label.Hide()
	entry.Hide()

	slider := widget.NewSlider(0, 10)
	btn := widget.NewButton("Send feedback", func() {
		text := fmt.Sprintf("You're rate our application as	%.0f.\nYour feedback: %s", slider.Value, entry.Text)
		fmt.Println(text)
	})
	btn.Hide()

	slider.OnChanged = func(value float64) {
		label.Show()
		entry.Show()
		btn.Show()
	}


	w.SetContent(
		container.NewVBox(
			title,
			slider,
			label,
			entry,
			btn,
		),
	)
	w.Show()
	a.Run()
}
