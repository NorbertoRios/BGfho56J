package interfaces

//IProcesses ...
type IProcesses interface {
	Immobilizer(int, string) IProcess
	Synchronization() IProcess
	All() []IProcess
}
