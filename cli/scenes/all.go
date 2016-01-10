package scenes

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/cli/config"
	"github.com/yesnault/ghue/cli/internal"
	"github.com/yesnault/ghue/sdk/scenes"
)

var cmdScenesAll = &cobra.Command{
	Use:   "all",
	Short: "Get All scenes: ghue scenes all",
	Long:  `Get all scenes: ghue scenes all`,
	Run: func(cmd *cobra.Command, args []string) {
		allCmd()
	},
}

func allCmd() {
	connection := config.ReadConfig()
	result, errHUE, err := scenes.GetAllScenes(connection)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
