package request

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/request"
)

//NewBOCommandRequest ...
func NewBOCommandRequest(_identity, _callbackID string, _command string) interfaces.IBOCommandRequest {
	req := &BOCommandRequest{
		command: _command,
	}
	req.FacadeCallbackID = _callbackID
	req.DeviceIdentity = _identity
	return req
}

//BOCommandRequest ...
type BOCommandRequest struct {
	request.Request
	command string
}

//Command ...
func (r *BOCommandRequest) Command() string {
	return r.command
}
