package worker

import (
	"geometris-go/connection/interfaces"
	"geometris-go/message"
)

//IWorkerPool ...
type IWorkerPool interface {
	Flush(*message.RawMessage, interfaces.IChannel)
	Run()
}
