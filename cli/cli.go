/*
Package cli implements all CLI functions and features
*/
package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/richardnascimento18/forge/log"
)

func OnBootAnalyzer() {
	testLog := log.Log{Timestamp: time.Now(), Action: "This is an action", Layer: "CLI LAYER", Message: "test log", State: "all good"}
	if err := testLog.WriteLogFile(); err != nil {
		PrintErrorMessage("Could not write log file: " + err.Error())
	}

	args := os.Args[1:]
	if len(args) == 0 {
		PrintErrorMessage("usage: forge run <config-file>")
	}

	cmd := args[0]
	cmdArgs := args[1:]

	switch cmd {
	case "help":
		printHelp()
	case "run":
		handleRun(cmdArgs)
	}
}

func printHelp() {
	fmt.Fprint(os.Stdout, `Usage: forge <command> [options]

Description:
	Declarative task execution engine for reproducible workflows.

Commands:
	init        Initialize a Forge project
	run         Execute a job
	status      Show system status
	help        Show help

Options:
	-h, --help      Show help
	-v, --version   Show version
`)
}

func handleRun(args []string) {
	if len(args) != 1 {
		PrintErrorMessage("no config file has been passed: usage: forge run <config-file>")
	}

	configPath := args[0]

	if _, err := os.Stat(configPath); err != nil {
		PrintErrorMessage("cannot read config file: " + err.Error())
	}
}

func PrintErrorMessage(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
