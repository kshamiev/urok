package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	content := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		newWidget(20)...,
	)
	myWindow := myApp.NewWindow("GridWrapLayout")
	myWindow.Resize(fyne.NewSize(600.0, 400.0))
	myWindow.CenterOnScreen()
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func newWidget(count int) []fyne.CanvasObject {
	res := make([]fyne.CanvasObject, count)
	for i := 0; i < count; i++ {
		res[i] = widget.NewEntry()
	}
	return res
}
