package controller

import (
	"geometris-go/connection/interfaces"
	"geometris-go/message/factory"
	"geometris-go/worker"
)

//NewRawDataController ...
func NewRawDataController(wp *worker.WorkersPool) *RawDataController {
	return &RawDataController{
		factory:     factory.New(),
		workersPool: wp,
	}
}

//RawDataController ...
type RawDataController struct {
	factory     *factory.MessageFactory
	workersPool *worker.WorkersPool
}

//Process ..
func (controller *RawDataController) Process(packet []byte, channel interfaces.IChannel) {
	rMessage := controller.factory.BuildMessage(packet)
	controller.workersPool.Flush(rMessage, channel)
}
