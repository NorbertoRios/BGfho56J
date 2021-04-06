package interfaces

//IConfigTask ....
type IConfigTask interface {
	ITask
	Command() string
}
