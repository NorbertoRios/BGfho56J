package interfaces

//ICommandsManager ...
type ICommandsManager interface {
	Command() string
	NextExist() bool
}
