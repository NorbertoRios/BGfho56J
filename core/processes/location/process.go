package location

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/location/task"
)

//New ...
func New(_symbol string) interfaces.IProcess {
	process := &Process{}
	process.History = list.New()
	process.ProcessSymbol = _symbol
	return process
}

//Process ...
type Process struct {
	process.Process
}

//NewRequest ....
func (p *Process) NewRequest(_request interface{}, _device interfaces.IDevice) interfaces.IProcessResponse {
	immoReq, _ := _request.(interfaces.IRequest)
	return p.TasksCompetitiveness(task.New(immoReq), _device)
}
