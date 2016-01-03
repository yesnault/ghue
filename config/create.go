package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/internal"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/config"
	"github.com/yesnault/ghue/version"
)

var (
	save bool
	ip   string
)

func init() {
	cmdConfigRegister.Flags().BoolVar(&save, "save", false, "Save configuration after verify in $HOME/.ghue/config.json")
	cmdConfigRegister.Flags().StringVar(&ip, "ip", "", "IP of your bridge. See http://www.developers.meethue.com/documentation/getting-started to check how to discover your bridge's IP.")
}

var cmdConfigRegister = &cobra.Command{
	Use:   "create",
	Short: "Creates a new user, don't forget to press the button on the bridge: ghue config create [--save]",
	Long: `
	Creates a new user. The link button on the bridge must be pressed and this command executed within 30 seconds.

	Usage:
	Creates a new user and save configuration in $HOME/.ghue/config.json
	ghue config create --save

	`,
	Run: func(cmd *cobra.Command, args []string) {
		createCmd()
	},
}

func createCmd() {
	connection := &common.Connection{
		Host:    ip,
		Verbose: internal.Verbose,
	}
	create := &config.Create{
		DeviceType: "ghue" + version.VERSION,
	}

	createResult, errHUE, err := config.CreateAPI(connection, create)
	internal.CheckErrors(err, errHUE)

	output, err := json.Marshal(createResult)
	internal.Check(err)

	if save {
		ghueFile := &GHUEFile{
			Host:     connection.Host,
			Username: createResult.Success.Username,
			Version:  version.VERSION,
		}
		jsonStr, err := json.MarshalIndent(ghueFile, "", "  ")
		internal.Check(err)
		jsonStr = append(jsonStr, '\n')
		dir := path.Dir(ConfigFile)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			internal.Check(os.Mkdir(dir, 0740))
		}
		internal.Check(ioutil.WriteFile(ConfigFile, jsonStr, 0600))
		fmt.Printf("file %s is created\n", ConfigFile)
	}

	fmt.Println(string(output))
}
