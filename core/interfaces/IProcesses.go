package interfaces

//IProcesses ...
type IProcesses interface {
	Immobilizer(int, string) IProcess
	Synchronization() IProcess
	LocationRequest() IProcess
	NewLocationProcess(string)
	Configuration() IProcess
	All() []IProcess
}
