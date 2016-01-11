package rules

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// Rule struct
type Rule struct {
	Actions []struct {
		Address string `json:"address"`
		Body    struct {
			BriInc         int `json:"bri_inc"`
			Transitiontime int `json:"transitiontime"`
		} `json:"body"`
		Method string `json:"method"`
	} `json:"actions"`
	Conditions []struct {
		Address  string `json:"address"`
		Operator string `json:"operator"`
		Value    string `json:"value"`
	} `json:"conditions"`
	Created        string `json:"created"`
	Lasttriggered  string `json:"lasttriggered"`
	Name           string `json:"name"`
	Owner          string `json:"owner"`
	Status         string `json:"status"`
	Timestriggered int    `json:"timestriggered"`
}

// GetAllRules GET on /api/<username>/rules
func GetAllRules(connection *common.Connection) (map[string]*Rule, *common.ErrorHUE, error) {
	rules := map[string]*Rule{}
	path := fmt.Sprintf("/api/" + connection.Username + "/rules")
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return rules, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return rules, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &rules)
	if err != nil {
		log.Errorf("Error with unmarshalling GetAllRules: %s", err.Error())
		return rules, nil, err
	}
	return rules, nil, nil
}
