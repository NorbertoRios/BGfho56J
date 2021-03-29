package states

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states/response"
)

//NewActual ...
func NewActual(_request interfaces.IRequest) interfaces.ITaskState {
	return &Actual{
		request: _request,
	}
}

//Actual ..
type Actual struct {
	Base
	request interfaces.IRequest
}

//IsClosed ...
func (s *Actual) IsClosed() bool {
	return true
}

//FacadeResponse ....
func (s *Actual) FacadeResponse() string {
	response := response.NewFacadeResponse(s.request.CallbackID(), "Actual", true)
	return response.String()
}
