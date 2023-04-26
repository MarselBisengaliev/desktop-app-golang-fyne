package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	item1 := fyne.NewMenuItem("Actions", nil)
	item2 := fyne.NewMenuItem("Say", nil)

	item1.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Print", func() { fmt.Println("Printed") }),
		fyne.NewMenuItem("Save", func() { fmt.Println("Saved") }),
		fyne.NewMenuItem("Cut", func() { fmt.Println("Cuted") }),
		fyne.NewMenuItem("Copy", func() { fmt.Println("Copied") }),
		fyne.NewMenuItem("Open", func() { fmt.Println("Opened") }),
	)

	item2.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Hi", func() { fmt.Println("Hello") }),
		fyne.NewMenuItem("Bye", func() { fmt.Println("Bye bye") }),
		fyne.NewMenuItem("Lol", func() { fmt.Println("Kek") }),
	)

	menu := fyne.NewMenu("Buttons", item1, item2)

	main := fyne.NewMainMenu(menu)

	w.SetMainMenu(main)
	w.ShowAndRun()
}
