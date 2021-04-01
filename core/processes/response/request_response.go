package response

import (
	"geometris-go/core/interfaces"
)

//NewProcessResponse ...
func NewProcessResponse() interfaces.IProcessResponse {
	return &ProcessResponse{
		newTasks:   []interfaces.ITask{},
		dirtyTasks: []interfaces.ITask{},
		states:     []interfaces.IDeviceState{},
	}
}

//ProcessResponse ...
type ProcessResponse struct {
	newTasks   []interfaces.ITask
	dirtyTasks []interfaces.ITask
	states     []interfaces.IDeviceState
}

//States ...
func (r *ProcessResponse) States() []interfaces.IDeviceState {
	return r.states
}

//AppendState ...
func (r *ProcessResponse) AppendState(_state ...interfaces.IDeviceState) {
	r.states = append(r.states, _state...)
}

//AppendNewTask ...
func (r *ProcessResponse) AppendNewTask(_task interfaces.ITask) {
	r.newTasks = append(r.newTasks, _task)
}

//AppendDirtyTask ...
func (r *ProcessResponse) AppendDirtyTask(_task interfaces.ITask) {
	r.dirtyTasks = append(r.dirtyTasks, _task)
}

//NewTasks ...
func (r *ProcessResponse) NewTasks() []interfaces.ITask {
	return r.newTasks
}

//DirtyTasks ...
func (r *ProcessResponse) DirtyTasks() []interfaces.ITask {
	return r.dirtyTasks
}
