package interfaces

import (
	"container/list"
	"geometris-go/message/interfaces"
)

//IProcess ...
type IProcess interface {
	Symbol() string
	Start(IDevice)
	NewRequest(interface{}, IDevice) IProcessResponse
	MessageArrived(interfaces.IMessage, IDevice) IProcessResponse
	Stop(IDevice, string)
	Pause()
	Resume()
	TasksHistory() *list.List
	Current() ITask
}
