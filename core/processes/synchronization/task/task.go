package task

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/synchronization/states"
	"geometris-go/core/processes/task"
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

//Stop ...
func (t *Task) Stop(_description string) {
	t.ChangeState(states.NewClosed(_description))
}
