package worker

import (
	"geometris-go/connection/interfaces"
)

//IWorkerPool ...
type IWorkerPool interface {
	Flush([]byte, interfaces.IChannel)
	Run()
}
