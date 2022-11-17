package cmd

import (
	"zut/internal/cmd/search"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search into the list of commands",
	Run:   search.Run,
}
