package request

import "strings"

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
	return strings.ReplaceAll(r.DeviceIdentity, "genx_", "")
}
