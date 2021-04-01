package worker

import (
	"geometris-go/connection/interfaces"
	message "geometris-go/message/interfaces"
)

//EntryData ...
type EntryData struct {
	Message message.IMessage
	Channel interfaces.IChannel
}
