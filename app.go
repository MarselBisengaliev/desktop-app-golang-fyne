package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))


	pb := widget.NewProgressBarInfinite()
	pb.Hide()

	title := widget.NewLabel("CREATE POST")

	postTitle := widget.NewEntry()
	postTitle.SetPlaceHolder("Your post title")

	postText := widget.NewEntry()
	postText.SetPlaceHolder("Your post text")

	submit := widget.NewButton("Submit", func() {	
		pb.Show()
		time.Sleep(time.Second * 3)
		pb.Hide()

		dialog.ShowInformation("Post creation", "You have created an post", w)
	})

	w.SetContent(
		container.NewVBox(pb, title, postTitle, postText, submit),
	)
	w.Show()
	a.Run()
}
