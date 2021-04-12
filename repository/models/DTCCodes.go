package models

import (
	"geometris-go/core/interfaces"
	"time"
)

//NewDTCCodes ...
func NewDTCCodes(_state interfaces.IDeviceState) *DTCCodes {
	if codes, f := _state.StateMap()["DTC"]; !f {
		return &DTCCodes{Codes: []string{}, CodesUpdateDateTime: time.Now().UTC()}
	} else {
		return &DTCCodes{
			Codes:               codes.Value().([]string),
			CodesUpdateDateTime: codes.CreatedAt(),
		}
	}
}

//DTCCodes ...
type DTCCodes struct {
	Codes               []string
	CodesUpdateDateTime time.Time
}

//Count returns count of dtc codes
func (dtc *DTCCodes) Count() int {
	return len(dtc.Codes)
}
