package synchronization

import (
	"container/list"
	"context"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/response"
	"geometris-go/core/processes/synchronization/task"
	"geometris-go/logger"
	message "geometris-go/message/interfaces"
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

//MessageArrived ...
func (p *Process) MessageArrived(_message message.IMessage, _device interfaces.IDevice) interfaces.IProcessResponse {
	resp := response.NewProcessResponse()
	commands := p.CurrentTask.NewMessageArrived(_message, _device)
	p.ExecuteCommands(commands, _device)
	if p.CurrentTask.IsClosed() {
		logger.Logger().WriteToLog(logger.Info, "[Synchronization | MessageArrived] Synch task Is Closed")
		resp.AppendDirtyTask(p.CurrentTask)
		p.SaveTask(p.CurrentTask)
		p.CurrentTask = nil
		if p.CuncelFunc != nil {
			p.CuncelFunc()
		}
	}
	return resp
}
