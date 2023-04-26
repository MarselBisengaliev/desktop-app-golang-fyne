package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func createRectangle(color color.RGBA) *canvas.Rectangle {
	rectangle := canvas.NewRectangle(color)
	rectangle.SetMinSize(fyne.NewSize(400, 400))

	return rectangle
}

func createImage(file string) *canvas.Image {
	image := canvas.NewImageFromFile(file)
	image.SetMinSize(fyne.NewSize(400, 400))

	return image
}

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	img1 := createImage("icon.png")
	img2 := createImage("icon2.jpg")

	rec1 := createRectangle(color.RGBA{79, 32, 106, 1})
	rec2 := createRectangle(color.RGBA{157, 51, 51, 1})
	rec3 := createRectangle(color.RGBA{118, 183, 39, 1})

	scroll := container.NewVScroll(
		container.NewVBox(
			img1,
			img2,
			rec1,
			rec2,
			rec3,
		),
	)

	w.SetContent(scroll)
	w.ShowAndRun()
}
