package interfaces

import (
	"container/list"
)

//ITaskState ...
type ITaskState interface {
	NewMessageArrived(interface{}, IDevice, ITask) *list.List
	IsClosed() bool
	Start(ITask) *list.List
	Stop() *list.List
	Resume(ITask)
	Pause(ITask)
	FacadeResponse() string
}
