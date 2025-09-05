package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("apprun-vf")
	w.Resize(fyne.NewSize(1200, 800))
	w.SetFixedSize(true)

	// ____________________________________ place to create ____________________________________ //

	// ____________________________________ place to create ____________________________________ //

	set := func(array []fyne.CanvasObject, func_ func(fyne.CanvasObject)) {
		for i := 0; i < len(array); i++ {
			func_(array[i])
		}
	}
	itemArray := []fyne.CanvasObject{red_button()}
	set(itemArray, w.SetContent)

	icon, _ := fyne.LoadResourceFromPath("./public/logo-white.svg")
	w.SetIcon(icon)
	w.ShowAndRun()
}

func red_button() *fyne.Container {
	btn := widget.NewButton("Visit", nil)

	btn_color := canvas.NewRectangle(
		color.NRGBA{R: 255, G: 0, B: 0, A: 255})

	container := container.New(
		layout.NewCenterLayout(),
		btn_color,
		btn,
	)

	return container
}
