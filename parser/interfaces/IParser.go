package interfaces

import (
	"geometris-go/configuration"
	message "geometris-go/message/interfaces"
	messageTypes "geometris-go/message/types"
)

//IParser ...
type IParser interface {
	Parse(*messageTypes.RawLocationMessage, ISynchParameter) message.IMessage
	ReportConfig() *configuration.ReportConfiguration
}
