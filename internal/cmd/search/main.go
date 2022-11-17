package search

import (
	"fmt"
	"strings"
	"zut/internal/commands"
	"zut/internal/commands/file"

	"github.com/spf13/cobra"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	cmds := searchCommands(args)

	for _, c := range cmds {
		fmt.Println(c.Name, "-", c.Description)
	}
}

func searchCommands(args []string) []*commands.Command {
	var commands []*commands.Command

	for _, fileCommand := range file.Commands {
		for _, arg := range args {
			if strings.Contains(fileCommand.Description, arg) {
				commands = append(commands, fileCommand)
				break
			}
		}
	}

	return commands
}
