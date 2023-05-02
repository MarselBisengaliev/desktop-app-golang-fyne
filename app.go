package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))

	evaluations := [][]string{
		{"John", "12", "12", "12"},
		{"Andrew", "11", "11", "11"},
		{"Kate", "10", "10", "10"},
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(evaluations), len(evaluations[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Default text")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(evaluations[tci.Row][tci.Col])
		},
	)

	w.SetContent(table)
	w.Show()
	a.Run()
}
