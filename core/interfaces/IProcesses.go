package interfaces

//IProcesses ...
type IProcesses interface {
	Immobilizer(int, string) IProcess
	Synchronization() IProcess
	Configuration() IProcess
	All() []IProcess
}
