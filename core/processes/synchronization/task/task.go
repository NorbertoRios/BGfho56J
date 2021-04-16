package task

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/immobilizer/task"
	"geometris-go/core/processes/synchronization/states"
)

//New ...
func New(_crc string) interfaces.ITask {
	task := &Task{crc: _crc}
	task.CurrentState = states.NewStartState(_crc)
	return task
}

//Task ...
type Task struct {
	task.Task
	crc string
}
