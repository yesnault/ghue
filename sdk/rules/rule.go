package rules

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// GetRule GET on /api/<username>/rules/<id>
func GetRule(connection *common.Connection, id string) (*Rule, *common.ErrorHUE, error) {
	rule := &Rule{}
	path := fmt.Sprintf("/api/" + connection.Username + "/rules/" + id)
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return rule, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return rule, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &rule)
	if err != nil {
		log.Errorf("Error with unmarshalling GetRule: %s", err.Error())
		return rule, nil, err
	}
	return rule, nil, nil
}
