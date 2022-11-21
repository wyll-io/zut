package search

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"zut/internal/commands"
	"zut/internal/commands/file"
	"zut/pkg/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	cmds := searchCommands(args)

	if len(cmds) == 0 {
		fmt.Println("no command has been found")
		return
	}

	for _, c := range cmds {
		// Check if command is installed
		isInstalled := utils.CommandIsInstalled(c.Name)

		if isInstalled {
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Fprintf(color.Output, "%s - %s (Command is %s)", c.Name, c.Description, green("installed"))
		} else {
			red := color.New(color.FgRed).SprintFunc()
			fmt.Fprintf(color.Output, "%s - %s (Command is %s)", c.Name, c.Description, red("not installed"))

			if utils.AskForConfirmation("Do you want to install it ?") {
				cmd := exec.Command("apt", "install", c.Name)
				out, err := cmd.Output()
				if err != nil {
					log.Fatalf("running apt list: %s", err)
				}

				fmt.Print(out)
			}
		}

		fmt.Println()
	}
}

func searchCommands(args []string) []*commands.Command {
	var commands []*commands.Command

	for _, fileCommand := range file.Commands {
		for _, arg := range args {
			// Search in description
			if strings.Contains(fileCommand.Description, arg) {
				commands = append(commands, fileCommand)
				break
			}

			// Search in tags
			for _, tag := range fileCommand.Tags {
				if strings.Contains(tag, arg) {
					commands = append(commands, fileCommand)
					break
				}
			}
		}
	}

	return commands
}
