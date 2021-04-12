package interfaces

//IProcessResponse ...
type IProcessResponse interface {
	AppendNewTask(ITask)
	AppendDirtyTask(ITask)
	AppendState(...IDirtyState)
	NewTasks() []ITask
	DirtyTasks() []ITask
	States() []IDirtyState
}
