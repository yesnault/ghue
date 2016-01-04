package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yesnault/ghue/cli/internal"
	"github.com/yesnault/ghue/sdk/common"
)

func init() {
	Cmd.AddCommand(cmdConfigRegister)
	Cmd.AddCommand(cmdConfigGet)
}

// Cmd config
var Cmd = &cobra.Command{
	Use:     "config",
	Short:   "Config commands: ghue config --help",
	Long:    `Config commands: ghue config <command>`,
	Aliases: []string{"c"},
}

var (
	// ConfigFile is $HOME/.ghue/config.json per default
	// contains Host and username
	ConfigFile string
)

// ReadConfig reads config in .ghue/config per default
func ReadConfig() *common.Connection {
	connection := &common.Connection{}
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
		viper.ReadInConfig() // Find and read the config file
		if internal.Verbose {
			fmt.Printf("Using config file %s\n", ConfigFile)
		}

		connection = &common.Connection{
			Host:     viper.GetString("host"),
			Username: viper.GetString("username"),
			Verbose:  internal.Verbose,
		}
	}
	return connection
}

// GHUEFile struct
type GHUEFile struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Version  string `json:"version"`
}
