package types

import (
	"fmt"
	"geometris-go/message/interfaces"
	"strings"
)

//NewAckMessage ...
func NewAckMessage(_serial, _command string) interfaces.IMessage {
	return &AckMessage{
		Base: Base{
			identity: fmt.Sprintf("geometris_%v", _serial),
		},
		commands: _command,
	}
}

//AckMessage represents ack message
type AckMessage struct {
	Base
	commands string
}

//Commands ...
func (m *AckMessage) Commands() string {
	return m.commands
}

//Parameters ...
func (m *AckMessage) Parameters() map[string]string {
	rawParams := strings.Split(m.commands, ";")
	parameters := make(map[string]string)
	for _, param := range rawParams {
		keyValue := strings.Split(param, "=")
		parameters[keyValue[0]] = keyValue[1]
	}
	return parameters
}
