package main

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))

	btn := widget.NewButton("Click me", func() {

		dialog.ShowCustomConfirm(
			"Say to me true",
			"Absolutely yes",
			"Are you kidding me?",
			widget.NewLabel("Do you love me?"),
			func(isLove bool) {
				if isLove {
					dialog.ShowInformation(
						"I love you too!",
						"You're the best",
						w,
					)
					return
				}

				dialog.ShowError(
					errors.New("you can't say nothing other than YES"),
					w,
				)

				dialog.ShowCustom("Do you understand?", "I got it", widget.NewLabel("You need to say Yes"), w)
			},
			w,
		)
	})

	w.SetContent(btn)
	w.ShowAndRun()
}
