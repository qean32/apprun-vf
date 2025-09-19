package main

import (
	"fmt"
	"os/exec"
)

func mainF() {
	// Apprun_(`start-process 'D:\Desktop\zapret-discord-youtube-1.8.4\general(ALT2).bat'`)
	// Apprun_(`cmd.exe -/c 'D:\Desktop\zapret-discord-youtube-1.8.4\general(ALT2).bat'`)
	// Apprun_(`Invoke-Expression -Command 'D:\Desktop\zapret-discord-youtube-1.8.4\general(ALT2).bat'`)
	// Apprun_(`powershell.exe -Command 'D:\Desktop\zapret-discord-youtube-1.8.4\general(ALT2).bat'`)
	// Apprun_(`cmd.exe '/C' 'D:\Desktop\zapret-discord-youtube-1.8.4\general(ALT2).bat'`)
	// command := exec.Command("CMD.exe", "/C", `D:\Desktop\zapret-discord-youtube-1.8.4\general(ALT2).bat`)
	// command.Run()
	// runScript("", "D:/Desktop/zapret-discord-youtube-1.8.4/general(ALT2).bat")
	cmd := exec.Command("CMD.exe", "/C", "D:/Desktop/zapret-discord-youtube-1.8.4/general(ALT2).bat")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске команды:", err)
		return
	}
}

func Apprun_(path string) {
	cmd := exec.Command(path)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при запуске команды:", err)
		return
	}
}

// func runScript(folder, filename string) bool {
// 	fmt.Println("Running", folder, filename)
// 	cmd := exec.Command("start-process " + filename)
// 	cmd.Dir = folder
// 	// Stream to std out
// 	cmd.Stdout = os.Stdout
// 	cmd.Stdin = os.Stdin
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Run()
// 	if err != nil && err.Error() != "exit status 1" {
// 		fmt.Println("Error running script: " + err.Error())
// 		return false
// 	}
// 	return true
// }
