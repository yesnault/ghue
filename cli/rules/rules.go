package rules

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdRulesAll)
	Cmd.AddCommand(cmdRulesGet)
}

// Cmd rules
var Cmd = &cobra.Command{
	Use:     "rules",
	Short:   "Rules commands: ghue rules --help",
	Long:    `Rules commands: ghue rules <command>`,
	Aliases: []string{"rule", "g"},
}
