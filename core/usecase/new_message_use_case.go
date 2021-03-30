package usecase

import (
	core "geometris-go/core/interfaces"
	"geometris-go/message/factory"
	message "geometris-go/message/interfaces"
	"geometris-go/repository"
	"geometris-go/storage"
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
func (usecase *UDPMessageUseCase) Launch(rawData []byte) {
	messageFactory := factory.New()
	message := messageFactory.BuildMessage(rawData)
	device := storage.Storage().Device(message.Identity())
	if device == nil {

	}
	processes := device.Processes().All()
	for _, p := range processes {
		processResp := p.MessageArrived(message, device)
		usecase.flushProcessResults(processResp)
	}
	usecase.unitOfWork.Commit()
}

func (usecase *UDPMessageUseCase) createDevice(_message message.IMessage) core.IDevice {
	deviceActivity := usecase.mysqlRepository.Load(_message.Identity())
	//Convert activity to sensors
	processes := device.Processes().All()
	for _, p := range processes {
		processResp := p.NewRequest(_message, device)
		usecase.flushProcessResults(processResp)
	}
	return device
}
