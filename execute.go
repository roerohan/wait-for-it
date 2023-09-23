package main

import (
	"os"
	"os/exec"
)

// Execute is used to run a command and print the value in stdout and stderr.
//
// The return value contains the command's exit code.
func Execute(command []string) int {
	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}

		return -1
	}

	return 0
}
