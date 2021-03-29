package task

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/immobilizer/states"
	"geometris-go/core/processes/task"
)

//New ...
func New(_facadeRequest interfaces.IImmoRequest) interfaces.ITask {
	task := &Task{}
	task.FacadeRequest = _facadeRequest
	task.CurrentState = states.NewStartState(_facadeRequest)
	return task
}

//Task ...
type Task struct {
	task.Task
}
