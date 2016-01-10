package sensors

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/cli/config"
	"github.com/yesnault/ghue/cli/internal"
	"github.com/yesnault/ghue/sdk/sensors"
)

var cmdSensorsGet = &cobra.Command{
	Use:     "get",
	Short:   "Get sensor attributes and state: ghue sensor <id>",
	Long:    `Get sensor attributes and state: ghue sensor <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue sensors state --help")
		} else {
			getCmd(args[0])
		}
	},
}

func getCmd(id string) {
	connection := config.ReadConfig()
	result, errHUE, err := sensors.GetSensor(connection, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
