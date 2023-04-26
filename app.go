package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createText() *widget.Label {
	text := widget.NewLabel("Irure occaecat voluptate consectetur exercitation veniam quis cillum laborum. Eiusmod laboris officia qui mollit nulla ipsum. Commodo enim eiusmod Lorem tempor nulla occaecat duis exercitation. Proident enim voluptate exercitation id.")
	text.Wrapping = fyne.TextWrapBreak

	return text
}

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	title := widget.NewLabel("ORTHOGRAPHY")

	rightLabel := widget.NewLabel("RIGHT")
	wrongLabel := widget.NewLabel("WRONG")

	rightText := createText()
	wrongText := createText()

	rightBox := container.NewVBox(rightLabel, rightText)
	wrongBox := container.NewVBox(wrongLabel, wrongText)

	split := container.NewHSplit(rightBox, wrongBox)

	content := container.NewVBox(title, split)

	w.SetContent(content)
	w.ShowAndRun()
}
