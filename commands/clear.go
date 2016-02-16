package commands

import (
	"fmt"
	"log"
	"os"
)

// clearFiles removes all local instances of files downloaded from a remote
func clearFiles() {
	err := os.RemoveAll("/tmp/ssedit")
	if err != nil {
		log.Fatal(err)
	}
}

// clearCmd defines the "clear" subcommand
var clearCmd = &Command{
	Usage: func(cmd string) {
		fmt.Printf("Usage: %s\n\n", cmd)
	},
	RunMethod: func(args []string) error {
		clearFiles()
		return nil
	},
}
