package extractor

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/immobilizer/request"
)

//NewImmoExtractor ...
func NewImmoExtractor(_device interfaces.IDevice, _request interfaces.IRequest) interfaces.IExtractor {
	return &ImmoExtractor{
		device:  _device,
		request: _request,
	}
}

//ImmoExtractor ..
type ImmoExtractor struct {
	device  interfaces.IDevice
	request interfaces.IRequest
}

//Extract ...
func (e *ImmoExtractor) Extract() interfaces.IProcess {
	immoReq, s := e.request.(interfaces.IImmoRequest)
	if !s {
		return nil
	}
	process := e.device.Processes().Immobilizer(request.NewIndex(immoReq).Int(), immoReq.Trigger())
	return process
}
