package main

import (
	"os/exec"
	"runtime"
)

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = exec.Command("cmd").Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = exec.Command("bash").Stdout
		cmd.Run()
	}
}
