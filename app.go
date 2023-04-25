package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	fileItem1 := fyne.NewMenuItem("New File", func() { os.Create("file.txt") })
	fileItem2 := fyne.NewMenuItem("Save", func() { fmt.Println("File Saved!") })

	menu1 := fyne.NewMenu("File", fileItem1, fileItem2)

	actionsItem1 := fyne.NewMenuItem("Hello", func() { fmt.Println("Hello from menu") })
	actionsItem2 := fyne.NewMenuItem("Bye", func() { fmt.Println("Bye bye") })
	actionsItem3 := fyne.NewMenuItem("Button", func() { fmt.Println("Clicked") })

	menu2 := fyne.NewMenu("Actions", actionsItem1, actionsItem2, actionsItem3)

	mainMenu := fyne.NewMainMenu(menu1, menu2)
	w.SetMainMenu(mainMenu)

	w.ShowAndRun()
}
