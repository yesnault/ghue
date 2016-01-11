package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/cli/config"
	"github.com/yesnault/ghue/cli/internal"
	"github.com/yesnault/ghue/sdk/rules"
)

var cmdRulesGet = &cobra.Command{
	Use:     "get",
	Short:   "Get rule attributes and state: ghue rule <id>",
	Long:    `Get rule attributes and state: ghue rule <id>`,
	Aliases: []string{"show", "g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue rules state --help")
		} else {
			getCmd(args[0])
		}
	},
}

func getCmd(id string) {
	connection := config.ReadConfig()
	result, errHUE, err := rules.GetRule(connection, id)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
