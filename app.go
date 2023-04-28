package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	names := []string{"John", "Kate", "Denis", "Andrew", "Rostik"}

	list := widget.NewList(
		func() int {
			return len(names)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("Create item", func() {
				fmt.Println("Hello")
			})
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Button).SetText(names[id])
		},
	)

	w.SetContent(list)
	w.ShowAndRun()
}
