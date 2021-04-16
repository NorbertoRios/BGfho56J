package interfaces

import (
	"container/list"
	"geometris-go/message/types"
)

//ILocationInProgressState ...
type ILocationInProgressState interface {
	ITaskState
	NewLocationMessageArrived(*types.RawLocationMessage, IProcessResponse, IDevice) *list.List
	NewSynchParameter(string, string)
}
