package controller

import (
	"geometris-go/connection/interfaces"
	"geometris-go/worker"
)

//NewRawDataController ...
func NewRawDataController(wp worker.IWorkerPool) interfaces.IController {
	return &RawDataController{
		workersPool: wp,
	}
}

//RawDataController ...
type RawDataController struct {
	workersPool worker.IWorkerPool
}

//Process ..
func (controller *RawDataController) Process(packet []byte, channel interfaces.IChannel) {
	controller.workersPool.Flush(packet, channel)
}
