package synchronization

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/synchronization/task"
	"sync"
)

//New ...
func New(_symbol string) interfaces.IProcess {
	process := &Process{}
	process.History = list.New()
	process.ProcessSymbol = _symbol
	process.Mutex = &sync.Mutex{}
	return process
}

//Process ...
type Process struct {
	process.Process
}

//NewRequest ...
func (p *Process) NewRequest(_request interface{}, _device interfaces.IDevice) interfaces.IProcessResponse {
	request, _ := _request.(string)
	return p.TasksCompetitiveness(task.New(request), _device)
}
