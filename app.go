package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Telegram account")
	w.Resize(fyne.NewSize(300, 400))

	url, err := url.Parse("https://t.me/marsel_bisengaliev")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	link := widget.NewHyperlink("Write me!", url)

	w.SetContent(container.NewVBox(
		link,
	))

	w.ShowAndRun()
}
