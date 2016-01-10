package scenes

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// GetScene GET on /api/<username>/scenes/<id>
func GetScene(connection *common.Connection, id string) (*Scene, *common.ErrorHUE, error) {
	scene := &Scene{}
	path := fmt.Sprintf("/api/" + connection.Username + "/scenes/" + id)
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return scene, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return scene, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &scene)
	if err != nil {
		log.Errorf("Error with unmarshalling GetScene: %s", err.Error())
		return scene, nil, err
	}
	return scene, nil, nil
}
