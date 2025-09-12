package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "lib"
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
	itemArray := []fyne.CanvasObject{
		// button("button", "./public/logo.svg", func() {}, [2]float32{500, 500})
	}
	set(itemArray, w.SetContent)

	icon, _ := fyne.LoadResourceFromPath("./public/logo-white.svg")
	w.SetIcon(icon)
	w.ShowAndRun()
}
