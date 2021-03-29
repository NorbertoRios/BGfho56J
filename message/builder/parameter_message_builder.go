package builder

import (
	"geometris-go/logger"
	"geometris-go/message/interfaces"
	"geometris-go/message/types"
	"regexp"
)

//NewParametersMessageBuilder ...
func NewParametersMessageBuilder(_pattern *regexp.Regexp) interfaces.IBuilder {
	return &ParametersMessageBuilder{
		Base: Base{
			pattern: _pattern,
		},
	}
}

//ParametersMessageBuilder ...
type ParametersMessageBuilder struct {
	Base
}

//Build ...
func (builder *ParametersMessageBuilder) Build(_message string) interfaces.IMessage {
	serial := builder.extractKeys("Serial", _message)
	parameters := builder.extractKeys("Parameters", _message)
	if serial == "" || parameters == "" {
		logger.Logger().WriteToLog(logger.Error, "[ParametersMessageBuilder | Build] \"Serial\" or \"Parameters\" values are empty. Message : "+_message)
		return nil
	}
	return types.NewParametersMessage(serial, parameters)
}
