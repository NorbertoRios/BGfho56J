package usecase

import (
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/device"
	core "geometris-go/core/interfaces"
	"geometris-go/message/interfaces"
	iParser "geometris-go/parser/interfaces"
	"geometris-go/rabbitlogger"
	"geometris-go/repository"
	"geometris-go/storage"
	"geometris-go/unitofwork"
	uowInterfaces "geometris-go/unitofwork/interfaces"
)

//NewUDPMessageUseCase ...
func NewUDPMessageUseCase(_mysql, _rabbit repository.IRepository, _parser iParser.IParser) *UDPMessageUseCase {
	usecase := &UDPMessageUseCase{parser: _parser}
	usecase.mysqlRepository = _mysql
	usecase.rabbitRepository = _rabbit
	return usecase
}

//UDPMessageUseCase ...
type UDPMessageUseCase struct {
	Base
	parser iParser.IParser
}

//Launch ...
func (usecase *UDPMessageUseCase) Launch(_message interfaces.IMessage, _channel connInterfaces.IChannel) {
	rabbitlogger.Logger().Log(_message.Content(), _message.Identity())
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
	activity := usecase.mysqlRepository.Load(_message.Identity())
	dev := device.NewDevice(_message.Identity(), activity.Software.SyncParams, activity.LastMessageID, activity.State(usecase.parser.ReportConfig()), _channel, usecase.parser)
	storage.Storage().AddDevice(dev)
	processes := dev.Processes().All()
	for _, process := range processes {
		resp := process.NewRequest(_message, dev)
		usecase.flushProcessResults(resp, _uow)
	}
	return dev
}
