package process

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/response"
	"geometris-go/logger"
	"geometris-go/message"
)

//Process ...
type Process struct {
	History     *list.List
	CurrentTask interfaces.ITask
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

}

//TasksCompetitiveness ...
func (p *Process) TasksCompetitiveness(_newTask interfaces.ITask, _device interfaces.IDevice) interfaces.IProcessResponse {
	resp := new(response.ProcessResponse)
	if p.CurrentTask != nil {
		p.CurrentTask.Stop("Deprecated")
		p.SaveTask(p.CurrentTask)
		resp.AppendDirtyTask(p.CurrentTask)
	}
	resp.AppendNewTask(_newTask)
	p.CurrentTask = _newTask
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
