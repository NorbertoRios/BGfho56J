package task

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	message "geometris-go/message/interfaces"
)

//Task ...
type Task struct {
	CurrentState  interfaces.ITaskState
	FacadeRequest interfaces.IRequest
}

//Start ...
func (t *Task) Start() *list.List {
	return t.CurrentState.Start(t)
}

//Pause ...
func (t *Task) Pause() {

}

//FacadeResponse ...
func (t *Task) FacadeResponse() string {
	return t.CurrentState.FacadeResponse()
}

//NewMessageArrived ...
func (t *Task) NewMessageArrived(_message message.IMessage, _device interfaces.IDevice) *list.List {
	return t.CurrentState.NewMessageArrived(_message, _device, t)
}

//Stop ...
func (t *Task) Stop(_description string) {
	t.ChangeState(states.NewClose(t.FacadeRequest, _description))
}

//Actual ...
func (t *Task) Actual() {
	t.ChangeState(states.NewActual(t.FacadeRequest))
}

//IsClosed ...
func (t *Task) IsClosed() bool {
	return t.CurrentState.IsClosed()
}

//Request ...
func (t *Task) Request() interfaces.IRequest {
	return t.FacadeRequest
}

//ChangeState ...
func (t *Task) ChangeState(state interfaces.ITaskState) {
	t.CurrentState = state
}
