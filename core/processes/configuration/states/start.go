package states

import (
	"container/list"
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
	"geometris-go/logger"
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
	command := cfgTask.Command()
	if command == "" {
		logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[Configuration | Start] Task %v unexpected closed. No commands to send.", command))
		_task.ChangeState(states.NewClose(s.request, "Commands is empty"))
		return cList
	}
	cList.PushBack(commands.NewSendMessageCommand("SETPARAMS " + command + " ACK;"))
	inProgress := NewInProgressState(command, 60, _task)
	_task.ChangeState(inProgress)
	inProgress.Run()
	return cList
}
