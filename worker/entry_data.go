package worker

import (
	"geometris-go/connection/interfaces"
	"geometris-go/message"
)

//EntryData ...
type EntryData struct {
	RawMessage *message.RawMessage
	Channel    interfaces.IChannel
}
