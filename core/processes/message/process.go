package message

import (
	"container/list"
	"geometris-go/core/interfaces"
	process "geometris-go/core/processes"
	"sync"

	"geometris-go/core/processes/message/task"
	"geometris-go/core/processes/response"
	message "geometris-go/message/interfaces"
	"geometris-go/message/types"
)

//New ...
func New(_symbol string, _synchParams map[string]string) interfaces.IProcess {
	p := &Process{}
	p.History = list.New()
	p.CurrentTask = task.New(_synchParams)
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
	return response.NewProcessResponse()
}

//MessageArrived ...
func (p *Process) MessageArrived(_message message.IMessage, _device interfaces.IDevice) interfaces.IProcessResponse {
	rawLocationMessage, s := _message.(*types.RawLocationMessage)
	if !s {
		return response.NewProcessResponse()
	}
	commands, resp := p.CurrentTask.(interfaces.ILocationTask).NewLocationMessageArrived(rawLocationMessage, _device)
	p.ExecuteCommands(commands, _device)
	return resp
}
