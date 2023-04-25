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

	btn1 := widget.NewButton("Button 1", func() {})
	btn2 := widget.NewButton("Button 2", func() {})
	btn3 := widget.NewButton("Button 3", func() {})
	btn4 := widget.NewButton("Button 4", func() {})

	label1 := widget.NewLabel("Label 1")
	label2 := widget.NewLabel("Label 2")
	label3 := widget.NewLabel("Label 3")
	label4 := widget.NewLabel("Label 4")

	btn_box := container.NewHBox(btn1, btn2, btn3, btn4)	
	label_box := container.NewHBox(label1, label2, label3, label4)
	content := container.NewVBox(label_box, btn_box)

	w.SetContent(content)

	w.ShowAndRun()
}
