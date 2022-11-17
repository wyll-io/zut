package cmd

import (
	"zut/internal/cmd/networking"

	"github.com/spf13/cobra"
)

var networkingCmd = &cobra.Command{
	Use:   "networking",
	Short: "List all networking commands",
	Run:   networking.Run,
}
