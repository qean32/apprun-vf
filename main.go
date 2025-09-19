package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	// "strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

const YBtn = 48

func main() {
	btnSize := [2]float32{120, 120}
	app := app.New()
	window := app.NewWindow("apprun-vf")
	window.Resize(fyne.NewSize(900, 600))
	window.SetFixedSize(true)
	input := TextEntery([2]float32{450, 250}, [2]float32{86, 288})
	runPath := "./run/relax.txt"
	runFn := func(path string) {
		runPath = path
		content, err := os.ReadFile(path)
		if err != nil {
		}
		input.Disable()
		input.SetText(string(content))
		fmt.Println(runPath)
	}
	disableInput := func(path string) {
		input.Enable()
	}

	// ------------------------------------------------- //

	btnSave := Button("Сохранить", "", runFn, "", [2]float32{120, 40}, [2]float32{86, 214})
	btnChange := Button("Изменить", "", disableInput, "", [2]float32{120, 40}, [2]float32{226, 214})
	btnRun := Button("Запустить", "", GlobalAppRun, runPath, [2]float32{120, 40}, [2]float32{689, 499})
	btnGame := Button("", "./public/game#1.svg", runFn, "./run/game#1.txt", btnSize, [2]float32{86, YBtn})
	btnGame_ := Button("", "./public/game#2.svg", runFn, "./run/game#2.txt", btnSize, [2]float32{251, YBtn})
	btnWork := Button("", "./public/work.svg", runFn, "./run/work.txt", btnSize, [2]float32{416, YBtn})
	btnRelax := Button("", "./public/relax.svg", runFn, "./run/relax.txt", btnSize, [2]float32{581, YBtn})
	inputExample := TextEntery([2]float32{245, 180}, [2]float32{569, 288})

	// ------------------------------------------------- //

	content, err := os.ReadFile("./run/example.txt")
	if err != nil {
	}
	inputExample.SetText(string(content))
	input.Disable()
	inputExample.Disable()
	contentPrime, err := os.ReadFile(runPath)
	if err != nil {
	}
	input.SetText(string(contentPrime))

	window.SetContent(container.NewWithoutLayout(
		btnGame,
		btnGame_,
		btnRelax,
		btnWork,
		input,
		inputExample,
		btnSave,
		btnChange,
		btnRun,
	))
	icon, _ := fyne.LoadResourceFromPath("./public/logo-white.svg")
	window.SetIcon(icon)
	window.ShowAndRun()
}

func GlobalAppRun(path string) {
	array, err := Read(path)
	if err != nil {
	}

	for _, path_ := range array {
		Apprun(path_)
	}
}

func Button(placeHolder string, imgPath string, fn func(path string), path string, size [2]float32, position [2]float32) *fyne.Container {
	btn := widget.NewButton(placeHolder, func() { fn(path) })
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

func TextEntery(size [2]float32, position [2]float32) *widget.Entry {
	textInput := widget.NewMultiLineEntry()
	textInput.SetPlaceHolder("Enter text here... n/")
	textInput.Resize(fyne.NewSize(size[0], size[1]))
	textInput.Move(fyne.NewPos(position[0], position[1]))

	return textInput
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

func Apprun(path string) {
	// primePath := strings.Replace(strings.Replace(path, `\`, "/", -1), `"`, `'`, -1)
	cmd := exec.Command(`start "C:/Program Files/BraveSoftware/Brave-Browser/Application/brave.exe"`)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске команды:", err)
		fmt.Println(path)
		return
	}
	fmt.Println("Команда успешно выполнена")
}
