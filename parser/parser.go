package parser

import (
	"geometris-go/configuration"
	"geometris-go/core/sensors"
	"geometris-go/core/sensors/builder"
	"geometris-go/logger"
	message "geometris-go/message/interfaces"
	messageTypes "geometris-go/message/types"
	"geometris-go/types"
	"strings"
)

var parser *Parser

//NewWithDiffConfig ...
func NewWithDiffConfig(_dir, _file string) *Parser {
	if parser == nil {
		parser = &Parser{
			reportConfig: configuration.ReportConfig(types.NewFileWithDir(_dir, _file)),
		}
	}
	return parser
}

//New ...
func New() *Parser {
	if parser == nil {
		parser = &Parser{
			reportConfig: configuration.ReportConfig(types.NewFile("/config/initialize/ReportConfiguration.xml")),
		}
	}
	return parser
}

//Parser ...
type Parser struct {
	reportConfig *configuration.ReportConfiguration
}

//Parse ...
func (p *Parser) Parse(_rawMessage message.IMessage, _param12 string) message.IMessage {
	_param12 = strings.Trim(_param12, ";")
	if _param12 == "" {
		logger.Logger().WriteToLog(logger.Error, "[Parser | Parse] Value of 12th parameter is empty for message: "+"_rawMessage"+". Identity: "+_rawMessage.Identity())
		return _rawMessage
	}
	rawMessage, s := _rawMessage.(*messageTypes.RawLocationMessage)
	if !s {
		logger.Logger().WriteToLog(logger.Error, "[Parser | Parse] Unexpected type for message: "+"_rawMessage"+". Identity: ", _rawMessage.Identity())
		return _rawMessage
	}
	sb := builder.NewSensorBuilder()
	keyValue := strings.Split(_param12, "=")
	fields := p.reportConfig.GetFieldsByIds(strings.Split(keyValue[1], "."))
	messageSensors := []sensors.ISensor{}
	values := rawMessage.RawData()
	for i, field := range fields {
		value := values[i]
		messageSensors = append(messageSensors, sb.Build(field.Name, value, field.ValueType)...)
	}
	return messageTypes.NewLocationMessage(_rawMessage.Identity(), messageSensors)
}
