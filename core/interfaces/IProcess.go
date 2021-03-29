package interfaces

import (
	"container/list"
	"geometris-go/message"
)

//IProcess ...
type IProcess interface {
	Start(IDevice)
	NewRequest(interface{}, IDevice) IProcessResponse
	MessageArrived(message.IMessage, IDevice) IProcessResponse
	Stop(IDevice, string)
	Pause()
	TasksHistory() *list.List
	Current() ITask
}
