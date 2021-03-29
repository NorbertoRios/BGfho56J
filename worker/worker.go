package worker

import (
	"geometris-go/interfaces/unitofwork"
	"geometris-go/worker/usecase"
	"sync"
)

//NewWorker ...
func NewWorker(_uow unitofwork.IUnitOfWork) *Worker {
	return &Worker{
		Uow:            _uow,
		messageChannel: make(chan *EntryData, 1000000),
		Devices:        make(map[string]bool),
		Mutex:          &sync.Mutex{},
	}
}

//Worker ...
type Worker struct {
	Mutex          *sync.Mutex
	Uow            unitofwork.IUnitOfWork
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
				device := w.Uow.Device(entryData.RawMessage.Identity())
				if device == nil {
					w.Uow.Register(entryData.RawMessage.Identity(), entryData.Channel)
					device = w.Uow.Device(entryData.RawMessage.Identity())
				} else {
					device.NewChannel(entryData.Channel)
				}
				usecase.NewMessageIncomeUseCase(entryData.RawMessage, device, w.Uow).Launch()
				// if !w.uow.Commit() {
				// 	logger.Logger().WriteToLog(logger.Fatal, "[Worker | Run] Something went wrong while commit changes to database")
				// }
			}
		}
	}
}
