package scenes

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// Scene struct
type Scene struct {
	Appdata     struct{}    `json:"appdata"`
	Lastupdated interface{} `json:"lastupdated"`
	Lights      []string    `json:"lights"`
	Locked      bool        `json:"locked"`
	Name        string      `json:"name"`
	Owner       string      `json:"owner"`
	Picture     string      `json:"picture"`
	Recycle     bool        `json:"recycle"`
	Version     int         `json:"version"`
	Lightstates struct {
		One struct {
			Bri int       `json:"bri"`
			On  bool      `json:"on"`
			Xy  []float64 `json:"xy"`
		} `json:"1"`
	} `json:"lightstates,omitempty"`
}

// GetAllScenes GET on /api/<username>/scenes
func GetAllScenes(connection *common.Connection) (map[string]*Scene, *common.ErrorHUE, error) {
	scenes := map[string]*Scene{}
	path := fmt.Sprintf("/api/" + connection.Username + "/scenes")
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return scenes, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return scenes, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &scenes)
	if err != nil {
		log.Errorf("Error with unmarshalling GetAllScenes: %s", err.Error())
		return scenes, nil, err
	}
	return scenes, nil, nil
}
