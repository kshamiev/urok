package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	os.Setenv("FYNE_THEME", "light")

	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.Size{
		Width:  300.0,
		Height: 200.0,
	})
	hello := widget.NewLabel("Hello Fyne 1!")
	w.SetContent(container.NewVBox(
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		hello,
	))
	w.SetMaster()
	w.Show()

	w2 := a.NewWindow("Hello")
	w2.Resize(fyne.Size{
		Width:  600.0,
		Height: 400.0,
	})
	hello2 := widget.NewLabel("Hello Fyne 2!")
	txt := widget.NewEntry()
	w2.SetContent(container.NewVBox(
		widget.NewButton("Hi!", func() {
			hello2.SetText("Welcome :)")
		}),
		hello2,
		txt,
	))
	w2.Show()

	txt.OnChanged = func(s string) {
		hello2.SetText(s)
	}

	fmt.Println("1")
	// w.ShowAndRun()
	a.Run()
	fmt.Println("2")
}
