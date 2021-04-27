package parser

import (
	"geometris-go/parser/interfaces"
	"strings"
)

//NewSynchParameter ...
func NewSynchParameter(_parameter string) interfaces.ISynchParameter {
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
