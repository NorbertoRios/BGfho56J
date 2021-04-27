package worker

import (
	"geometris-go/core/usecase"
	"geometris-go/parser/interfaces"
	"geometris-go/repository"
	"geometris-go/storage"
	"sync"
	"time"
)

//NewWorker ...
func NewWorker(_mysql, _rabbit repository.IRepository, _garbageDuration int, _parser interfaces.IParser) *Worker {
	return &Worker{
		garbageDuration: _garbageDuration,
		mysql:           _mysql,
		rabbit:          _rabbit,
		messageChannel:  make(chan *EntryData, 1000000),
		Devices:         make(map[string]bool),
		Mutex:           &sync.Mutex{},
		parser:          _parser,
	}
}

//Worker ...
type Worker struct {
	garbageDuration int
	Mutex           *sync.Mutex
	mysql           repository.IRepository
	rabbit          repository.IRepository
	messageChannel  chan *EntryData
	Devices         map[string]bool
	parser          interfaces.IParser
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
	ticker := time.NewTicker(300 * time.Second)
	for {
		select {
		case entryData := <-w.messageChannel:
			{
				usecase := usecase.NewUDPMessageUseCase(w.mysql, w.rabbit, w.parser)
				usecase.Launch(entryData.Message, entryData.Channel)
			}
		case <-ticker.C:
			{
				w.garbageDevices()
			}
		}
	}
}

func (w *Worker) garbageDevices() {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	for identity := range w.Devices {
		if time.Now().UTC().Sub(storage.Storage().Device(identity).Channel().LastActivity()).Seconds() >= float64(w.garbageDuration) {
			storage.Storage().RemoveDevice(identity)
			delete(w.Devices, identity)
		}
	}
}
