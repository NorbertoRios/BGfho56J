package builder

import (
	"geometris-go/logger"
	"geometris-go/message/interfaces"
	"geometris-go/message/types"
	"regexp"
)

//NewLocationMessageBuilder ...
func NewLocationMessageBuilder(_pattern *regexp.Regexp) interfaces.IBuilder {
	return &LocationMessageBuilder{
		Base: Base{
			pattern: _pattern,
		},
	}
}

//LocationMessageBuilder ...
type LocationMessageBuilder struct {
	Base
}

//Build ...
func (builder *LocationMessageBuilder) Build(_message string) interfaces.IMessage {
	serial := builder.extractKeys("Serial", _message)
	rawData := builder.extractKeys("Data", _message)
	crc := builder.extractKeys("Crc", _message)
	if serial == "" || rawData == "" {
		logger.Logger().WriteToLog(logger.Error, "[LocationMessageBuilder | Build] \"Serial\" or \"RawData\" values are empty. Message : "+_message)
		return nil
	}
	return types.NewRawLocationMessage(serial, rawData, crc)
}
