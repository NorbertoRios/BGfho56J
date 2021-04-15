package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
)

//NewStartState ...
func NewStartState(_request interfaces.IBOCommandRequest) interfaces.ITaskState {
	return &Start{
		request: _request,
	}
}

//Start ...
type Start struct {
	states.Base
	request interfaces.IBOCommandRequest
}

//Start ...
func (s *Start) Start(_task interfaces.ITask) *list.List {
	cList := list.New()
	cList.PushBack(commands.NewSendMessageCommand(s.request.Command()))
	_task.ChangeState(states.NewDone(s.request))
	return cList
}
