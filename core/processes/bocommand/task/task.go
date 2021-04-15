package task

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/bocommand/states"
	"geometris-go/core/processes/task"
)

//New ...
func New(_request interfaces.IBOCommandRequest) interfaces.ITask {
	t := &Task{}
	t.FacadeRequest = _request
	t.CurrentState = states.NewStartState(_request)
	return t
}

//Task ...
type Task struct {
	task.Task
}
