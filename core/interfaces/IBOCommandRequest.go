package interfaces

//IBOCommandRequest ...
type IBOCommandRequest interface {
	IRequest
	Command() string
}
