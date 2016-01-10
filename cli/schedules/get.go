package schedules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/cli/config"
	"github.com/yesnault/ghue/cli/internal"
	"github.com/yesnault/ghue/sdk/schedules"
)

var cmdSchedulesGet = &cobra.Command{
	Use:     "get",
	Short:   "Get schedule attributes and state: ghue schedule <id>",
	Long:    `Get schedule attributes and state: ghue schedule <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue schedules state --help")
		} else {
			getCmd(args[0])
		}
	},
}

func getCmd(id string) {
	connection := config.ReadConfig()
	result, errHUE, err := schedules.GetSchedule(connection, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
