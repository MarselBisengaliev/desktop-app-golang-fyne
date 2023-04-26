package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createColoredButton(c color.RGBA, b *widget.Button) *fyne.Container {
	button := container.New(
		layout.NewMaxLayout(),

		b,
		canvas.NewRectangle(c),
	)

	return button
}

func createImagedButton(img *canvas.Image, b *widget.Button) *fyne.Container {
	button := container.New(
		layout.NewMaxLayout(),

		b,
		img,
	)

	return button
}

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	w.SetContent(
		container.NewVBox(
			createColoredButton(
				color.RGBA{90, 60, 90, 1},
				widget.NewButton(
					"Click!",
					nil,
				),
			),
			createColoredButton(
				color.RGBA{146, 117, 86, 1},
				widget.NewButton(
					"LOL!",
					nil,
				),
			),
			createImagedButton(
				canvas.NewImageFromFile("icon.png"),
				widget.NewButton(
					"KEK!",
					nil,
				),
			),
		),
	)
	w.ShowAndRun()
}
