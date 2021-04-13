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
		rawData: strings.Split(_data, ","),
		data:    _data,
	}
}

//RawLocationMessage ...
type RawLocationMessage struct {
	Base
	rawData []string
	data    string
}

//RawByteData ...
func (rlm *RawLocationMessage) RawByteData() []byte {
	return []byte(rlm.data)
}

//RawData ...
func (rlm *RawLocationMessage) RawData() []string {
	return rlm.rawData
}
