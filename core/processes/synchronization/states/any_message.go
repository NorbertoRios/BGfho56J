package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
)

//NewAnyMessageState ...
func NewAnyMessageState(_crc, _message string) interfaces.ITaskState {
	return &AnyMessageState{
		message: _message,
		crc:     _crc,
	}
}

//AnyMessageState ...
type AnyMessageState struct {
	states.Base
	message string
	crc     string
}

//NewMessageArrived ...
func (s *AnyMessageState) NewMessageArrived(_message interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	cList := list.New()
	cList.PushBack(commands.NewSendMessageCommand(s.message))
	inProgress := NewInProgressState(s.crc, s.message, 30, _task)
	inProgress.Run()
	_task.ChangeState(inProgress)
	return cList
}
