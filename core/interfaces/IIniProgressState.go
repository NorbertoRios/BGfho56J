package interfaces

//IInProgressState ...
type IInProgressState interface {
	ITaskState
	Run()
}
