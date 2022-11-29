package main

import (
	"flag"
	"fmt"
	"github.com/roerohan/wait-for-it/pkg/wait"
	"os"
)

var (
	timeout  int
	services wait.Services
	quiet    bool
	strict   bool
)

func init() {
	flag.IntVar(&timeout, "t", 15, "Timeout in seconds, zero for no timeout")
	flag.BoolVar(&quiet, "q", false, "Quiet, don't output any status messages")
	flag.BoolVar(&strict, "s", false, "Only execute subcommand if the test succeeds")
	flag.Var(&services, "w", "Services to be waiting for, in the form `host:port`")
}

// Log is used to log with prefix wait-for-it:
func log(message string) {
	if quiet {
		return
	}

	wait.Log("wait-for-it: " + message)
}

func main() {
	flag.Parse()
	args := os.Args

	if len(services) != 0 {
		log(fmt.Sprintf("waiting %d seconds for %s", timeout, services.String()))
		ok := services.Wait(timeout)

		if !ok {
			log(fmt.Sprintf("timeout occured after waiting for %d seconds", timeout))
			if strict {
				log("strict mode, refusing to execute subprocess")
				return
			}
		}
	}

	var command []string

	for i, arg := range args {
		if arg == "--" {
			if (i + 1) < len(args) {
				command = args[i+1:]
			}

			break
		}
	}

	if len(command) == 0 {
		return
	}

	Execute(command)
}
