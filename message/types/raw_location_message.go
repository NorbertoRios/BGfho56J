package types

import (
	"fmt"
	"geometris-go/message/interfaces"
	"strings"
)

//NewRawLocationMessage ...
func NewRawLocationMessage(_serial, _data, _crc string) interfaces.IMessage {
	message := &RawLocationMessage{
		rawData: strings.Split(_data, ","),
		data:    _data,
		crc:     _crc,
	}
	message.identity = fmt.Sprintf("geometris_%v", _serial)
	return message
}

//RawLocationMessage ...
type RawLocationMessage struct {
	Base
	rawData []string
	data    string
	crc     string
}

//CRC ...
func (rlm *RawLocationMessage) CRC() string {
	return rlm.crc
}

//RawByteData ...
func (rlm *RawLocationMessage) RawByteData() []byte {
	return []byte(rlm.data)
}

//RawData ...
func (rlm *RawLocationMessage) RawData() []string {
	return rlm.rawData
}
