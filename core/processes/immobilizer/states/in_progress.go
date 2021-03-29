package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/core/processes/watchdog"
	"geometris-go/message"
)

//NewInProgressState ...
func NewInProgressState(_ackMessage string, _timeout int, _task interfaces.ITask) *InProgress {
	return &InProgress{
		ackMessage: _ackMessage,
		watchdog:   watchdog.New(_task, NewAnyMessageState(_ackMessage), _timeout),
	}
}

//InProgress ...
type InProgress struct {
	states.Base
	ackMessage string
	watchdog   *watchdog.Watchdog
}

//NewMessageArrived ...
func (s *InProgress) NewMessageArrived(msg interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	switch msg.(type) {
	case *message.AckMessage:
		{
			ack, _ := msg.(*message.AckMessage)
			if ack.Value == s.ackMessage {
				s.watchdog.Stop()
				_task.ChangeState(states.NewDone(_task.Request().(interfaces.IRequest)))
			}
		}
	}
	return list.New()
}

//Run ...
func (s *InProgress) Run() {
	s.watchdog.Start()
}
