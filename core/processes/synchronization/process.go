package synchronization

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/response"
	"geometris-go/core/processes/synchronization/task"
)

//New ...
func New() interfaces.IProcess {
	process := &Process{}
	process.History = list.New()
	return process
}

//Process ...
type Process struct {
	process.Process
}

//NewRequest ...
func (p *Process) NewRequest(_request interface{}, _device interfaces.IDevice) interfaces.IProcessResponse {
	return p.TasksCompetitiveness(task.New(), _device)
}

//TasksCompetitiveness ...
func (p *Process) TasksCompetitiveness(_newTask interfaces.ITask, _device interfaces.IDevice) interfaces.IProcessResponse {
	resp := response.NewProcessResponse()
	if p.CurrentTask != nil {
		p.CurrentTask.Stop("Deprecated")
		p.SaveTask(p.CurrentTask)
		resp.AppendDirtyTask(p.CurrentTask)
	}
	p.CurrentTask = _newTask
	p.ExecuteCommands(p.CurrentTask.Start(), _device)
	resp.AppendNewTask(_newTask)
	return resp
}
