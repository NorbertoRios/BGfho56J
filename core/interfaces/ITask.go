package interfaces

import (
	"container/list"
	"geometris-go/message/interfaces"
)

//ITask ...
type ITask interface {
	Start() *list.List
	NewMessageArrived(interfaces.IMessage, IDevice) *list.List
	Stop(string)
	Actual()
	IsClosed() bool
	Request() IRequest
	ChangeState(ITaskState)
	FacadeResponse() string
	Pause()
	Resume()
	State() interface{}
}
