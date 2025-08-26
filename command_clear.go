package main

import (
	"os"
	"os/exec"
	"runtime"
)

func commandClear(config *config, args ...string) error {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	return nil
}
