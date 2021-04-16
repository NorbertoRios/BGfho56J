package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
)

//NewStartState ...
func NewStartState(_crc string) interfaces.ITaskState {
	return &Start{
		crc: _crc,
	}
}

//Start ...
type Start struct {
	states.Base
	crc string
}

//Start ...
func (s *Start) Start(_task interfaces.ITask) *list.List {
	cList := list.New()
	cList.PushBack(commands.NewSendMessageCommand("DIAG PARAM=12 VIAUDP"))
	inProgress := NewInProgressState(s.crc, "DIAG PARAM=12 VIAUDP", 30, _task)
	inProgress.Run()
	_task.ChangeState(inProgress)
	return cList
}
