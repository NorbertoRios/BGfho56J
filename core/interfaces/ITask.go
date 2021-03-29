package interfaces

import (
	"container/list"
	"geometris-go/message"
)

//ITask ...
type ITask interface {
	Start() *list.List
	NewMessageArrived(message.IMessage, IDevice) *list.List
	Stop(string)
	Actual()
	IsClosed() bool
	Request() IRequest
	ChangeState(ITaskState)
	FacadeResponse() string
}
