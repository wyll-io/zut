package search

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"zut/internal/commands"
	"zut/internal/commands/file"
	"zut/pkg/utils"

	"github.com/spf13/cobra"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	cmds := searchCommands(args)

	for _, c := range cmds {
		fmt.Println(c.Name, "-", c.Description)
		if utils.CommandExists(c.Name) {
			fmt.Println("Command is installed")
		} else {
			fmt.Println("Command is not installed")
			if utils.AskForConfirmation("Do you want to install it ?") {
				cmd := exec.Command("apt", "install", c.Name)
				out, err := cmd.Output()
				if err != nil {
					log.Fatalf("running apt list: %s", err)
				}

				fmt.Print(out)
			}
		}
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
