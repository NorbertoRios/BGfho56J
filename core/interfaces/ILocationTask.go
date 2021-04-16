package interfaces

import (
	"container/list"
	"geometris-go/message/types"
)

//ILocationTask ...
type ILocationTask interface {
	ITask
	NewLocationMessageArrived(*types.RawLocationMessage, IProcessResponse, IDevice) *list.List
	NewSyncParam(string, string)
}
