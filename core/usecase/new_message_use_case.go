package usecase

import (
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/device"
	"geometris-go/core/sensors"
	"geometris-go/message/interfaces"
	"geometris-go/repository"
	"geometris-go/storage"
	"geometris-go/unitofwork"
)

//NewUDPMessageUseCase ...
func NewUDPMessageUseCase(_mysql, _rabbit repository.IRepository) *UDPMessageUseCase {
	usecase := &UDPMessageUseCase{}
	usecase.mysqlRepository = _mysql
	usecase.rabbitRepository = _rabbit
	return usecase
}

//UDPMessageUseCase ...
type UDPMessageUseCase struct {
	Base
}

//Launch ...
func (usecase *UDPMessageUseCase) Launch(message interfaces.IMessage, _channel connInterfaces.IChannel) {
	dev := storage.Storage().Device(message.Identity())
	if dev == nil {
		dev = device.NewDevice(message.Identity(), "", make(map[string]sensors.ISensor), _channel)
	}
	dev.NewChannel(_channel)
	uow := unitofwork.New(usecase.mysqlRepository, usecase.rabbitRepository)
	processes := dev.Processes().All()
	for _, p := range processes {
		processResp := p.MessageArrived(message, dev)
		usecase.flushProcessResults(processResp, uow)
	}
	uow.Commit()
}
