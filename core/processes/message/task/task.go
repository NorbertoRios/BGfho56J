package task

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/message/states"
	"geometris-go/core/processes/task"
	"geometris-go/message/types"
)

//New ...
func New(_synchParams map[string]string) interfaces.ILocationTask {
	task := &Task{}
	task.CurrentState = states.NewInProgressState(_synchParams)
	return task
}

//Task ...
type Task struct {
	task.Task
}

//NewLocationMessageArrived ...
func (t *Task) NewLocationMessageArrived(_message *types.RawLocationMessage, _device interfaces.IDevice) (*list.List, interfaces.IProcessResponse) {
	return t.CurrentState.(interfaces.ILocationInProgressState).NewLocationMessageArrived(_message, _device)
}

//NewSyncParam ...
func (t *Task) NewSyncParam(crc, syncParam string) {
	t.CurrentState.(interfaces.ILocationInProgressState).NewSynchParameter(crc, syncParam)
}
