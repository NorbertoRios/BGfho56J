package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/logger"
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
	logger.Logger().WriteToLog(logger.Info, "[MessageProcess | StartState] Message process is started")
	inProgress := NewInProgressState()
	_task.ChangeState(inProgress)
	return cList
}
