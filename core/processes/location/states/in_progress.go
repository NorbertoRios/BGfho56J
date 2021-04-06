package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/core/processes/watchdog"
	"geometris-go/message/types"
)

//NewInProgressState ...
func NewInProgressState(_message string, _duration int, _task interfaces.ITask) interfaces.IInProgressState {
	state := &InProgress{}
	state.Watchdog = watchdog.New(_task, states.NewAnyMessageState(_message, NewInProgressState), _duration)
	return state
}

//InProgress ...
type InProgress struct {
	states.InProgress
	request interfaces.IRequest
}

//NewMessageArrived ...
func (s *InProgress) NewMessageArrived(msg interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	switch msg.(type) {
	case *types.RawLocationMessage:
		{
			s.Watchdog.Stop()
			_task.ChangeState(states.NewDone(_task.Request().(interfaces.IRequest)))
		}
	}
	return list.New()
}
