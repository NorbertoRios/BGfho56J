package interfaces

//IConfigRequest ...
type IConfigRequest interface {
	IRequest
	Commands() []string
}
