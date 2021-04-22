package types

import (
	"fmt"
	"geometris-go/message/interfaces"
	"strings"
)

//NewParametersMessage ...
func NewParametersMessage(_serial, _parameters string) interfaces.IMessage {
	return &ParametersMessage{
		Base: Base{
			identity: fmt.Sprintf("geometris_%v", _serial),
		},
		parameters: _parameters,
	}
}

//ParametersMessage represent parameters message
type ParametersMessage struct {
	Base
	parameters string
}

//Content ...
func (m *ParametersMessage) Content() string {
	return m.parameters
}

//Parameters ...
func (m *ParametersMessage) Parameters() map[string]string {
	rawParams := strings.Split(m.parameters, ";")
	parameters := make(map[string]string)
	for _, param := range rawParams {
		keyValue := strings.Split(param, "=")
		parameters[keyValue[0]] = keyValue[1]
	}
	return parameters
}
