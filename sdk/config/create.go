package config

import (
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// Create represents body struct of http://www.developers.meethue.com/documentation/configuration-api#71_create_user
type Create struct {
	DeviceType string `json:"devicetype"`
	// Username is deprecated
	// Username   string `json:"username"`
}

// CreateResult struct
type CreateResult struct {
	Success struct {
		Username string `json:"username"`
	} `json:"success"`
}

// CreateAPI POST on /api to create a new user
func CreateAPI(connection *common.Connection, create *Create) (*CreateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(create)
	if err != nil {
		log.Errorf("Error with marshalling create: %s", err.Error())
		return &CreateResult{}, nil, err
	}
	bodyResponse, errHUE, err := internal.Request(connection, "POST", http.StatusOK, "/api/", bodyRequest)
	if errHUE != nil {
		log.Errorf("Error with requesting POST on /api (create a new user), HUE Error: %s", errHUE.Error.Description)
		return &CreateResult{}, errHUE, err
	}
	if err != nil {
		log.Errorf("Error with requesting POST on /api (create a new user): %s", err.Error())
		return &CreateResult{}, errHUE, err
	}
	var creates []CreateResult
	err = json.Unmarshal(bodyResponse, &creates)
	if err != nil {
		log.Errorf("Error with unmarshalling POST on /api (create a new user): %s", err.Error())
		return &CreateResult{}, nil, err
	}
	return &creates[0], nil, nil

}
