package lights

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/yesnault/ghue/sdk/common"
	"github.com/yesnault/ghue/sdk/internal"
)

// State of a light
type State struct {
	Alert     string    `json:"alert"`
	Bri       int       `json:"bri"`
	Colormode string    `json:"colormode"`
	Ct        int       `json:"ct"`
	Effect    string    `json:"effect"`
	Hue       int       `json:"hue"`
	On        bool      `json:"on"`
	Reachable bool      `json:"reachable"`
	Sat       int       `json:"sat"`
	Xy        []float64 `json:"xy"`
}

// SetStateValues of a light
type SetStateValues struct {
	On             string `json:"on"`
	Alert          string `json:"alert"`
	Bri            string `json:"bri"`
	Hue            string `json:"hue"`
	Sat            string `json:"sat"`
	XY             string `json:"xy"`
	Ct             string `json:"ct"`
	Effect         string `json:"effect"`
	TransitionTime string `json:"transitiontime"`
	BriInc         string `json:"bri_inc"`
	SatInc         string `json:"sat_inc"`
	HueInc         string `json:"hue_inc"`
	CtInc          string `json:"ct_inc"`
	XYInc          string `json:"xy_inc"`
}

// SetState PUT on /api/<username>/lights/<idLight>
func SetState(connection *common.Connection, id string, setState *SetStateValues) ([]interface{}, *common.ErrorHUE, error) {

	var ret []interface{}
	requestBody := make(map[string]interface{})

	if setState.Alert != "" {
		requestBody["alert"] = setState.Alert
	}
	if setState.Effect != "" {
		requestBody["effect"] = setState.Effect
	}
	if v, err := strconv.ParseBool(setState.On); err == nil {
		requestBody["on"] = v
	}

	toInt("bri", setState.Bri, requestBody)
	toInt("hue", setState.Hue, requestBody)
	toInt("sat", setState.Sat, requestBody)
	//toInt("xy", setState.XY, requestBody)
	toInt("ct", setState.Ct, requestBody)
	toInt("transitiontime", setState.TransitionTime, requestBody)
	toInt("bri_inc", setState.BriInc, requestBody)
	toInt("sat_inc", setState.SatInc, requestBody)
	toInt("hue_inc", setState.HueInc, requestBody)
	toInt("ct_inc", setState.CtInc, requestBody)
	toInt("xy_inc", setState.XYInc, requestBody)
	if setState.XY != "" {
		tuple := strings.Split(setState.XY, ",")
		if len(tuple) == 2 {
			x, e1 := strconv.ParseFloat(tuple[0], 64)
			y, e2 := strconv.ParseFloat(tuple[0], 64)
			if e1 == nil && e2 == nil {
				requestBody["xy"] = []float64{x, y}
			} else {
				return ret, nil, fmt.Errorf("Invalid value for xy")
			}
		}
	}

	bodyRequest, err := json.Marshal(requestBody)
	if err != nil {
		log.Errorf("Error with marshalling state: %s", err.Error())
		return ret, nil, err
	}

	path := fmt.Sprintf("/api/" + connection.Username + "/lights/" + id + "/state")
	bodyResponse, errHUE, err := internal.Request(connection, "PUT", http.StatusOK, path, bodyRequest)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return ret, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return ret, errHUE, err
	}

	err = json.Unmarshal(bodyResponse, &ret)
	if err != nil {
		log.Errorf("Error with unmarshalling SetState: %s", err.Error())
		return ret, nil, err
	}
	return ret, nil, nil
}

func toInt(key string, value string, requestBody map[string]interface{}) {
	if v, err := strconv.Atoi(value); err == nil {
		requestBody[key] = v
	}
}
