package interfaces

//IController ...
type IController interface {
	Process([]byte, IChannel)
}
