package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(300, 300))

	color := color.NRGBA{R: 57, G: 8, B: 137, A: 255}
	label := canvas.NewText("Color", color)

	w.SetContent(label)

	w.ShowAndRun()
}
