package interfaces

import "container/list"

//ICommand ...
type ICommand interface {
	Execute(IDevice) *list.List
}
