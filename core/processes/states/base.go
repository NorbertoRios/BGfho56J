package states

import (
	"container/list"
	"geometris-go/core/interfaces"
)

//Base ...
type Base struct {
}

//NewMessageArrived ...
func (s *Base) NewMessageArrived(_message interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	return list.New()
}

//Pause ...
func (s *Base) Pause(_task interfaces.ITask) {
	_task.ChangeState(NewPauseState(s))
}

//Resume ...
func (s *Base) Resume(_task interfaces.ITask) {
}

//IsClosed ...
func (s *Base) IsClosed() bool {
	return false
}

//Start ...
func (s *Base) Start(_task interfaces.ITask) *list.List {
	return list.New()
}

//Stop ...
func (s *Base) Stop() *list.List {
	return list.New()
}

//FacadeResponse ...
func (s *Base) FacadeResponse() string {
	return ""
}
