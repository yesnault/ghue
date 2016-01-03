package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION of ghue
const VERSION = "0.1.0"

// Cmd version
var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Display Version of ghue: ghue version",
	Long:    `ghue version`,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version ghue: %s\n", VERSION)
	},
}
