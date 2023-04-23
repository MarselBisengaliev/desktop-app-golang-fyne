package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(300, 500))

	title := widget.NewLabel("Order placement")

	nameLabel := widget.NewLabel("Your name:")
	name:= widget.NewEntry()

	foodLabel := widget.NewLabel("Choice a food for order")
	food := widget.NewCheckGroup([]string{"Pizza", "Cake", "Nuggets", "Burger", "Coca-Cola"}, func([]string) {})

	genderLabel := widget.NewLabel("Your gender:")
	gender := widget.NewRadioGroup([]string{"Male", "Female"}, func(s string) {})

	result := widget.NewLabel("")

	btn := widget.NewButton("Order", func() {
		username := name.Text
		selectedFoods := food.Selected
		selectedGender := gender.Selected

		parsedSelectedFoods := ""
		for i, foodEl := range selectedFoods {
			if i + 1 == len(selectedFoods) {
				parsedSelectedFoods += foodEl
				break
			}
			parsedSelectedFoods += foodEl + ", "
		}

		result.SetText(fmt.Sprintf("%s customer %s ordered: %s.", selectedGender, username, parsedSelectedFoods))
	})

	w.SetContent(container.NewVBox(
		title,
		nameLabel,
		name,
		foodLabel,
		food,
		genderLabel,
		gender,
		btn,
		result,
	))

	w.ShowAndRun()
}
