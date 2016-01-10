package schedules

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdSchedulesAll)
	Cmd.AddCommand(cmdSchedulesGet)
}

// Cmd schedules
var Cmd = &cobra.Command{
	Use:     "schedules",
	Short:   "Schedules commands: ghue schedules --help",
	Long:    `Schedules commands: ghue schedules <command>`,
	Aliases: []string{"schedule", "g"},
}
