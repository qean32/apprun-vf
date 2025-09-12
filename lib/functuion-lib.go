package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

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

func write(write []string, path string) error {
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
