package parser

import (
	"geometris-go/configuration"
	"geometris-go/core/sensors"
	"geometris-go/core/sensors/builder"
	message "geometris-go/message/interfaces"
	messageTypes "geometris-go/message/types"
	"geometris-go/parser/interfaces"
	"geometris-go/types"
)

//NewWithDiffConfig ...
func NewWithDiffConfig(_dir, _file string) interfaces.IParser {
	return &Parser{
		reportConfig: configuration.NewReportConfig(types.NewFileWithDir(_dir, _file)),
	}
}

//New ...
func New() interfaces.IParser {
	return &Parser{
		reportConfig: configuration.NewReportConfig(types.NewFile("/config/initializer/ReportConfiguration.xml")),
	}
}

//Parser ...
type Parser struct {
	reportConfig *configuration.ReportConfiguration
}

//Parse ...
func (p *Parser) Parse(_rawMessage *messageTypes.RawLocationMessage, _synchParam interfaces.ISynchParameter) message.IMessage {
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

//ReportConfig ...
func (p *Parser) ReportConfig() *configuration.ReportConfiguration {
	return p.reportConfig
}
