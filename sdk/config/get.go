package config

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// Config struct
type Config struct {
	Name          string `json:"name"`
	Zigbeechannel int    `json:"zigbeechannel"`
	Bridgeid      string `json:"bridgeid"`
	Mac           string `json:"mac"`
	Dhcp          bool   `json:"dhcp"`
	Ipaddress     string `json:"ipaddress"`
	Netmask       string `json:"netmask"`
	Gateway       string `json:"gateway"`
	Proxyaddress  string `json:"proxyaddress"`
	Proxyport     int    `json:"proxyport"`
	Utc           string `json:"UTC"`
	Localtime     string `json:"localtime"`
	Timezone      string `json:"timezone"`
	Modelid       string `json:"modelid"`
	Swversion     string `json:"swversion"`
	Apiversion    string `json:"apiversion"`
	Swupdate      struct {
		Updatestate    int  `json:"updatestate"`
		Checkforupdate bool `json:"checkforupdate"`
		Devicetypes    struct {
			Bridge  bool          `json:"bridge"`
			Lights  []interface{} `json:"lights"`
			Sensors []interface{} `json:"sensors"`
		} `json:"devicetypes"`
		URL    string `json:"url"`
		Text   string `json:"text"`
		Notify bool   `json:"notify"`
	} `json:"swupdate"`
	Linkbutton       bool   `json:"linkbutton"`
	Portalservices   bool   `json:"portalservices"`
	Portalconnection string `json:"portalconnection"`
	Portalstate      struct {
		Signedon      bool   `json:"signedon"`
		Incoming      bool   `json:"incoming"`
		Outgoing      bool   `json:"outgoing"`
		Communication string `json:"communication"`
	} `json:"portalstate"`
	Factorynew       bool   `json:"factorynew"`
	Replacesbridgeid string `json:"replacesbridgeid"`
	Backup           struct {
		Status    string `json:"status"`
		Errorcode int    `json:"errorcode"`
	} `json:"backup"`
	Whitelist map[string]*Whitelisted `json:"whitelist"`
}

// Whitelisted struct
type Whitelisted struct {
	LastUseDate string `json:"last use date"`
	CreateDate  string `json:"create date"`
	Name        string `json:"name"`
}

// Get GET on /api/<username>/config
// see http://www.developers.meethue.com/documentation/configuration-api#72_get_configuration
func Get(connection *common.Connection) (*Config, *common.ErrorHUE, error) {
	var config Config
	path := fmt.Sprintf("/api/" + connection.Username + "/config")
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return &config, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return &config, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &config)
	if err != nil {
		log.Errorf("Error with unmarshalling GetAllLights: %s", err.Error())
		return &config, nil, err
	}
	return &config, nil, nil
}
