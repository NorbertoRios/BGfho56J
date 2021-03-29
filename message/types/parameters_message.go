package types

import (
	"fmt"
	"geometris-go/message/interfaces"
	"strings"
)

//NewParametersMessage ...
func NewParametersMessage(_serial, _parameters string) interfaces.IMessage {
	rawParams := strings.Split(_parameters, ";")
	parameters := make(map[string]string)
	for _, param := range rawParams {
		keyValue := strings.Split(param, "=")
		parameters[keyValue[0]] = keyValue[1]
	}
	return &ParametersMessage{
		Base: Base{
			identity: fmt.Sprintf("geometris_%v", _serial),
		},
		Parameters: parameters,
	}
}

//ParametersMessage represent parameters message
type ParametersMessage struct {
	Base
	Parameters map[string]string
}
