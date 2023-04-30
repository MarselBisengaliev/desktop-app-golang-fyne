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
	w.Resize(fyne.NewSize(500, 500))

	entry := widget.NewMultiLineEntry()
	btn := widget.NewButton("Show the text!", func() {
		label := widget.NewLabel(entry.Text)
		label.Wrapping = fyne.TextWrapBreak

		w2 := a.NewWindow("Window 2")
		w2.Resize(fyne.NewSize(500, 500))

		w2.SetContent(
			container.NewVScroll(
				label,
			),
		)
		w2.Show()
	})

	w.SetContent(
		container.NewVBox(entry, btn),
	)
	w.Show()
	w.SetMaster()
	a.Run()
}
