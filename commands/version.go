package commands

import (
	"fmt"
	"github.com/hazbo/ssedit/meta"
)

// displayVersion displays the current version of ssedit along with a small
// thank you message.
func displayVersion() {
	msg := fmt.Sprintf("%s v%s\nAuthor: %s %s\n",
		meta.Application,
		meta.Version,
		meta.Author,
		meta.Copyright,
	)
	fmt.Printf(msg)
	fmt.Printf(`
Thank you for using ssedit! All source code is available at
github.com/hazbo/ssedit, free under the MIT license.
`)
}

// versionCmd defines the "version" subcommand
var versionCmd = &Command{
	Usage: func(cmd string) {
		fmt.Printf("Usage: %s\n\n", cmd)
	},
	RunMethod: func(args []string) error {
		displayVersion()
		return nil
	},
}
