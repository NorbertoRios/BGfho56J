package interfaces

//IProcessResponse ...
type IProcessResponse interface {
	AppendNewTask(ITask)
	AppendDirtyTask(ITask)
	AppendState(...IDeviceState)
	NewTasks() []ITask
	DirtyTasks() []ITask
	States() []IDeviceState
}
