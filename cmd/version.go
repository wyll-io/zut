package cmd

import (
	"zut/internal/cmd/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version",
	Run:   version.Run,
}
