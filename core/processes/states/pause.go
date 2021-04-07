package states

import "geometris-go/core/interfaces"

//NewPauseState ...
func NewPauseState(_state interface{}) interfaces.ITaskState {
	return &Pause{
		stateSnapshot: _state,
	}
}

//Pause ...
type Pause struct {
	Base
	stateSnapshot interface{}
}

//Resume ...
func (state *Pause) Resume(_task interfaces.ITask) {
	switch state.stateSnapshot.(type) {
	case interfaces.IInProgressState:
		{
			taskState := state.stateSnapshot.(interfaces.IInProgressState)
			taskState.Run()
			_task.ChangeState(taskState)
		}
	default:
		{
			taskState := state.stateSnapshot.(interfaces.ITaskState)
			_task.ChangeState(taskState)
		}
	}
}
