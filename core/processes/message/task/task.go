package task

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/message/states"
	"geometris-go/core/processes/task"
)

//New ...
func New() interfaces.ITask {
	task := &Task{}
	task.CurrentState = states.NewStartState()
	return task
}

//Task ...
type Task struct {
	task.Task
}
