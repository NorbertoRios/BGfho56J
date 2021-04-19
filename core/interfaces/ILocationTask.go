package interfaces

import (
	"container/list"
	"geometris-go/message/types"
)

//ILocationTask ...
type ILocationTask interface {
	ITask
	NewLocationMessageArrived(*types.RawLocationMessage, IDevice) (*list.List, IProcessResponse)
	NewSyncParam(string, string)
}
