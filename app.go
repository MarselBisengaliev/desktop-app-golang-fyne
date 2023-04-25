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
	w.Resize(fyne.NewSize(500, 500))

	ic, err := fyne.LoadResourceFromPath("icon.png")

	if err != nil {
		fmt.Println(err.Error())	
		return
	}

	w.SetIcon(ic)
	a.SetIcon(ic)
	
	w.SetContent(widget.NewLabel("Marsel App"))

	w.ShowAndRun()
}
