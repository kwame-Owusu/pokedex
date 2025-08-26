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
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
