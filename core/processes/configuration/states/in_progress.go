package states

import (
	"container/list"
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
	"geometris-go/core/processes/watchdog"
	"geometris-go/message/types"
)

//NewInProgressState ...
func NewInProgressState(_ackMessage string, _timeout int, _task interfaces.ITask) interfaces.IInProgressState {
	state := &InProgress{}
	state.AckMessage = _ackMessage
	state.Watchdog = watchdog.New(_task, states.NewAnyMessageState(_ackMessage, NewInProgressState), _timeout)
	return state
}

//InProgress ...
type InProgress struct {
	states.InProgress
}

//NewMessageArrived ...
func (s *InProgress) NewMessageArrived(msg interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	cList := list.New()
	switch msg.(type) {
	case *types.AckMessage:
		{
			ack, _ := msg.(*types.AckMessage)
			if ack.Commands() == s.AckMessage {
				s.Watchdog.Stop()
				if _task.(interfaces.IConfigTask).CommandsManager().NextExist() {
					command := fmt.Sprintf("SETPARAMS %v ACK", _task.(interfaces.IConfigTask).CommandsManager().Command())
					cList.PushBack(commands.NewSendMessageCommand(command))
					_task.ChangeState(NewInProgressState(command, 300, _task))
				} else {
					_task.ChangeState(states.NewDone(_task.Request().(interfaces.IRequest)))
				}
			}
		}
	}
	return cList
}
