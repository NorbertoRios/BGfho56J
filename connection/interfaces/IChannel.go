package interfaces

//IChannel ...
type IChannel interface {
	Send(string) error
	Type() string
}
