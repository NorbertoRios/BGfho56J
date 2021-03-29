package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/relaydrive"
	"geometris-go/core/processes/states"
)

//NewStartState ...
func NewStartState(_request interfaces.IImmoRequest) interfaces.ITaskState {
	return &Start{
		request: _request,
	}
}

//Start ...
type Start struct {
	states.Base
	request interfaces.IImmoRequest
}

//Start ...
func (s *Start) Start(_task interfaces.ITask) *list.List {
	cList := list.New()
	relay := relaydrive.NewSetRelayDrive(s.request)
	cList.PushBack(commands.NewSendMessageCommand(relay.String()))
	inProgress := NewInProgressState(relay.String(), 300, _task)
	inProgress.Run()
	_task.ChangeState(inProgress)
	return cList
}
