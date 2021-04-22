package states

import (
	"geometris-go/core/interfaces"
	"geometris-go/response"
)

//NewDone ...
func NewDone(_request interfaces.IRequest) interfaces.ITaskState {
	return &Done{
		request: _request,
	}
}

//Done ...
type Done struct {
	Base
	request interfaces.IRequest
}

//IsClosed ...
func (s *Done) IsClosed() bool {
	return true
}

//FacadeResponse ....
func (s *Done) FacadeResponse() string {
	response := response.NewFacadeResponse(s.request.CallbackID(), "Done", true)
	return response.String()
}
