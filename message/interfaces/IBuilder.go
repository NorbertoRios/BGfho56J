package interfaces

//IBuilder ...
type IBuilder interface {
	Build(string) IMessage
	IsParsable([]byte) bool
}
