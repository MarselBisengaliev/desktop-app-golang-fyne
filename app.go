package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Student struct {
	marks []int
	name  string
}

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))

	students := []Student{
		{marks: []int{9, 9, 9}, name: "Kate"},
		{marks: []int{10, 10, 10}, name: "John"},
		{marks: []int{11, 11, 11}, name: "Andrew"},
		{marks: []int{12, 8, 10}, name: "Vlad"},
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(students), len(students[0].marks) + 1
		},
		
		func() fyne.CanvasObject {
			return widget.NewLabel("Default text")
		},
		
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			if tci.Col == 0 {
				co.(*widget.Label).SetText(students[tci.Row].name)
				return
			}

			co.(*widget.Label).SetText(fmt.Sprint(students[tci.Row].marks[tci.Col-1]))
		},
	)

	w.SetContent(table)
	w.Show()
	a.Run()
}
