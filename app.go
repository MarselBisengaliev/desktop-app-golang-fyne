package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createCenterLayout(widget fyne.Widget) *fyne.Container {
	c := container.New(
		layout.NewCenterLayout(),
		widget,
	)

	return c
}

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")

	l := widget.NewLabel("Label")
	e := widget.NewEntry()
	b := widget.NewButton("Button", func() {
		fmt.Println(e.Text)
	})

	cl := createCenterLayout(l)
	ce := createCenterLayout(e)
	cb := createCenterLayout(b)

	w.SetContent(
		container.NewVBox(
			cl,
			ce,
			cb,
		),
	)
	w.ShowAndRun()
}
