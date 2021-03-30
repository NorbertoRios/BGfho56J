package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/core/processes/watchdog"
	"geometris-go/message/types"
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
	case *types.ParametersMessage:
		{
			param, _ := msg.(*types.ParametersMessage)
			if value, f := param.Parameters()["12"]; f {
				s.complete(_device, _task, value)
			}
		}
	case *types.AckMessage:
		{
			ack, _ := msg.(*types.AckMessage)
			if value, f := ack.Parameters()["12"]; f {
				s.complete(_device, _task, value)
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
