package states

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
)

//NewDone ...
func NewDone() interfaces.ITaskState {
	return &Done{}
}

//Done ...
type Done struct {
	states.Base
}

//IsClosed ...
func (s *Done) IsClosed() bool {
	return true
}
