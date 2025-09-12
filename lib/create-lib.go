package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func button(placeHolder string, imgPath string, fn func(), size [2]float32) *fyne.Container {
	btn := widget.NewButton(placeHolder, func() {})
	img := canvas.NewImageFromFile(imgPath)
	img.Resize(fyne.NewSize(100, 100))
	container := container.New(
		layout.NewStackLayout(),

		img,
		btn,
	)
	btn.Resize(fyne.NewSize(size[0], size[1]))
	container.Resize(fyne.NewSize(size[0], size[1]))
	fn()

	return container
}
