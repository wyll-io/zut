package cmd

import (
	"zut/internal/cmd/fav"

	"github.com/spf13/cobra"
)

var favCmd = &cobra.Command{
	Use:   "fav",
	Short: "List all favorites commands",
	Run:   fav.Run,
}
