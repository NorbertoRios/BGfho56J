package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/message/types"
)

//InProgress ...
type InProgress struct {
	Base
	AckMessage string
	Watchdog   interfaces.IWatchdog
}

//Pause ...
func (s *InProgress) Pause(_task interfaces.ITask) {
	s.Watchdog.Stop()
	_task.ChangeState(NewPauseState(s))
}

//NewMessageArrived ...
func (s *InProgress) NewMessageArrived(msg interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	switch msg.(type) {
	case *types.AckMessage:
		{
			ack, _ := msg.(*types.AckMessage)
			if ack.Commands() == s.AckMessage {
				s.Watchdog.Stop()
				_task.ChangeState(NewDone(_task.Request().(interfaces.IRequest)))
			}
		}
	}
	return list.New()
}

//Run ...
func (s *InProgress) Run() {
	s.Watchdog.Start()
}
