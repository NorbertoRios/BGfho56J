package interfaces

import (
	"container/list"
	"geometris-go/message/types"
)

//ILocationInProgressState ...
type ILocationInProgressState interface {
	ITaskState
	NewLocationMessageArrived(*types.RawLocationMessage, IDevice) (*list.List, IProcessResponse)
	NewSynchParameter(string, string)
}
