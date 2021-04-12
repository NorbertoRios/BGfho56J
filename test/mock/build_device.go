package mock

import (
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/interfaces"
	"geometris-go/core/sensors"
	message "geometris-go/message/interfaces"
	"geometris-go/storage"
	unitofwork "geometris-go/unitofwork/interfaces"
)

//NewDeviceBuilder ...
func NewDeviceBuilder(_message message.IMessage, _channel connInterfaces.IChannel, _syncParam string, _unitOfWork unitofwork.IUnitOfWork) *BuildDevice {
	return &BuildDevice{
		message:    _message,
		channel:    _channel,
		syncParam:  _syncParam,
		unitOfWork: _unitOfWork,
	}
}

//BuildDevice ...
type BuildDevice struct {
	message    message.IMessage
	channel    connInterfaces.IChannel
	syncParam  string
	unitOfWork unitofwork.IUnitOfWork
}

//Build ...
func (b *BuildDevice) Build() interfaces.IDevice {
	dev := NewDevice(b.message.Identity(), b.syncParam, []sensors.ISensor{}, b.channel)
	storage.Storage().AddDevice(dev)
	processes := dev.Processes().All()
	for _, process := range processes {
		_response := process.NewRequest(b.message, dev)
		b.unitOfWork.AddDirtyStates(_response.States()...)
		b.unitOfWork.AddDirtyTasks(_response.DirtyTasks()...)
		b.unitOfWork.AddNewTasks(_response.NewTasks()...)
	}
	return dev
}
