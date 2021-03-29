package request

import (
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
	"strings"
)

//NewState ...
func NewState(_request interfaces.IImmoRequest) *State {
	return &State{
		request: _request,
	}
}

//State ...
type State struct {
	request interfaces.IImmoRequest
}

//String ...
func (s *State) String() string {
	actionPattern := map[string]string{
		"mobilehigh": "OFF",
		"mobilelow":  "ON",
		"armedhigh":  "ON",
		"armedlow":   "OFF",
	}
	key := strings.TrimSpace(s.request.State()) + strings.TrimSpace(s.request.Trigger())
	if action, f := actionPattern[key]; f {
		return action
	}
	logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[State | String] Unexpected action value. Incoming state:%v. Incoming trigger:%v", s.request.State(), s.request.Trigger()))
	return ""
}

//Int ...
func (s *State) Int() int {
	switch strings.ToUpper(s.String()) {
	case "ON":
		return 1
	default:
		return 0
	}
}
