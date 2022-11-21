package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.0.1"

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	fmt.Println("zut version:", Version)
}
