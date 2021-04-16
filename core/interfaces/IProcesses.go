package interfaces

//IProcesses ...
type IProcesses interface {
	RemoveProcess(string)
	NewSynchParameter(string, string)
	BOCommand(string) IProcess
	Immobilizer(int, string) IProcess
	ProcessComplete(string)
	Synchronization() IProcess
	LocationRequest() IProcess
	Configuration() IProcess
	All() []IProcess
}
