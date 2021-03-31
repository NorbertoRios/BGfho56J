package usecase

import (
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/message/factory"
	"geometris-go/repository"
	"geometris-go/storage"
	"geometris-go/unitofwork"
)

//NewUDPMessageUseCase ...
func NewUDPMessageUseCase(_mysql, _rabbit repository.IRepository) *UDPMessageUseCase {
	return &UDPMessageUseCase{
		Base: Base{
			mysqlRepository:  _mysql,
			rabbitRepository: _rabbit,
		},
	}
}

//UDPMessageUseCase ...
type UDPMessageUseCase struct {
	Base
}

//Launch ...
func (usecase *UDPMessageUseCase) Launch(rawData []byte, _channel connInterfaces.IChannel) {
	messageFactory := factory.New()
	message := messageFactory.BuildMessage(rawData)
	device := storage.Storage().Device(message.Identity())
	// if device == nil {

	// }
	uow := unitofwork.New(usecase.mysqlRepository, usecase.rabbitRepository)
	processes := device.Processes().All()
	for _, p := range processes {
		processResp := p.MessageArrived(message, device)
		usecase.flushProcessResults(processResp, uow)
	}
	uow.Commit()
}

//func (usecase *UDPMessageUseCase) createDevice(_message message.IMessage) core.IDevice {
//	deviceActivity := usecase.mysqlRepository.Load(_message.Identity()).(*models.DeviceActivity)
//	//Convert activity to sensors
//	processes := device.Processes().All()
//	for _, p := range processes {
//		processResp := p.NewRequest(_message, device)
//		usecase.flushProcessResults(processResp)
//	}
//	return device
//}
