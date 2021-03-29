package relaydrive

import (
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/immobilizer/request"
)

//NewSetRelayDrive ...
func NewSetRelayDrive(_request interfaces.IImmoRequest) *SetRelayDrive {
	return &SetRelayDrive{
		request: _request,
	}
}

//SetRelayDrive ...
type SetRelayDrive struct {
	request interfaces.IImmoRequest
}

func (srd *SetRelayDrive) String() string {
	return fmt.Sprintf("SETRELAYDRIVE%v%v SERIALFILTER %v;BACKUPNVRAM", request.NewIndex(srd.request).Int(), request.NewState(srd.request).String(), srd.request.Serial())
}
