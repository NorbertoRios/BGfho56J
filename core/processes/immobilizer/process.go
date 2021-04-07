package immobilizer

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/immobilizer/task"
)

//New ...
func New(_index int, _trigger, _symbol string) interfaces.IProcess {
	process := &Process{
		index:   _index,
		trigger: _trigger,
	}
	process.History = list.New()
	process.ProcessSymbol = _symbol
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
	immoReq, _ := _request.(interfaces.IImmoRequest)
	return p.TasksCompetitiveness(task.New(immoReq), _device)
}
