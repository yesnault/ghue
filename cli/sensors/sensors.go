package sensors

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdSensorsAll)
	Cmd.AddCommand(cmdSensorsGet)
}

// Cmd sensors
var Cmd = &cobra.Command{
	Use:     "sensors",
	Short:   "Sensors commands: ghue sensors --help",
	Long:    `Sensors commands: ghue sensors <command>`,
	Aliases: []string{"sensor", "g"},
}
