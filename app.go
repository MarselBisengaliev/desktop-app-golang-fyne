package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))

	btn1 := widget.NewButton("Set light theme", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	btn2 := widget.NewButton("Set dark theme", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})

	btn_box := container.NewHBox(btn1, btn2)	
	content := container.NewVBox(btn_box)

	w.SetContent(content)

	w.ShowAndRun()
}
