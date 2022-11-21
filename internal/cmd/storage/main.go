package storage

import (
	"fmt"
	"log"
	"os/exec"

	"zut/internal/commands/storage"
	"zut/pkg/utils"

	"github.com/spf13/cobra"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	for _, c := range storage.Commands {
		fmt.Println(c.Name, "-", c.Description)
		if utils.CommandIsInstalled(c.Name) {
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
