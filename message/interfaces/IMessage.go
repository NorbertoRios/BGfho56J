package interfaces

//IMessage ...
type IMessage interface {
	Identity() string
	Content() string
}
