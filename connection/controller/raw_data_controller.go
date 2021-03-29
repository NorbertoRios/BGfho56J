package controller

import (
	"geometris-go/connection/interfaces"
	"geometris-go/message"
	"geometris-go/worker"
)

//NewRawDataController ...
func NewRawDataController(wp *worker.WorkersPool) *RawDataController {
	return &RawDataController{
		factory:     message.Factory(),
		workersPool: wp,
	}
}

//RawDataController ...
type RawDataController struct {
	factory     *message.RawMessageFactory
	workersPool *worker.WorkersPool
}

//Process ..
func (controller *RawDataController) Process(packet []byte, channel interfaces.IChannel) {
	rMessage := controller.factory.BuildRawMessage(packet)
	controller.workersPool.Flush(rMessage, channel)
}
