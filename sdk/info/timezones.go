package info

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// GetAllTimezones GET on /api/<username>/info/timezones
func GetAllTimezones(connection *common.Connection) ([]string, *common.ErrorHUE, error) {
	timezones := []string{}
	path := fmt.Sprintf("/api/" + connection.Username + "/info/timezones")
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return timezones, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return timezones, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &timezones)
	if err != nil {
		log.Errorf("Error with unmarshalling GetAllTimezones: %s", err.Error())
		return timezones, nil, err
	}
	return timezones, nil, nil
}
