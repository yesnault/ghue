package scenes

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdScenesAll)
	Cmd.AddCommand(cmdScenesGet)
}

// Cmd scenes
var Cmd = &cobra.Command{
	Use:     "scenes",
	Short:   "Scenes commands: ghue scenes --help",
	Long:    `Scenes commands: ghue scenes <command>`,
	Aliases: []string{"scene", "g"},
}
