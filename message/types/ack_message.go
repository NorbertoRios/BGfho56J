package types

import (
	"fmt"
	"geometris-go/message/interfaces"
)

//NewAckMessage ...
func NewAckMessage(_serial, _command string) interfaces.IMessage {
	return &AckMessage{
		Base: Base{
			identity: fmt.Sprintf("geometris_%v", _serial),
		},
		Command: _command,
	}
}

//AckMessage represents ack message
type AckMessage struct {
	Base
	Command string
}
