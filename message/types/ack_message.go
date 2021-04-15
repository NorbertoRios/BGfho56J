package types

import (
	"fmt"
	"geometris-go/message/interfaces"
	"strings"
)

//NewAckMessage ...
func NewAckMessage(_serial, _command string) interfaces.IMessage {
	message := &AckMessage{commands: _command}
	message.identity = fmt.Sprintf("geometris_%v", _serial)
	return message
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
	commands := strings.ReplaceAll(m.commands, "SETPARAMS", "")
	commands = strings.ReplaceAll(commands, "ACK", "")
	commands = strings.Trim(commands, " ")
	rawParams := strings.Split(commands, ";")
	parameters := make(map[string]string)
	for _, param := range rawParams {
		if param == "" {
			continue
		}
		keyValue := strings.Split(param, "=")
		parameters[keyValue[0]] = param
	}
	return parameters
}
