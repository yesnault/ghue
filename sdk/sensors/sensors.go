package sensors

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// Sensor struct
type Sensor struct {
	Config struct {
		Lat           string `json:"lat"`
		Long          string `json:"long"`
		On            bool   `json:"on"`
		Sunriseoffset int    `json:"sunriseoffset"`
		Sunsetoffset  int    `json:"sunsetoffset"`
	} `json:"config"`
	Manufacturername string `json:"manufacturername"`
	Modelid          string `json:"modelid"`
	Name             string `json:"name"`
	State            struct {
		Daylight    interface{} `json:"daylight"`
		Lastupdated string      `json:"lastupdated"`
	} `json:"state"`
	Swversion string `json:"swversion"`
	Type      string `json:"type"`
}

// GetAllSensors GET on /api/<username>/sensors
func GetAllSensors(connection *common.Connection) (map[string]*Sensor, *common.ErrorHUE, error) {
	sensors := map[string]*Sensor{}
	path := fmt.Sprintf("/api/" + connection.Username + "/sensors")
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return sensors, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return sensors, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &sensors)
	if err != nil {
		log.Errorf("Error with unmarshalling GetAllSensors: %s", err.Error())
		return sensors, nil, err
	}
	return sensors, nil, nil
}
