package parser

import (
	"geometris-go/logger"
	"strings"
)

//NewSynchParameter ...
func NewSynchParameter(_parameter string) *SynchParameter {
	if _parameter == "" {
		logger.Logger().WriteToLog(logger.Info, "[NewSynchParameter] Synch parameter is empty")
	}
	return &SynchParameter{
		parameter: _parameter,
	}
}

//SynchParameter ...
type SynchParameter struct {
	parameter string
}

//ColumnsID ...
func (sp *SynchParameter) ColumnsID() []string {
	param12 := sp.parameter
	param12 = strings.Trim(param12, ";")
	return strings.Split(param12, ".")
}
