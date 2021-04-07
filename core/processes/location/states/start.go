package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
)

//NewStartState ...
func NewStartState(_request interfaces.IRequest) interfaces.ITaskState {
	return &Start{
		request: _request,
	}
}

//Start ...
type Start struct {
	states.Base
	request interfaces.IRequest
}

//Start ...
func (s *Start) Start(_task interfaces.ITask) *list.List {
	cList := list.New()
	cList.PushBack(commands.NewSendMessageCommand("POLLQ VIAUDP"))
	inProgress := NewInProgressState("POLLQ VIAUDP", 60, _task)
	inProgress.Run()
	_task.ChangeState(inProgress)
	return cList
}
