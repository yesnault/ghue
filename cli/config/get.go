package config

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/cli/internal"
	"github.com/yesnault/ghue/sdk/config"
)

var cmdConfigGet = &cobra.Command{
	Use:     "get",
	Short:   "Get configuration: ghue configuration get",
	Long:    `Get configuration: ghue configuration get`,
	Aliases: []string{"show"},
	Run: func(cmd *cobra.Command, args []string) {
		getCmd()
	},
}

func getCmd() {
	connection := ReadConfig()
	result, errHUE, err := config.Get(connection)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
