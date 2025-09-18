package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"

	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const NEW_FILE_NAME = "Untited"

func main_() {
	a := app.New()
	w := a.NewWindow("apprun-vf")
	w.Resize(fyne.NewSize(900, 600))
	w.SetFixedSize(true)
	tabs := AddTabs(w)
	entries := AddFirstTab(tabs)
	AddProgramMenu(w, tabs, entries)

	icon, _ := fyne.LoadResourceFromPath("./public/logo-white.svg")
	w.SetIcon(icon)
	w.ShowAndRun()
}

func Read_(path string) ([]string, error) {
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

func NewTextEntery() *widget.Entry {
	textInput := widget.NewMultiLineEntry()
	textInput.SetPlaceHolder("Enter text here... n/")
	content, err := os.ReadFile("./run/game.txt")
	if err != nil {
	}
	textInput.SetText(string(content))

	return textInput
}

func AddTabs(w fyne.Window) *container.AppTabs {
	tabs := container.NewAppTabs()
	w.SetContent(tabs)

	return tabs
}

func AddFirstTab(tabs *container.AppTabs) []*widget.Entry {
	entries := []*widget.Entry{}
	firstEntery := NewTextEntery()
	tabs.Append(container.NewTabItem(NEW_FILE_NAME, firstEntery))

	return append(entries, firstEntery)
}

func AddProgramMenu(w fyne.Window, tabs *container.AppTabs, entries []*widget.Entry) {
	new := fyne.NewMenuItem("New", func() {
		entry := NewTextEntery()
		entries = append(entries, entry)

		tabs.Append(container.NewTabItem(NEW_FILE_NAME, entry))
	})

	save := fyne.NewMenuItem("Save", func() {
		tabText := tabs.Selected().Text

		if tabText != NEW_FILE_NAME {
			file, err := os.Create(tabText)

			if err != nil {
				fmt.Println("content")
			}

			defer file.Close()
			file.WriteString(entries[tabs.SelectedIndex()].Text)
		} else {
			SaveFile(w, tabs, entries)
		}
	})

	saveAs := fyne.NewMenuItem("Save As", func() {
		SaveFile(w, tabs, entries)
	})

	open := fyne.NewMenuItem("Open", func() {

	})

	menu := fyne.NewMenu("File", new, save, saveAs, open)
	w.SetMainMenu(fyne.NewMainMenu(menu))
}

func SaveFile(w fyne.Window, tabs *container.AppTabs, entries []*widget.Entry) {
	dialog.ShowFileSave(
		func(writer fyne.URIWriteCloser, err error) {
			if writer != nil {

				io.WriteString(writer, entries[tabs.SelectedIndex()].Text)
				tabs.Selected().Text = writer.URI().Path()
				tabs.Refresh()
			}
		},
		w,
	)
}
