package request

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/request"
)

//NewConfigRequest ...
func NewConfigRequest(_identity, _callbackID string, _commands []string) interfaces.IConfigRequest {
	req := &Request{
		commands: _commands,
	}
	req.FacadeCallbackID = _callbackID
	req.DeviceIdentity = _identity
	return req
}

//Request ...
type Request struct {
	request.Request
	commands []string
}

//Commands ...
func (r *Request) Commands() []string {
	return r.commands
}
