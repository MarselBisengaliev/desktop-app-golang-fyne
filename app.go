package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(250, 250))

	label := widget.NewLabel("Notebook")

	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("Enter here...")

	btn := widget.NewButton("Save", func() {
		file, err := os.Create("info.txt")

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		defer file.Close()

		file.WriteString(entry.Text)
	})

	content := container.NewVBox(label, entry, btn)

	w.SetContent(content)

	w.ShowAndRun()
}
