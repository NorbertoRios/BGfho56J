package interfaces

//IChangeOutputStateRequest ...
type IChangeOutputStateRequest interface {
	IRequest
	Port() string
}
