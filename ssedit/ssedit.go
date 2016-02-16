package main

import (
	"fmt"
	"github.com/hazbo/ssedit/commands"
	"os"
	"os/signal"
)

func printHelp() {
	fmt.Printf(`usage: ssedit <command> [<args>]

individual usage:
help <command>

commands:
clear      Clears the local instances of remote files
help       Displays this text or command individual usage
open       Opens a remote file or directory ready to edit
version    Displays the version of ssedit

`)
}

func main() {
	// Clear /tmp/ssedit on Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		subcommand := commands.CommandMap["clear"]
		subcommand.RunMethod([]string{})
		os.Exit(0)
	}()

	if len(os.Args) == 1 {
		printHelp()
	}
	if len(os.Args) == 2 && os.Args[1] == "help" {
		printHelp()
	}
	// TODO: check that subcommand actually exists
	if len(os.Args) == 2 && os.Args[1] != "help" {
		commands.CommandMap[os.Args[1]].RunMethod([]string{})
	}
	if len(os.Args) > 2 {
		passedCommand := os.Args[1]

		if passedCommand == "help" {
			// Print the usage for the given subcommand
			commands.CommandMap[os.Args[2]].Usage(os.Args[2])
			os.Exit(0)
		}

		var args []string
		for i := 2; i < len(os.Args); i++ {
			args = append(args, os.Args[i])
		}

		subcommand := commands.CommandMap[passedCommand]
		subcommand.RunMethod(args)
	}
}
