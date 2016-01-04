package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/sdk/common"
)

// Cmd version
var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Display Version of ghue: ghue version",
	Long:    `ghue version`,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version ghue: %s\n", common.VERSION)
	},
}
