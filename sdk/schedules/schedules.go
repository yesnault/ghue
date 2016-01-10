package schedules

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// Schedule struct
type Schedule struct {
	Command struct {
		Address string `json:"address"`
		Body    struct {
			Scene string `json:"scene"`
		} `json:"body"`
		Method string `json:"method"`
	} `json:"command"`
	Created     string `json:"created"`
	Description string `json:"description"`
	Localtime   string `json:"localtime"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Time        string `json:"time"`
	Autodelete  bool   `json:"autodelete"`
}

// GetAllSchedules GET on /api/<username>/schedules
func GetAllSchedules(connection *common.Connection) (map[string]*Schedule, *common.ErrorHUE, error) {
	schedules := map[string]*Schedule{}
	path := fmt.Sprintf("/api/" + connection.Username + "/schedules")
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return schedules, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return schedules, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &schedules)
	if err != nil {
		log.Errorf("Error with unmarshalling GetAllSchedules: %s", err.Error())
		return schedules, nil, err
	}
	return schedules, nil, nil
}
