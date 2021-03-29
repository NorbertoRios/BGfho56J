package request

import (
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
)

//NewIndex ...
func NewIndex(_request interfaces.IChangeOutputStateRequest) *Index {
	return &Index{
		request: _request,
	}
}

//Index ...
type Index struct {
	request interfaces.IChangeOutputStateRequest
}

//String ...
func (i *Index) String() string {
	return fmt.Sprintf("%v", i.Int())
}

//Int ...
func (i *Index) Int() int {
	var j int
	if _, err := fmt.Sscanf(i.request.Port(), "OUT%1d", &j); err == nil {
		return j + 1
	}
	logger.Logger().WriteToLog(logger.Error, "[Index | Int] Cant convert \"PORT\":"+i.request.Port()+" value to integer.")
	return -1
}
