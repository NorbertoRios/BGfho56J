package states

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states"
	"geometris-go/response"
)

//NewClosed ...
func NewClosed(_description string) interfaces.ITaskState {
	return &Closed{
		description: _description,
	}
}

//Closed ...
type Closed struct {
	description string
	states.Base
}

//IsClosed ...
func (s *Closed) IsClosed() bool {
	return true
}

//FacadeResponse ....
func (s *Closed) FacadeResponse() string {
	response := response.NewFacadeResponse("synchronization", s.description, false)
	return response.String()
}
