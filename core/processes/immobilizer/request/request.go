package request

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/request"
)

//NewImmoRequest ...
func NewImmoRequest(_callbackID, _identity, _port, _trigger, _state string, _safety bool) interfaces.IImmoRequest {
	req := &ImmoRequest{
		port:    _port,
		trigger: _trigger,
		state:   _state,
		safety:  _safety,
	}
	req.FacadeCallbackID = _callbackID
	req.DeviceIdentity = _identity
	return req
}

//ImmoRequest ...
type ImmoRequest struct {
	request.Request
	port    string
	trigger string
	state   string
	safety  bool
}

//Port ...
func (r *ImmoRequest) Port() string {
	return r.port
}

//Trigger ...
func (r *ImmoRequest) Trigger() string {
	return r.trigger
}

//Safety ...
func (r *ImmoRequest) Safety() bool {
	return r.safety
}

//State ...
func (r *ImmoRequest) State() string {
	return r.state
}
