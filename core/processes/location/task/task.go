package task

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/immobilizer/task"
	"geometris-go/core/processes/location/states"
)

//New ...
func New(_request interfaces.IRequest) interfaces.ITask {
	task := &Task{}
	task.FacadeRequest = _request
	task.CurrentState = states.NewStartState(_request)
	return task
}

//Task ...
type Task struct {
	task.Task
}
