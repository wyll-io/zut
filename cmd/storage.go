package cmd

import (
	"zut/internal/cmd/storage"

	"github.com/spf13/cobra"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "List all storage commands",
	Run:   storage.Run,
}
