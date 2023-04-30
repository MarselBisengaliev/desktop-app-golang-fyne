package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createContainer(text string) *fyne.Container {
	c := container.NewVBox(
		widget.NewLabel(text),
		widget.NewEntry(),
		widget.NewButton("CLICK", nil),
	)

	return c
}

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(500, 500))

	img, err := fyne.LoadResourceFromPath("./icon.png")

	if err != nil {
		fmt.Println("Img has not been found")
		return
	}

	c1 := createContainer("Marsel")
	c2 := createContainer("John")
	c3 := createContainer("Kate")
	c4 := createContainer("Andrew")

	t := container.NewAppTabs(
		container.NewTabItemWithIcon("Tab 1", img, c1),
		container.NewTabItemWithIcon("Tab 2", img, c2),
		container.NewTabItemWithIcon("Tab 3", img, c3),	
		container.NewTabItemWithIcon("Tab 4", img, c4),	
	)

	t.SetTabLocation(container.TabLocationTop)

	w.SetContent(t)
	w.Show()
	a.Run()
}
