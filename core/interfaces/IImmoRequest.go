package interfaces

//IImmoRequest ...
type IImmoRequest interface {
	IChangeOutputStateRequest
	Trigger() string
	Safety() bool
	State() string
}
