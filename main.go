package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	// "strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

const YBtn = 48

var currentPath = "./run/relax.txt"

func main() {
	btnSize := [2]float32{120, 120}
	app := app.New()
	window := app.NewWindow("apprun-vf")
	window.Resize(fyne.NewSize(900, 600))
	window.SetFixedSize(true)

	// ------------------------------------------------- //

	input := textEntery([2]float32{450, 250}, [2]float32{86, 288})
	inputExample := textEntery([2]float32{245, 180}, [2]float32{569, 288})

	// ------------------------------------------------- //

	selectFunction := func(path string) {
		currentPath = path
		content, err := os.ReadFile(path)
		if err != nil {
		}
		input.Disable()
		input.SetText(string(content))
		fmt.Println(currentPath)
	}
	disableInput := func() {
		input.Enable()
	}

	// ------------------------------------------------- //

	btnSave := widget.NewButton("Сохранить", func() { write(input.Text) })
	btnSave.Resize(fyne.NewSize(120, 40))
	btnSave.Move(fyne.NewPos(86, 214))

	btnChange := widget.NewButton("Изменить", func() { disableInput() })
	btnChange.Resize(fyne.NewSize(120, 40))
	btnChange.Move(fyne.NewPos(226, 214))

	btnRun := widget.NewButton("Запустить", func() { globalAppRun(window) })
	btnRun.Resize(fyne.NewSize(120, 40))
	btnRun.Move(fyne.NewPos(689, 499))

	// ------------------------------------------------- //

	btnGame := button("", "./public/game#1.svg", selectFunction, "./run/game#1.txt", btnSize, [2]float32{86, YBtn})
	btnGame_ := button("", "./public/game#2.svg", selectFunction, "./run/game#2.txt", btnSize, [2]float32{251, YBtn})
	btnWork := button("", "./public/work.svg", selectFunction, "./run/work.txt", btnSize, [2]float32{416, YBtn})
	btnRelax := button("", "./public/relax.svg", selectFunction, "./run/relax.txt", btnSize, [2]float32{581, YBtn})

	// ------------------------------------------------- //

	content, err := os.ReadFile("./run/example.txt")
	if err != nil {
	}
	inputExample.SetText(string(content))
	input.Disable()
	inputExample.Disable()
	contentPrime, err := os.ReadFile(currentPath)
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

func button(placeHolder string, imgPath string, fn func(path string), path string, size [2]float32, position [2]float32) *fyne.Container {
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

func textEntery(size [2]float32, position [2]float32) *widget.Entry {
	textInput := widget.NewMultiLineEntry()
	textInput.SetPlaceHolder("Enter text here... n/")
	textInput.Resize(fyne.NewSize(size[0], size[1]))
	textInput.Move(fyne.NewPos(position[0], position[1]))

	return textInput
}

func read(path string) ([]string, error) {
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

func write(write string) {
	fmt.Println(currentPath)
	fmt.Println(write)
	file, err := os.Create(currentPath)

	if err != nil {
		fmt.Println("content")
	}

	defer file.Close()
	file.WriteString(write)
}

func globalAppRun(window fyne.Window) {
	array, err := read(currentPath)
	if err != nil {
	}

	for _, path_ := range array {
		fmt.Println("run " + path_)
		appRun(path_)
	}
	fmt.Println("--------------------------")
	window.Close()
}

func appRun(path string) {
	extension := strings.Split(path, `.`)
	if extension[len(extension)-1] == "bat" {
		cmd := exec.Command("CMD.exe", "/C", path)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Ошибка при запуске команды:", err)
			return
		}
	} else {
		cmd := exec.Command(path)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Ошибка при запуске команды:", err)
			return
		}
	}
}
