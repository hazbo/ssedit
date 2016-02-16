package commands

// Command represents the definition of a single command.
type Command struct {
	Usage     func(string)
	RunMethod func([]string) error
}

// Run executes the command with the given args.
func (cmd *Command) Run(args []string) error {
	return cmd.RunMethod(args)
}

// CommandMap defines all of the available subcommands.
var CommandMap = map[string]*Command{
	"clear":   clearCmd,
	"open":    openCmd,
	"version": versionCmd,
}
