package worker

import (
	"geometris-go/connection/interfaces"
	"geometris-go/interfaces/unitofwork"
	"geometris-go/message"
)

//NewWorkerPool ...
func NewWorkerPool(workersCount int, _uow unitofwork.IUnitOfWork) IWorkerPool {
	return &WorkersPool{
		Pool: NewPool(workersCount, _uow),
	}
}

//WorkersPool ...
type WorkersPool struct {
	Pool *Pool
}

//Flush ...
func (wp *WorkersPool) Flush(rMessage *message.RawMessage, channel interfaces.IChannel) {
	data := &EntryData{RawMessage: rMessage, Channel: channel}
	for _, worker := range wp.Pool.all() {
		if worker.DeviceExist(rMessage.Identity()) {
			worker.Push(data)
			return
		}
	}
	worker := wp.Pool.next()
	worker.NewDevice(rMessage.Identity())
	worker.Push(data)
}

//Run ...
func (wp *WorkersPool) Run() {
	for _, worker := range wp.Pool.Workers {
		go worker.Run()
	}
}
