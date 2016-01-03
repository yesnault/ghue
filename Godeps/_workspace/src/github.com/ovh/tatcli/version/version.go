package version

import (
	"fmt"

	"github.com/ovh/tatcli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// VERSION of tatcli
const VERSION = "0.56.1"

var versionNewLine bool

func init() {
	Cmd.Flags().BoolVarP(&versionNewLine, "versionNewLine", "", true, "New line after version number. If true, display Version Engine too")
}

// Cmd version
var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Display Version of tatcli and tat engine if configured : tatcli version",
	Long:    `tatcli version`,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		if versionNewLine {
			fmt.Printf("Version tatcli : %s\n", VERSION)
			internal.ReadConfig()
			if viper.GetString("url") == "" {
				internal.Exit("Version Engine : No Engine Configured. See tatcli config --help\n")
			} else {
				fmt.Printf("Version Engine on %s : %s\n", viper.GetString("url"), internal.GetWantReturn("/version"))
			}
		} else {
			fmt.Print(VERSION)
		}
	},
}
