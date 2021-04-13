package message

import (
	"container/list"
	"fmt"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"geometris-go/logger"
	"sync"

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
		syncParam: _syncParam,
	}
	p.History = list.New()
	p.CurrentTask = task.New()
	p.ProcessSymbol = _symbol
	p.Mutex = &sync.Mutex{}
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
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	resp := response.NewProcessResponse()
	rawLocationMessage, s := _message.(*types.RawLocationMessage)
	if !s {
		return resp
	}
	messageParser := parser.New()
	locationMessage := messageParser.Parse(_message, p.syncParam).(*messageTypes.LocationMessage)
	logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[Message Arrived] Arrived new location message: %v", locationMessage.Sensors()))
	commads := p.CurrentTask.NewMessageArrived(locationMessage, _device)
	p.ExecuteCommands(commads, _device)
	dirtyState := response.NewDirtyState(_device.Identity(), p.syncParam, _device.State(), rawLocationMessage.RawByteData())
	resp.AppendState(dirtyState)
	return resp
}
