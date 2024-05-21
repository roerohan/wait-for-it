package main

import (
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// Execute is used to run a command and print the value in stdout and stderr.
//
// The return value contains the command's exit code.
func Execute(command []string) int {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		Log("unable to start " + err.Error())
		return -1
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	err = cmd.Process.Signal(os.Interrupt)
	if err != nil {
		Log("unable send interrupt " + err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
		return -1
	}
	return 0
}
