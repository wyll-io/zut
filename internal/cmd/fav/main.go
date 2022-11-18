package fav

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type FavCommands []string

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Displaying all your favorite commands")
	fmt.Println()

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Read file
	data, err := os.ReadFile(dirname + "/zut.yml")
	if err != nil {
		log.Fatal(err)
	}

	var favCommands *FavCommands

	err = yaml.Unmarshal(data, &favCommands)
	if err != nil {
		log.Fatal(err)
	}

	for _, favCommand := range *favCommands {
		fmt.Println(favCommand)
	}
}
