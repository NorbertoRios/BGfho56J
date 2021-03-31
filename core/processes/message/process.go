package message

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/core/processes/immobilizer/task"
	"geometris-go/core/processes/response"
	message "geometris-go/message/interfaces"
	messageTypes "geometris-go/message/types"
	"geometris-go/parser"
	"geometris-go/types"
)

//New ...
func New(_syncParam string) interfaces.IProcess {
	return &Process{
		Process: process.Process{
			History: list.New(),
		},
		syncParam: _syncParam,
	}
}

//Process ....
type Process struct {
	process.Process
	syncParam string
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

//MessageArrived ...
func (p *Process) MessageArrived(_message message.IMessage, _device interfaces.IDevice) interfaces.IProcessResponse {
	resp := response.NewProcessResponse()
	messageParser := parser.New(types.NewFile("/config/initialize/ReportConfiguration.xml"))
	locationMessage := messageParser.Parse(_message, p.syncParam).(*messageTypes.LocationMessage)
	p.CurrentTask.NewMessageArrived(locationMessage, _device)
	resp.AppendState(_device.State())
}
