package factory

import (
	"fmt"
	"geometris-go/logger"
	"geometris-go/message/builder"
	"geometris-go/message/interfaces"
	"regexp"
)

//New returns raw message factory
func New() *MessageFactory {
	location, _ := regexp.Compile("^(?P<Serial>[a-zA-Z0-9]{12}),(?P<Data>(.*))$")
	ack, _ := regexp.Compile("^(?P<Serial>[a-zA-Z0-9]{12}) ACK <(?P<Command>(.*))>")
	diagParameters, _ := regexp.Compile("^(?P<Serial>[a-zA-Z0-9]{12}) PARAMETERS (?P<Parameters>(.*));$")
	factory := &MessageFactory{
		builders: []interfaces.IBuilder{
			builder.NewAckMessageBuilder(ack),
			builder.NewParametersMessageBuilder(diagParameters),
			builder.NewLocationMessageBuilder(location),
		},
	}
	return factory
}

//MessageFactory factory for raw message
type MessageFactory struct {
	builders []interfaces.IBuilder
}

//BuildMessage ....
func (factory *MessageFactory) BuildMessage(_packet []byte) interfaces.IMessage {
	message := string(_packet)
	logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[MessageFactory | BuildMessage] New message: String:%v", message))
	for _, builder := range factory.builders {
		if builder.IsParsable(_packet) {
			return builder.Build(message)
		}
	}
	logger.Logger().WriteToLog(logger.Error, fmt.Sprintf("[MessageFactory | BuildMessage] Builder not found for message: \n\tByte:%v\n\tString:%v", _packet, message))
	return nil
}
