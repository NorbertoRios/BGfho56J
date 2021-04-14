package worker

import (
	"fmt"
	"geometris-go/connection/interfaces"
	"geometris-go/logger"
	"geometris-go/message/factory"
	"geometris-go/repository"
	"sync"
)

//NewWorkerPool ...
func NewWorkerPool(workersCount int, _mysql, _rabbit repository.IRepository) IWorkerPool {
	return &WorkersPool{
		Pool:  NewPool(workersCount, _mysql, _rabbit),
		mutex: &sync.Mutex{},
	}
}

//WorkersPool ...
type WorkersPool struct {
	Pool  *Pool
	mutex *sync.Mutex
}

//Flush ...
func (wp *WorkersPool) Flush(rawData []byte, channel interfaces.IChannel) {
	wp.mutex.Lock()
	defer func() {
		if r := recover(); r != nil {
			logger.Logger().WriteToLog(logger.Error, "[WorkersPool | Recovered] Recovered with error: ", r)
		}
	}()
	defer wp.mutex.Unlock()
	messageFactory := factory.New()
	message := messageFactory.BuildMessage(rawData)
	if message == nil {
		return
	}
	data := &EntryData{Message: message, Channel: channel}
	for id, worker := range wp.Pool.all() {
		if worker.DeviceExist(message.Identity()) {
			logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[WorkersPool | Flush] Message %v added to worker: %v", message.Identity(), id))
			worker.Push(data)
			return
		}
	}
	worker := wp.Pool.next()
	worker.NewDevice(message.Identity())
	worker.Push(data)
}
