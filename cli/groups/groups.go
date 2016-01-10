package groups

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdGroupsAll)
	Cmd.AddCommand(cmdGroupsGet)
}

// Cmd groups
var Cmd = &cobra.Command{
	Use:     "groups",
	Short:   "Groups commands: ghue groups --help",
	Long:    `Groups commands: ghue groups <command>`,
	Aliases: []string{"group", "g"},
}
