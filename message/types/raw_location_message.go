package types

import (
	"fmt"
	"geometris-go/message/interfaces"
	"strings"
)

//NewRawLocationMessage ...
func NewRawLocationMessage(_serial, _data string) interfaces.IMessage {
	return &RawLocationMessage{
		Base: Base{
			identity: fmt.Sprintf("geometris_%v", _serial),
		},
		RawData: strings.Split(_data, ","),
	}
}

//RawLocationMessage ...
type RawLocationMessage struct {
	Base
	RawData []string
}
