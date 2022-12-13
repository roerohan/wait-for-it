package main

import (
	"flag"
	"fmt"
	"github.com/roerohan/wait-for-it/pkg/wait"
	"os"
	"time"
)

var (
	reqTimeout int
	maxTimeout int
	services   wait.Services
	quiet      bool
	strict     bool
)

func init() {
	flag.IntVar(&reqTimeout, "t", 15, "Service request timeout in seconds, zero for no timeout")
	flag.IntVar(&maxTimeout, "m", 30, "Max service timeout to retry request in seconds, zero for no max service timeout")
	flag.BoolVar(&quiet, "q", false, "Quiet, don't output any status messages")
	flag.BoolVar(&strict, "s", false, "Only execute subcommand if the test succeeds")
	flag.Var(&services, "w", "Dependency services to be waiting for, in the form `host:port`")
}

// Log is used to log with prefix wait-for-it:
func log(message string) {
	if quiet {
		return
	}

	wait.Log(message)
}

func main() {
	flag.Parse()
	args := os.Args

	if len(services) != 0 {
		log(fmt.Sprintf("waiting %d seconds for %s for a max of %d seconds", reqTimeout, services.String(), maxTimeout))
		err := wait.ForDependencies(services, time.Duration(reqTimeout), time.Duration(maxTimeout))
		if err != nil {
			log(fmt.Sprintf("wait.ForDependencies failed with err %v", err))
			if strict {
				log("strict mode, refusing to execute subprocess")
				os.Exit(1)
			}
			os.Exit(1)
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
