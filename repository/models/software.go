package models

import (
	"encoding/json"
	"geometris-go/logger"
)

//NewSoftware ...
func NewSoftware(_syncParam map[string]string, _firmware string) *Software {
	return &Software{
		SyncParams: _syncParam,
		Firmware:   _firmware,
	}
}

//Software ...
type Software struct {
	SyncParams map[string]string `json:"syncParams"`
	Firmware   string            `json:"firmware"`
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
