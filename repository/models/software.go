package models

import (
	"encoding/json"
	"geometris-go/logger"
)

//NewSoftware ...
func NewSoftware(_syncParam, _firmware string) *Software {
	return &Software{
		SyncParam: _syncParam,
		Firmware:  _firmware,
	}
}

//Software ...
type Software struct {
	SyncParam string `json:"syncParam"`
	Firmware  string `json:"firmware"`
}

//Marshal ...
func (s *Software) Marshal() string {
	jSoft, jErr := json.Marshal(s)
	if jErr != nil {
		logger.Logger().WriteToLog(logger.Error, "[Software | Marshal] Error while marshaling software. ", jErr)
		return ""
	}
	return string(jSoft)
}
