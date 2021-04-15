package interfaces

//IProcesses ...
type IProcesses interface {
	RemoveProcess(string)
	BOCommand(string) IProcess
	Immobilizer(int, string) IProcess
	ProcessComplete(string)
	Synchronization() IProcess
	LocationRequest() IProcess
	NewLocationProcess(string)
	Configuration() IProcess
	All() []IProcess
}
