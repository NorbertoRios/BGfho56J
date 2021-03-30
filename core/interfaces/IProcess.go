package interfaces

import (
	"container/list"
	"geometris-go/message/interfaces"
)

//IProcess ...
type IProcess interface {
	Start(IDevice)
	NewRequest(interface{}, IDevice) IProcessResponse
	MessageArrived(interfaces.IMessage, IDevice) IProcessResponse
	Stop(IDevice, string)
	Pause()
	TasksHistory() *list.List
	Current() ITask
}
