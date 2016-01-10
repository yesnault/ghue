package sensors

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// GetSensor GET on /api/<username>/sensors/<id>
func GetSensor(connection *common.Connection, id string) (*Sensor, *common.ErrorHUE, error) {
	sensor := &Sensor{}
	path := fmt.Sprintf("/api/" + connection.Username + "/sensors/" + id)
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return sensor, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return sensor, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &sensor)
	if err != nil {
		log.Errorf("Error with unmarshalling GetSensor: %s", err.Error())
		return sensor, nil, err
	}
	return sensor, nil, nil
}
