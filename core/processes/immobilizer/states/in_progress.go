package states

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/core/processes/watchdog"
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
