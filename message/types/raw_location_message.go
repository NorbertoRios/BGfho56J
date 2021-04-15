package types

import (
	"fmt"
	"geometris-go/message/interfaces"
	"strings"
)

//NewRawLocationMessage ...
func NewRawLocationMessage(_serial, _data string) interfaces.IMessage {
	message := &RawLocationMessage{
		rawData: strings.Split(_data, ","),
		data:    _data,
	}
	message.identity = fmt.Sprintf("geometris_%v", _serial)
	return message
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
