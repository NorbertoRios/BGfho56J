package worker

import (
	"geometris-go/core/usecase"
	"geometris-go/repository"
	"sync"
)

//NewWorker ...
func NewWorker(_mysql, _rabbit repository.IRepository) *Worker {
	return &Worker{
		mysql:          _mysql,
		rabbit:         _rabbit,
		messageChannel: make(chan *EntryData, 1000000),
		Devices:        make(map[string]bool),
		Mutex:          &sync.Mutex{},
	}
}

//Worker ...
type Worker struct {
	Mutex          *sync.Mutex
	mysql          repository.IRepository
	rabbit         repository.IRepository
	messageChannel chan *EntryData
	Devices        map[string]bool
}

//NewDevice ...
func (w *Worker) NewDevice(identity string) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	w.Devices[identity] = true
}

//DeviceExist ...
func (w *Worker) DeviceExist(identity string) bool {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	return w.Devices[identity]
}

//Push ...
func (w *Worker) Push(data *EntryData) {
	w.messageChannel <- data
}

//Run ...
func (w *Worker) Run() {
	for {
		select {
		case entryData := <-w.messageChannel:
			{
				usecase := usecase.NewUDPMessageUseCase(w.mysql, w.rabbit)
				usecase.Launch(entryData.Message, entryData.Channel)
			}
		}
	}
}
