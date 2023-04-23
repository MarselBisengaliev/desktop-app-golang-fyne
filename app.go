package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")

	label1 := widget.NewLabel("Enter a first number: ")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Enter a second number: ")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("")

	btn := widget.NewButton("Calculate", func() {
		n1, err1 := strconv.ParseFloat(entry1.Text, 64)
		n2, err2 := strconv.ParseFloat(entry2.Text, 64)

		if err1 != nil || err2 != nil {
			answer.SetText("Validation error!")
			return
		}

		sum := n1 + n2
		sub := n1 - n2
		mul := n1 * n2
		div := n1 / n2

		answer.SetText(fmt.Sprintf("(+) %f\n(-) %f\n(*) %f\n(/) %f\n", sum, sub, mul, div))
	})

	w.SetContent(container.NewVBox(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		btn,
		answer,
	)))

	w.ShowAndRun()
}
