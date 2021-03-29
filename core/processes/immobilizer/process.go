package immobilizer

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/immobilizer/task"
	"geometris-go/core/processes/response"
)

//New ...
func New(_index int, _trigger string) interfaces.IProcess {
	process := &Process{
		index:   _index,
		trigger: _trigger,
	}
	process.History = list.New()
	return process
}

//Process ...
type Process struct {
	process.Process
	index   int
	trigger string
}

//NewRequest ....
func (p *Process) NewRequest(_request interface{}, _device interfaces.IDevice) interfaces.IProcessResponse {
	immoReq, s := _request.(interfaces.IImmoRequest)
	if !s {
		return response.NewProcessResponse()
	}
	return p.TasksCompetitiveness(task.New(immoReq), _device)
}

//TasksCompetitiveness ...
func (p *Process) TasksCompetitiveness(_newTask interfaces.ITask, _device interfaces.IDevice) interfaces.IProcessResponse {
	resp := response.NewProcessResponse()
	if p.CurrentTask != nil {
		p.CurrentTask.Stop("Deprecated")
		resp.AppendDirtyTask(p.CurrentTask)
	}
	p.CurrentTask = _newTask
	p.ExecuteCommands(p.CurrentTask.Start(), _device)
	resp.AppendNewTask(_newTask)
	return resp
}
