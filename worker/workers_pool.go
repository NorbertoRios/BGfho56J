package worker

import (
	"geometris-go/connection/interfaces"
	"geometris-go/message/factory"
	"geometris-go/repository"
)

//NewWorkerPool ...
func NewWorkerPool(workersCount int, _mysql, _rabbit repository.IRepository) IWorkerPool {
	return &WorkersPool{
		Pool: NewPool(workersCount, _mysql, _rabbit),
	}
}

//WorkersPool ...
type WorkersPool struct {
	Pool *Pool
}

//Flush ...
func (wp *WorkersPool) Flush(rawData []byte, channel interfaces.IChannel) {
	messageFactory := factory.New()
	message := messageFactory.BuildMessage(rawData)
	data := &EntryData{Message: message, Channel: channel}
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
