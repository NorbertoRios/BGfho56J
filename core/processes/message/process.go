package message

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/logger"

	"geometris-go/core/processes/message/task"
	"geometris-go/core/processes/response"
	message "geometris-go/message/interfaces"
	"geometris-go/message/types"
	messageTypes "geometris-go/message/types"
	"geometris-go/parser"
)

//New ...
func New(_syncParam, _symbol string) interfaces.IProcess {
	p := &Process{
		Process: process.Process{
			History:       list.New(),
			CurrentTask:   task.New(),
			ProcessSymbol: _symbol,
		},
		syncParam: _syncParam,
	}
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

//MessageArrived ...
func (p *Process) MessageArrived(_message message.IMessage, _device interfaces.IDevice) interfaces.IProcessResponse {
	resp := response.NewProcessResponse()
	if _, s := _message.(*types.RawLocationMessage); !s {
		return resp
	}
	messageParser := parser.New()
	locationMessage := messageParser.Parse(_message, p.syncParam).(*messageTypes.LocationMessage)
	logger.Logger().WriteToLog(logger.Info, "[Message Arrived] Arrived new location message: ", locationMessage)
	commads := p.CurrentTask.NewMessageArrived(locationMessage, _device)
	p.ExecuteCommands(commads, _device)
	resp.AppendState(_device.State())
	return resp
}
