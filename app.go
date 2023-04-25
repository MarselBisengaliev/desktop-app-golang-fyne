package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createEntry(
	size fyne.Size,
	pos fyne.Position,
	placeholder string,
	isPassword bool,
	isMultiline bool,
) *widget.Entry {
	entry := widget.NewEntry()

	if isPassword {
		entry = widget.NewPasswordEntry()
	} else if isMultiline {
		entry = widget.NewMultiLineEntry()
	}

	entry.Resize(size)
	entry.Move(pos)
	entry.SetPlaceHolder(placeholder)

	return entry
}

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(400, 400))

	label := widget.NewLabel("REGISTRATION")
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Move(fyne.NewPos(140, 3))

	username := createEntry(
		fyne.NewSize(300, 40),
		fyne.NewPos(50, 50),
		"Username",
		false,
		false,
	)

	email := createEntry(
		fyne.NewSize(300, 40),
		fyne.NewPos(50, 110),
		"Email",
		false,
		false,
	)

	password := createEntry(
		fyne.NewSize(300, 40),
		fyne.NewPos(50, 170),
		"Password",
		true,
		false,
	)

	description := createEntry(
		fyne.NewSize(300, 80),
		fyne.NewPos(50, 230),
		"Description",
		false,
		true,
	)

	submit := widget.NewButton("SIGN UP", func() {
		file, err := os.Create("users.txt")

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		defer file.Close()

		data := fmt.Sprintf(
			"-Username: %s\n-Password: %s\n-Email: %s\n-Description: %s\n",
			username.Text,
			password.Text,
			email.Text,
			description.Text,
		)

		file.WriteString(data)
	})
	submit.Resize(fyne.NewSize(200, 40))
	submit.Move(fyne.NewPos(100, 340))

	container := container.NewWithoutLayout(
		label,
		username,
		email,
		password,
		description,
		submit,
	)

	w.SetContent(container)

	w.ShowAndRun()
}
