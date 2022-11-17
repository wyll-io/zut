package cmd

import (
	"os"

	"zut/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zut",
	Short: "Zut helps you to find the command that you search and how to use",
	Long:  ``,
	Run:   internal.Run,
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

// Execute the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
