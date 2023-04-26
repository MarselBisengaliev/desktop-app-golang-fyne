package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	card := widget.NewCard(
		"Marsel GROUP",
		"Read more about the user accept",
		widget.NewAccordion(
			widget.NewAccordionItem(
				"Pargrapgh 1",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Point 1",
						widget.NewLabel("Info about point 1"),
					),
					widget.NewAccordionItem(
						"Point 2",
						widget.NewLabel("Info about point 2"),
					),
					widget.NewAccordionItem(
						"Point 3",
						widget.NewLabel("Info about point 3"),
					),
				),
			),
			widget.NewAccordionItem(
				"Pargrapgh 2",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Point 1",
						widget.NewLabel("Info about point 1"),
					),
					widget.NewAccordionItem(
						"Point 2",
						widget.NewLabel("Info about point 2"),
					),
					widget.NewAccordionItem(
						"Point 3",
						widget.NewLabel("Info about point 3"),
					),
				),
			),
			widget.NewAccordionItem(
				"More information",
				widget.NewLabel("More information for user"),
			),
		),
	)

	w.SetContent(card)
	w.ShowAndRun()
}
