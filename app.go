package main

import (	
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(800, 500))

	entry := widget.NewMultiLineEntry()

	btn := widget.NewButton("Open window", func() {
		dialog.ShowFileSave(func(uc fyne.URIWriteCloser, err error) {
			io.WriteString(uc, entry.Text)
		}, w)
	})

	w.SetContent(
		container.NewVBox(entry, btn),
	)
	w.Show()
	a.Run()
}
