package main

import (
	"lib"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os/exec"
)

func voidFn() {
	fmt.Println("")
}

func apprun() {
	cmd := exec.Command("")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске команды:", err)
		return
	}
	fmt.Println("Команда успешно выполнена")
}

func main() {
	// writeLines([]string {"zxc", "zxc"}, "table.txt")
	// fmt.Println(readLines("table.txt"))

	// pw()
	// apprun()
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
	itemArray := []fyne.CanvasObject{button("button", "./public/logo.svg", voidFn, [2]float32{500, 500})}
	set(itemArray, w.SetContent)

	icon, _ := fyne.LoadResourceFromPath("./public/logo-white.svg")
	w.SetIcon(icon)
	w.ShowAndRun()
}

func button(placeHolder string, imgPath string, fn func(), size [2]float32) *fyne.Container {
	btn := widget.NewButton(placeHolder, func() { apprun() })
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
