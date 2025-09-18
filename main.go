package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

const YBtn = 100

func main() {
	btnSize := [2]float32{120, 120}
	app := app.New()
	window := app.NewWindow("apprun-vf")
	window.Resize(fyne.NewSize(900, 600))
	window.SetFixedSize(true)

	// ------------------------------------------------- //
	btnGame := Button("", "./public/game#1.svg", VoidFn, btnSize, [2]float32{100, YBtn})
	btnGame_ := Button("", "./public/game#2.svg", VoidFn, btnSize, [2]float32{300, YBtn})
	btnWork := Button("", "./public/work.svg", VoidFn, btnSize, [2]float32{440, YBtn})
	btnRelax := Button("", "./public/relax.svg", VoidFn, btnSize, [2]float32{660, YBtn})
	text := canvas.NewText("", color.Opaque)
	// ------------------------------------------------- //
	text.Move(fyne.NewPos(200, 300))

	window.SetContent(container.NewWithoutLayout(
		btnGame, btnGame_, btnRelax, btnWork,
		text,
	))
	icon, _ := fyne.LoadResourceFromPath("./public/logo-white.svg")
	window.SetIcon(icon)
	window.ShowAndRun()
}

func Button(placeHolder string, imgPath string, fn func(), size [2]float32, position [2]float32) *fyne.Container {
	btn := widget.NewButton(placeHolder, fn)
	img := canvas.NewImageFromFile(imgPath)
	img.Resize(fyne.NewSize(size[0], size[1]))
	container := container.NewWithoutLayout(
		btn,
		img,
	)
	btn.Resize(fyne.NewSize(size[0], size[1]))
	container.Move(fyne.NewPos(position[0], position[1]))

	return container
}

func VoidBlock() *fyne.Container {
	return container.NewPadded()
}

func Read(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var write []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		write = append(write, scanner.Text())
	}
	return write, scanner.Err()
}

func Write(write []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range write {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func VoidFn() {
	fmt.Println("--------")
}

func Apprun() {
	cmd := exec.Command("")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске команды:", err)
		return
	}
	fmt.Println("Команда успешно выполнена")
}
