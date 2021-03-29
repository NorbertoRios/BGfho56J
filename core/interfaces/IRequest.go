package interfaces

//IRequest ...
type IRequest interface {
	CallbackID() string
	Identity() string
	Serial() string
}
