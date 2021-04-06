package process

import (
	"container/list"
	"context"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/response"
	"geometris-go/logger"
	message "geometris-go/message/interfaces"
	"reflect"
)

//Process ...
type Process struct {
	History     *list.List
	CurrentTask interfaces.ITask
	CuncelFunc  context.CancelFunc
}

//Start ...
func (p *Process) Start(_device interfaces.IDevice) {
	if p.CurrentTask == nil {
		return
	}
	p.CurrentTask.Start()
}

//Stop ...
func (p *Process) Stop(_device interfaces.IDevice, _description string) {
	if p.CurrentTask == nil {
		return
	}
	p.CurrentTask.Stop(_description)
}

//Pause ...
func (p *Process) Pause() {
	p.CurrentTask.Pause()
}

//Resume ...
func (p *Process) Resume() {
	p.CurrentTask.Resume()
}

//TasksCompetitiveness ...
func (p *Process) TasksCompetitiveness(_newTask interfaces.ITask, _device interfaces.IDevice) interfaces.IProcessResponse {
	resp := response.NewProcessResponse()
	if p.CurrentTask == nil {
		p.CurrentTask = _newTask
		p.ExecuteCommands(p.CurrentTask.Start(), _device)
		resp.AppendNewTask(_newTask)
	} else {
		if reflect.DeepEqual(_newTask.Request(), p.CurrentTask.Request()) {
			_newTask.Stop("Duplicate")
			resp.AppendDirtyTask(_newTask)
		} else {
			p.CurrentTask.Stop("Deprecated")
			resp.AppendDirtyTask(p.CurrentTask)
			p.CurrentTask = _newTask
			p.ExecuteCommands(p.CurrentTask.Start(), _device)
			resp.AppendNewTask(_newTask)
		}
	}
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

//ExecuteCommands ...
func (p *Process) ExecuteCommands(_commands *list.List, _device interfaces.IDevice) {
	for _commands.Len() > 0 {
		cmd := _commands.Front()
		command, valid := cmd.Value.(interfaces.ICommand)
		if valid {
			nList := command.Execute(_device)
			if nList != nil && nList.Len() > 0 {
				_commands.PushFrontList(nList)
			}
			_commands.Remove(cmd)
		}
	}
}

//TasksHistory ...
func (p *Process) TasksHistory() *list.List {
	return p.History
}

//Current ...
func (p *Process) Current() interfaces.ITask {
	return p.CurrentTask
}

//SaveTask ...
func (p *Process) SaveTask(_task interfaces.ITask) {
	p.History.PushBack(_task)
}
