package lights

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/cli/config"
	"github.com/yesnault/ghue/cli/internal"
	"github.com/yesnault/ghue/sdk/lights"
)

// TODO Add All arguments from http://www.developers.meethue.com/documentation/lights-api#15_set_light_attributes_rename

var (
	on             string
	alert          string
	bri            string
	hue            string
	sat            string
	xy             string
	ct             string
	effect         string
	transitionTime string
	briInc         string
	satInc         string
	hueInc         string
	ctInc          string
	xyInc          string
)

func init() {
	cmdLightsState.Flags().StringVarP(&on, "on", "", "", "On/Off state of the light. On=true, Off=false")
	cmdLightsState.Flags().StringVarP(&alert, "alert", "", "", `The alert effect, is a temporary change to the bulb’s state, and has one of the following values:
				  “none” – The light is not performing an alert effect.
				  “select” – The light is performing one breathe cycle.
				  “lselect” – The light is performing breathe cycles for 15 seconds or until an “alert“: “none“ command is received.`)
	cmdLightsState.Flags().StringVarP(&bri, "bri", "", "", `The brightness value to set the light to.
				  Brightness is a scale from 1 (the minimum the light is capable of) to 254 (the maximum). Note: a brightness of 1 is not off.`)
	cmdLightsState.Flags().StringVarP(&hue, "hue", "", "", `The hue value to set light to.
				  The hue value is a wrapping value between 0 and 65535. Both 0 and 65535 are red, 25500 is green and 46920 is blue.`)
	cmdLightsState.Flags().StringVarP(&sat, "sat", "", "", `Saturation of the light. 254 is the most saturated (colored) and 0 is the least saturated (white).`)
	cmdLightsState.Flags().StringVarP(&xy, "xy", "", "", `The x and y coordinates of a color in CIE color space.
				  The first entry is the x coordinate and the second entry is the y coordinate. Both x and y must be between 0 and 1.
				  If the specified coordinates are not in the CIE color space, the closest color to the coordinates will be chosen.`)
	cmdLightsState.Flags().StringVarP(&ct, "ct", "", "", `The Mired Color temperature of the light. 2012 connected lights are capable of 153 (6500K) to 500 (2000K).`)
	cmdLightsState.Flags().StringVarP(&effect, "effect", "", "", `The dynamic effect of the light. Currently “none” and “colorloop” are supported.
				  Other values will generate an error of type 7.
				  Setting the effect to colorloop will cycle through all hues using the current brightness and saturation settings.`)
	cmdLightsState.Flags().StringVarP(&transitionTime, "transitionTime", "", "", `The duration of the transition from the light’s current state to the new state.
				  This is given as a multiple of 100ms and defaults to 4 (400ms).
				  For example, setting transistiontime:10 will make the transition last 1 second.`)
	cmdLightsState.Flags().StringVarP(&briInc, "briInc", "", "", `Increments or decrements the value of the brightness.
				  briInc is ignored if the bri attribute is provided. Any ongoing bri transition is stopped.
				  Setting a value of 0 also stops any ongoing transition.
				  The bridge will return the bri value after the increment is performed.`)
	cmdLightsState.Flags().StringVarP(&satInc, "satInc", "", "", `Increments or decrements the value of the sat.
				  satInc is ignored if the sat attribute is provided. Any ongoing sat transition is stopped.
				  Setting a value of 0 also stops any ongoing transition.
				  The bridge will return the sat value after the increment is performed.`)
	cmdLightsState.Flags().StringVarP(&hueInc, "hueInc", "", "", ` Increments or decrements the value of the hue.
				  hueInc is ignored if the hue attribute is provided. Any ongoing color transition is stopped.
				  Setting a value of 0 also stops any ongoing transition. The bridge will return the hue value after the increment is performed.
				  Note if the resulting values are < 0 or > 65535 the result is wrapped`)
	cmdLightsState.Flags().StringVarP(&ctInc, "ctInc", "", "", `Increments or decrements the value of the ct.
				  ctInc is ignored if the ct attribute is provided. Any ongoing color transition is stopped.
				  Setting a value of 0 also stops any ongoing transition.
				  The bridge will return the ct value after the increment is performed.`)
	cmdLightsState.Flags().StringVarP(&xyInc, "xyInc", "", "", `Increments or decrements the value of the xy.
				  xyInc is ignored if the xy attribute is provided. Any ongoing color transition is stopped.
				  Setting a value of 0 also stops any ongoing transition. Will stop at it's gamut boundaries.
				  The bridge will return the xy value after the increment is performed.`)
}

// TODO Add Long documentation

var cmdLightsState = &cobra.Command{
	Use:   "state",
	Short: "Set light state: ghue lights state <idLight> [--param1=value], [--param2=value]...",
	Long:  `Set light state: ghue lights state <idLight> [--param1=value], [--param2=value]...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see ./ghue lights state --help")
		} else {
			stateCmd(args[0])
		}
	},
}

func stateCmd(id string) {
	connection := config.ReadConfig()
	setState := &lights.SetStateValues{
		On:             on,
		Alert:          alert,
		Bri:            bri,
		Hue:            hue,
		Sat:            sat,
		XY:             xy,
		Ct:             ct,
		Effect:         effect,
		TransitionTime: transitionTime,
		BriInc:         briInc,
		SatInc:         satInc,
		HueInc:         hueInc,
		CtInc:          ctInc,
		XYInc:          xyInc,
	}
	result, errHUE, err := lights.SetState(connection, id, setState)
	internal.CheckErrors(err, errHUE)

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	internal.Check(err)
	internal.FormatOutputDef(jsonStr)
}
