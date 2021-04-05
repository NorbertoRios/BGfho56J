package interfaces

//IConfigTask ....
type IConfigTask interface {
	ITask
	CommandsManager() ICommandsManager
}
