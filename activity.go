package pushover

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func init() {
	_ = activity.Register(&Activity{}, New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

type Activity struct {
	appToken   string
	groupToken string
	active     bool
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	logger := ctx.Logger()
	logger.Debugf("Setting: %b", s)

	act := &Activity{
		appToken:   s.AppToken,
		groupToken: s.GroupToken,
		active:     s.Active,
	}

	return act, nil
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	output := &Output{}
	if !a.active || len(input.Message) == 0 {
		output.Status = 204
	} else {
		//do something
		success, err := a.sendPushover(input.Message)
		if err != nil {
			return false, err
		}

		if !success {
			output.Status = 400
		} else {
			output.Status = 200
		}
	}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}

type pushover struct {
	Token   string `json:"token"`
	User    string `json:"user"`
	Message string `json:"message"`
}

func (a *Activity) sendPushover(message string) (bool, error) {
	url := "https://api.pushover.net/1/messages.json"
	json, err := json.Marshal(pushover{Token: a.appToken, User: a.groupToken, Message: message})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		return false, err
	}
	return true, nil
}
