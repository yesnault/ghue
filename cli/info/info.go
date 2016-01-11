package info

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdInfoTimezones)
}

// Cmd info
var Cmd = &cobra.Command{
	Use:     "info",
	Short:   "Info commands: ghue info --help",
	Long:    `Info commands: ghue info <command>`,
	Aliases: []string{"info", "i"},
}
