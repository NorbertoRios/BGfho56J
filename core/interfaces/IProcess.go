package interfaces

import (
	"container/list"
	"context"
	"geometris-go/message/interfaces"
)

//IProcess ...
type IProcess interface {
	Start(IDevice)
	NewRequest(interface{}, IDevice) IProcessResponse
	MessageArrived(interfaces.IMessage, IDevice) IProcessResponse
	Stop(IDevice, string)
	Pause(context.Context)
	TasksHistory() *list.List
	Current() ITask
}
