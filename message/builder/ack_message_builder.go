package builder

import (
	"geometris-go/logger"
	"geometris-go/message/interfaces"
	"geometris-go/message/types"
	"regexp"
)

//NewAckMessageBuilder ...
func NewAckMessageBuilder(_pattern *regexp.Regexp) interfaces.IBuilder {
	return &AckMessageBuilder{
		Base: Base{
			pattern: _pattern,
		},
	}
}

//AckMessageBuilder ...
type AckMessageBuilder struct {
	Base
}

//Build ...
func (builder *AckMessageBuilder) Build(_message string) interfaces.IMessage {
	serial := builder.extractKeys("Serial", _message)
	command := builder.extractKeys("Command", _message)
	if serial == "" || command == "" {
		logger.Logger().WriteToLog(logger.Error, "[AckMessageBuilder | Build] \"Serial\" or \"Command\" values are empty. Message : "+_message)
		return nil
	}
	return types.NewAckMessage(serial, command)
}
