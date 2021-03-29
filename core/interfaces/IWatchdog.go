package interfaces

//IWatchdog ...
type IWatchdog interface {
	Start(ITask)
	Stop()
	Started() bool
}
