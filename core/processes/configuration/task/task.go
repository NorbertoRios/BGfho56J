package task

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/configuration/states"
	"geometris-go/core/processes/task"
)

//New ...
func New(_facadeRequest interfaces.IConfigRequest) interfaces.IConfigTask {
	task := &Task{
		commandsManager: NewCommandsManager(_facadeRequest.Commands()),
	}
	task.FacadeRequest = _facadeRequest
	task.CurrentState = states.NewStartState(_facadeRequest)
	return task
}

//Task ...
type Task struct {
	task.Task
	commandsManager interfaces.ICommandsManager
}

//CommandsManager ...
func (t *Task) CommandsManager() interfaces.ICommandsManager {
	return t.commandsManager
}
