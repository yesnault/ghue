package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	// Verbose conditions the quantity of output of api requests
	Verbose bool

	// ConfigFile is $HOME/.tatcli/config.json per default
	// contains user, password and url of tat
	ConfigFile string

	// SSLInsecureSkipVerify Skip certificate check with SSL connection
	SSLInsecureSkipVerify bool

	// Pretty prints json return in pretty format
	Pretty bool

	// URL of tat engine
	URL string

	// Username of tat user
	Username string

	// Password of tat user
	Password string
)

// ReadConfig reads config in .tatcli/config per default
func ReadConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
		viper.ReadInConfig() // Find and read the config file
		if Verbose {
			fmt.Printf("Using config file %s\n", ConfigFile)
		}
	}
}
