package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Execute is used to run a command and print
// the value in stdout and stderr
func Execute(command []string) {
	var out, stderr bytes.Buffer

	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(stderr.String())
		return
	}

	fmt.Println(out.String())
}
