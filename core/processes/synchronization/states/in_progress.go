package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/core/processes/watchdog"
	"geometris-go/message"
	"strings"
)

//NewInProgressState ..
func NewInProgressState(_paramMessage string, _duration int, _task interfaces.ITask) *InProgress {
	return &InProgress{
		watchdog: watchdog.New(_task, NewAnyMessageState(_paramMessage), _duration),
	}
}

//InProgress ...
type InProgress struct {
	states.Base
	watchdog *watchdog.Watchdog
}

//NewMessageArrived ...
func (s *InProgress) NewMessageArrived(msg interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	switch msg.(type) {
	case *message.ParametersMessage:
		{
			param, _ := msg.(*message.ParametersMessage)
			if value, f := param.Parameters["24"]; f {
				s.complete(_device, _task, value)
			}
		}
	case *message.AckMessage:
		{
			ack, _ := msg.(*message.AckMessage)
			param := strings.Split(ack.Value, "=")[0]
			if param == "24" {
				s.complete(_device, _task, ack.Value)
			}
		}
	}
	return list.New()
}

func (s *InProgress) complete(_device interfaces.IDevice, _task interfaces.ITask, _param24 string) {
	_device.New24Param(_param24)
	s.watchdog.Stop()
	_task.ChangeState(NewDone())
}

//Run ...
func (s *InProgress) Run() {
	s.watchdog.Start()
}
