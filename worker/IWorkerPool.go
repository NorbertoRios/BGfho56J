package worker

import (
	"geometris-go/connection/interfaces"
	message "geometris-go/message/interfaces"
)

//IWorkerPool ...
type IWorkerPool interface {
	Flush(message.IMessage, interfaces.IChannel)
	Run()
}
