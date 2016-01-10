package schedules

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// GetSchedule GET on /api/<username>/schedules/<id>
func GetSchedule(connection *common.Connection, id string) (*Schedule, *common.ErrorHUE, error) {
	schedule := &Schedule{}
	path := fmt.Sprintf("/api/" + connection.Username + "/schedules/" + id)
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return schedule, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return schedule, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &schedule)
	if err != nil {
		log.Errorf("Error with unmarshalling GetSchedule: %s", err.Error())
		return schedule, nil, err
	}
	return schedule, nil, nil
}
