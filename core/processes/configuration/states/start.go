package states

import (
	"container/list"
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
)

//NewStartState ...
func NewStartState(_request interfaces.IConfigRequest) interfaces.ITaskState {
	return &Start{
		request: _request,
	}
}

//Start ...
type Start struct {
	states.Base
	request interfaces.IConfigRequest
}

//Start ....
func (s *Start) Start(_task interfaces.ITask) *list.List {
	cfgTask, _ := _task.(interfaces.IConfigTask)
	cList := list.New()
	command := fmt.Sprintf("SETPARAMS %v ACK", cfgTask.CommandsManager().Command())
	cList.PushBack(commands.NewSendMessageCommand(command))
	_task.ChangeState(NewInProgressState(command, 300, _task))
	return cList
}
