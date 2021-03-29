package usecase

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/usecase"
	"geometris-go/interfaces/unitofwork"
	"geometris-go/logger"
	"geometris-go/message"
	"geometris-go/message/messagetype"
)

//NewMessageIncomeUseCase ...
func NewMessageIncomeUseCase(_rawMessage *message.RawMessage, _device interfaces.IDevice, _uow unitofwork.IUnitOfWork) *MessageIncomeUseCase {
	return &MessageIncomeUseCase{
		rawMessage: _rawMessage,
		device:     _device,
		uow:        _uow,
	}
}

//MessageIncomeUseCase ...
type MessageIncomeUseCase struct {
	device     interfaces.IDevice
	rawMessage *message.RawMessage
	uow        unitofwork.IUnitOfWork
}

//Launch ..
func (miu *MessageIncomeUseCase) Launch() {
	switch miu.rawMessage.MessageType {
	case messagetype.BinaryLocation:
		{
			parser := miu.device.Parser()
			if parser == nil {
				logger.Logger().WriteToLog(logger.Info, "[MessageIncomeUseCase | Launch] Cant process location message. Device parser is nil.")
				return
			}
			locationMessage := parser.Parse(miu.rawMessage).(*message.LocationMessage)
			for _, message := range locationMessage.Messages {
				usecase.NewMessageArrivedUseCase(miu.device, message).Launch()
				miu.uow.UpdateActivity(miu.rawMessage.Identity(), miu.device)
				miu.uow.UpdateState(miu.rawMessage.Identity(), miu.device)
			}
		}
	default:
		{
			msg := miu.prepareMessage()
			usecase.NewMessageArrivedUseCase(miu.device, msg).Launch()
			miu.uow.UpdateActivity(miu.rawMessage.Identity(), miu.device)
		}
	}
	processes := miu.device.DeviceProcesses()
	commands := processes.Update(miu.device)
	usecase.NewBaseUseCase(miu.device, commands).Launch()
	//commit
}

func (miu *MessageIncomeUseCase) prepareMessage() interface{} {
	prepare := NewPrepareMessage(miu.rawMessage)
	return prepare.PreparedMessage()
}
