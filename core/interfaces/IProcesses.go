package interfaces

//IProcesses ...
type IProcesses interface {
	Immobilizer(int, string) IProcess
	Synchronization() IProcess
	LocationRequest() IProcess
	Configuration() IProcess
	All() []IProcess
}
