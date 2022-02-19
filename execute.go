package main

import (
	"os"
	"os/exec"
)

// Execute is used to run a command and print
// the value in stdout and stderr
func Execute(command []string) {
	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
