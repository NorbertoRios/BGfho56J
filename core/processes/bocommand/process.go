package bocommand

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/bocommand/task"
	"geometris-go/core/processes/response"
	"geometris-go/logger"
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
	resp := response.NewProcessResponse()
	request, _ := _request.(interfaces.IBOCommandRequest)
	p.CurrentTask = task.New(request)
	commands := p.CurrentTask.Start()
	p.ExecuteCommands(commands, _device)
	if p.CurrentTask.IsClosed() {
		logger.Logger().WriteToLog(logger.Info, "[BOProcess | NewRequest] Task with id "+p.CurrentTask.Request().CallbackID()+". Is Closed")
		resp.AppendDirtyTask(p.CurrentTask)
		p.CurrentTask = nil
		_device.Processes().RemoveProcess(p.Symbol())
	}
	return resp
}
