package parser

import (
	"geometris-go/configuration"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	message "geometris-go/message/interfaces"
	messageTypes "geometris-go/message/types"
	"geometris-go/types"
	"strings"
)

var parser *Parser

//New ...
func New(file *types.File) *Parser {
	if parser == nil {
		provider := configuration.ConstructXMLProvider(file)
		parser = &Parser{
			reportConfig: configuration.ConstructReportConfiguration(provider),
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
	keyValue := strings.Split(_param12, "=")
	fields := p.reportConfig.GetFieldsByIds(strings.Split(keyValue[1], "."))
	messageSensors := []sensors.ISensor{}
	values := rawMessage.RawData()
	for i, field := range fields {
		value := values[i]
		sensor := sensors.NewSensor(field.ID, field.Name, p.castValueToType(value, field.ValueType))
		messageSensors = append(messageSensors, sensor)
	}
	return messageTypes.NewLocationMessage(_rawMessage.Identity(), messageSensors)
}

func (p *Parser) castValueToType(_value, _type string) interface{} {
	strValue := types.NewString(_value)
	switch _type {
	case "float64":
		return strValue.Float(64)
	case "uint32":
		return strValue.UInt(32)
	case "int":
		return strValue.IntValue()
	case "byte":
		return strValue.Byte()
	default:
		return _value
	}
}
