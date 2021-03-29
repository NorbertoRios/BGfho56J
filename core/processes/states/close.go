package states

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/states/response"
)

//NewClose ...
func NewClose(_request interfaces.IRequest, _description string) interfaces.ITaskState {
	return &Close{
		request:     _request,
		description: _description,
	}
}

//Close ...
type Close struct {
	Base
	request     interfaces.IRequest
	description string
}

//IsClosed ...
func (s *Close) IsClosed() bool {
	return true
}

//FacadeResponse ....
func (s *Close) FacadeResponse() string {
	response := response.NewFacadeResponse(s.request.CallbackID(), s.description, false)
	return response.String()
}
