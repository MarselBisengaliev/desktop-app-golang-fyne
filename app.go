package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(1200, 600))

	img1 := canvas.NewImageFromFile("icon.png")
	label := widget.NewLabel("Picture")

	w.SetContent(container.NewGridWithColumns(2,  img1, label))

	w.ShowAndRun()
}
