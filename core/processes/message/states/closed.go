package states

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
)

//NewClosed ...
func NewClosed() interfaces.ITaskState {
	return &Closed{}
}

//Closed ...
type Closed struct {
	states.Base
}

//IsClosed ...
func (s *Closed) IsClosed() bool {
	return true
}
