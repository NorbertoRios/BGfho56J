package synchronization

import (
	"container/list"
	"context"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/synchronization/task"
)

//New ...
func New(_canselFunc context.CancelFunc) interfaces.IProcess {
	process := &Process{}
	process.History = list.New()
	process.CuncelFunc = _canselFunc
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
