package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
)

//NewStartState ...
func NewStartState() interfaces.ITaskState {
	return &Start{}
}

//Start ...
type Start struct {
	states.Base
}

//Start ...
func (s *Start) Start(_task interfaces.ITask) *list.List {
	cList := list.New()
	cList.PushBack(commands.NewSendMessageCommand("DIAG PARAMS=12 VIAUDP"))
	inProgress := NewInProgressState("DIAG PARAMS=12 VIAUDP", 300, _task)
	_task.ChangeState(inProgress)
	return cList
}
