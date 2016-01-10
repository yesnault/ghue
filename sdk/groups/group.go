package groups

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// GetGroup GET on /api/<username>/groups/<id>
func GetGroup(connection *common.Connection, id string) (*Group, *common.ErrorHUE, error) {
	group := &Group{}
	path := fmt.Sprintf("/api/" + connection.Username + "/groups/" + id)
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return group, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return group, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &group)
	if err != nil {
		log.Errorf("Error with unmarshalling GetGroup: %s", err.Error())
		return group, nil, err
	}
	return group, nil, nil
}
