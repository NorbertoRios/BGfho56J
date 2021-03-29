package task

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/immobilizer/task"
	"geometris-go/core/processes/synchronization/states"
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
