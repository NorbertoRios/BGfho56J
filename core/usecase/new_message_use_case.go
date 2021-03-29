package usecase

import (
	core "geometris-go/core/interfaces"
	"geometris-go/message/factory"
	message "geometris-go/message/interfaces"
	"geometris-go/unitofwork/interfaces"
)

//NewUDPMessageUseCase ...
func NewUDPMessageUseCase(_unitOfWork interfaces.IUnitOfWork) *UDPMessageUseCase {
	usecase := &UDPMessageUseCase{}
	usecase.unitOfWork = _unitOfWork
	return usecase
}

//UDPMessageUseCase ...
type UDPMessageUseCase struct {
	Base
}

//Launch ...
func (usecase *UDPMessageUseCase) Launch(rawData []byte) {
	messageFactory := factory.New()
	message := messageFactory.BuildMessage(rawData)
	device := usecase.unitOfWork.Device(message.Identity())
	if device == nil {
		device = usecase.createDevice(message)
	}
	processes := device.Processes().All()
	for _, p := range processes {
		processResp := p.MessageArrived(message, device)
		usecase.flushProcessResults(processResp)
	}
	usecase.unitOfWork.Commit()
}

func (usecase *UDPMessageUseCase) createDevice(_message message.IMessage) core.IDevice {
	device := usecase.unitOfWork.NewDevice(_message.Identity())
	processes := device.Processes().All()
	for _, p := range processes {
		processResp := p.NewRequest(_message, device)
		usecase.flushProcessResults(processResp)
	}
	return device
}
