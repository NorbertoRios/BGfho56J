package message

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"

	"geometris-go/core/processes/message/states"
	"geometris-go/core/processes/message/task"
	"geometris-go/core/processes/response"
	message "geometris-go/message/interfaces"
	"geometris-go/message/types"
	messageTypes "geometris-go/message/types"
	"geometris-go/parser"
)

//New ...
func New(_syncParam string) interfaces.IProcess {
	p := &Process{
		Process: process.Process{
			History:     list.New(),
			CurrentTask: task.New(),
		},
		syncParam: _syncParam,
	}
	p.CurrentTask.ChangeState(states.NewInProgressState())
	return p
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
	if _, s := _message.(*types.RawLocationMessage); !s {
		return resp
	}
	messageParser := parser.New()
	locationMessage := messageParser.Parse(_message, p.syncParam).(*messageTypes.LocationMessage)
	commads := p.CurrentTask.NewMessageArrived(locationMessage, _device)
	p.ExecuteCommands(commads, _device)
	resp.AppendState(_device.State())
	return resp
}
