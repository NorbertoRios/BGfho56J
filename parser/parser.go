package parser

import (
	"geometris-go/configuration"
	"geometris-go/core/sensors"
	"geometris-go/core/sensors/builder"
	message "geometris-go/message/interfaces"
	messageTypes "geometris-go/message/types"
	"geometris-go/types"
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
			reportConfig: configuration.ReportConfig(types.NewFile("/config/initializer/ReportConfiguration.xml")),
		}
	}
	return parser
}

//Parser ...
type Parser struct {
	reportConfig *configuration.ReportConfiguration
}

//Parse ...
func (p *Parser) Parse(_rawMessage *messageTypes.RawLocationMessage, _synchParam *SynchParameter) message.IMessage {
	sb := builder.NewSensorBuilder()
	fields := p.reportConfig.GetFieldsByIds(_synchParam.ColumnsID())
	messageSensors := []sensors.ISensor{}
	values := _rawMessage.RawData()
	for i, value := range values {
		field := fields[i]
		messageSensors = append(messageSensors, sb.Build(field.Name, value, field.ValueType)...)
	}
	return messageTypes.NewLocationMessage(_rawMessage.Identity(), messageSensors)
}
