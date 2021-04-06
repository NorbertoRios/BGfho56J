package task

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/configuration/states"
	"geometris-go/core/processes/task"
	message "geometris-go/message/interfaces"
)

//New ...
func New(_facadeRequest interfaces.IConfigRequest) *Task {
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

//Command ...
func (t *Task) Command() string {
	return t.commandsManager.Command()
}

//Start ...
func (t *Task) Start() *list.List {
	return t.CurrentState.Start(t)
}

//Pause ...
func (t *Task) Pause() {
	t.CurrentState.Pause(t)
}

//NewMessageArrived ...
func (t *Task) NewMessageArrived(_message message.IMessage, _device interfaces.IDevice) *list.List {
	return t.CurrentState.NewMessageArrived(_message, _device, t)
}

//Resume ...
func (t *Task) Resume() {
	t.CurrentState.Resume(t)
}
