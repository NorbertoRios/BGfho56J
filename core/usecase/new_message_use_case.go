package usecase

import (
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/device"
	core "geometris-go/core/interfaces"
	"geometris-go/core/sensors"
	"geometris-go/message/interfaces"
	"geometris-go/repository"
	"geometris-go/storage"
	"geometris-go/unitofwork"
	uowInterfaces "geometris-go/unitofwork/interfaces"
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
func (usecase *UDPMessageUseCase) Launch(_message interfaces.IMessage, _channel connInterfaces.IChannel) {
	dev := storage.Storage().Device(_message.Identity())
	uow := unitofwork.New(usecase.mysqlRepository, usecase.rabbitRepository)
	if dev == nil {
		dev = usecase.BuildDevice(_message, _channel, uow)
	}
	dev.NewChannel(_channel)
	processes := dev.Processes().All()
	for _, p := range processes {
		if p.Current() == nil {
			continue
		}
		processResp := p.MessageArrived(_message, dev)
		usecase.flushProcessResults(processResp, uow)
	}
	uow.Commit()
}

//BuildDevice ...
func (usecase *UDPMessageUseCase) BuildDevice(_message interfaces.IMessage, _channel connInterfaces.IChannel, _uow uowInterfaces.IUnitOfWork) core.IDevice {
	_syncParam := ""
	dev := device.NewDevice(_message.Identity(), _syncParam, make(map[string]sensors.ISensor), _channel)
	storage.Storage().AddDevice(dev)
	processes := dev.Processes().All()
	for _, process := range processes {
		resp := process.NewRequest(_message, dev)
		usecase.flushProcessResults(resp, _uow)
	}
	return dev
}
