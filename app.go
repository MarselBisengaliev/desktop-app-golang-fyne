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
	w.Resize(fyne.NewSize(550, 550))

	card1 := widget.NewCard(
		"Marsel Blog",
		"Hello! My name is Marsel, i am 18 years old, from KZ",
		widget.NewButton("Like Marsel", func() {}),
	)

	card2 := widget.NewCard(
		"Title of Kek",
		"desc of kek",
		widget.NewButton("Kek me", func() {}),
	)

	row1 := container.NewVBox(card1, card2)

	w.SetContent(row1)

	w.ShowAndRun()
}
