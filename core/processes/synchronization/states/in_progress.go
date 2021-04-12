package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/core/processes/watchdog"
	"geometris-go/message/types"
)

//NewInProgressState ..
func NewInProgressState(_paramMessage string, _duration int, _task interfaces.ITask) interfaces.IInProgressState {
	state := &InProgress{}
	state.Watchdog = watchdog.New(_task, states.NewAnyMessageState(_paramMessage, NewInProgressState), _duration)
	return state
}

//InProgress ...
type InProgress struct {
	states.InProgress
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
	_device.Processes().NewLocationProcess("12=" + _param24)
	s.Watchdog.Stop()
	_task.ChangeState(NewDone())
}
