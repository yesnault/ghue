package lights

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdLightsAll)
	Cmd.AddCommand(cmdLightsState)
	Cmd.AddCommand(cmdLightsGet)
}

// Cmd lights
var Cmd = &cobra.Command{
	Use:     "lights",
	Short:   "Lights commands: ghue lights --help",
	Long:    `Lights commands: ghue lights <command>`,
	Aliases: []string{"light", "l"},
}
