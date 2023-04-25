package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))

	res, err := fyne.LoadResourceFromURLString(
		"https://dce0qyjkutl4h.cloudfront.net/wp-content/uploads/2020/11/Blog_Golang-use-cases.jpg",
	)

	if err != nil {
		fmt.Println(err.Error())
	}

	img := canvas.NewImageFromResource(res)

	w.SetContent(img)

	w.ShowAndRun()
}
