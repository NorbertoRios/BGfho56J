package states

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/response"
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

//FacadeResponse ...
func (s *Done) FacadeResponse() string {
	response := response.NewFacadeResponse("synchronization", "Done", true)
	return response.String()
}
