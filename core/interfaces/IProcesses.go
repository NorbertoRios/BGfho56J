package interfaces

//IProcesses ...
type IProcesses interface {
	Immobilizer(int, string) IProcess
	ProcessComplete(string)
	Synchronization() IProcess
	LocationRequest() IProcess
	NewLocationProcess(string)
	Configuration() IProcess
	All() []IProcess
}
