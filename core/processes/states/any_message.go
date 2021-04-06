package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
)

//NewAnyMessageState ...
func NewAnyMessageState(_message string, _constructor func(string, int, interfaces.ITask) interfaces.IInProgressState) interfaces.ITaskState {
	return &AnyMessageState{
		message:     _message,
		constructor: _constructor,
	}
}

//AnyMessageState ...
type AnyMessageState struct {
	Base
	message     string
	constructor func(string, int, interfaces.ITask) interfaces.IInProgressState
}

//NewMessageArrived ...
func (s *AnyMessageState) NewMessageArrived(_message interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	cList := list.New()
	cList.PushBack(commands.NewSendMessageCommand(s.message))
	inProgress := s.constructor(s.message, 300, _task)
	inProgress.Run()
	_task.ChangeState(inProgress)
	return cList
}
