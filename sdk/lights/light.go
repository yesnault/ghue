package lights

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// GetLight GET on /api/<username>/lights/<id>
func GetLight(connection *common.Connection, id string) (*Light, *common.ErrorHUE, error) {
	light := &Light{}
	path := fmt.Sprintf("/api/" + connection.Username + "/lights/" + id)
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return light, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return light, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &light)
	if err != nil {
		log.Errorf("Error with unmarshalling GetLight: %s", err.Error())
		return light, nil, err
	}
	return light, nil, nil
}
