package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))

	text := canvas.NewText("Hello world", color.White)
	rec := canvas.NewRectangle(color.White)
	rec.SetMinSize(fyne.NewSize(100, 100))

	cp := dialog.NewColorPicker(
		"Color picker",
		"pick an color",
		func(c color.Color) {
			text.Color = c
			rec.FillColor = c
			rec.Refresh()
		},
		w,
	)

	btn := widget.NewButton("Open color picker", func() {
		cp.Show()
	})

	w.SetContent(
		container.NewVBox(
			btn,
			text,
			rec,
		),
	)
	w.Show()
	a.Run()
}
