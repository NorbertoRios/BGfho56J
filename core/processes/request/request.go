package request

import (
	"geometris-go/core/interfaces"
	"strings"
)

//NewRequest ...
func NewRequest(_callbackID, _identity string) interfaces.IRequest {
	return &Request{
		FacadeCallbackID: _callbackID,
		DeviceIdentity:   _identity,
	}
}

//Request ...
type Request struct {
	FacadeCallbackID string
	DeviceIdentity   string
}

//CallbackID ...
func (r *Request) CallbackID() string {
	return r.FacadeCallbackID
}

//Identity ...
func (r *Request) Identity() string {
	return r.DeviceIdentity
}

//Serial ...
func (r *Request) Serial() string {
	return strings.ReplaceAll(r.DeviceIdentity, "geometris_", "")
}
