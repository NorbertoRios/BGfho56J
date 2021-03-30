package worker

import (
	"geometris-go/connection/interfaces"
	"geometris-go/interfaces/unitofwork"
	message "geometris-go/message/interfaces"
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
func (wp *WorkersPool) Flush(message message.IMessage, channel interfaces.IChannel) {
	data := &EntryData{RawMessage: message, Channel: channel}
	for _, worker := range wp.Pool.all() {
		if worker.DeviceExist(message.Identity()) {
			worker.Push(data)
			return
		}
	}
	worker := wp.Pool.next()
	worker.NewDevice(message.Identity())
	worker.Push(data)
}

//Run ...
func (wp *WorkersPool) Run() {
	for _, worker := range wp.Pool.Workers {
		go worker.Run()
	}
}
