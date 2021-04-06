package configuration

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/configuration/task"
	"geometris-go/core/processes/response"
	"geometris-go/logger"
	message "geometris-go/message/interfaces"
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
	configReq, _ := _request.(interfaces.IConfigRequest)
	return p.TasksCompetitiveness(task.New(configReq), _device)
}

//TasksCompetitiveness ...
func (p *Process) TasksCompetitiveness(_newTask interfaces.IConfigTask, _device interfaces.IDevice) interfaces.IProcessResponse {
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

//MessageArrived ...
func (p *Process) MessageArrived(_message message.IMessage, _device interfaces.IDevice) interfaces.IProcessResponse {
	resp := response.NewProcessResponse()
	commands := p.CurrentTask.NewMessageArrived(_message, _device)
	p.ExecuteCommands(commands, _device)
	if p.CurrentTask.IsClosed() {
		logger.Logger().WriteToLog(logger.Info, "[Process | MessageArrived] Task with id "+p.CurrentTask.Request().CallbackID()+". Is Closed")
		resp.AppendDirtyTask(p.CurrentTask)
		p.SaveTask(p.CurrentTask)
		p.CurrentTask = nil
		if p.CuncelFunc != nil {
			p.CuncelFunc()
		}
	}
	return resp
}
