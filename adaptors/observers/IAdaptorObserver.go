package observers

import (
	"geometris-go/adaptors/dto"
	"geometris-go/core/sensors"
)

//IAdaptorObserver ...
type IAdaptorObserver interface {
	Notify(*dto.DtoMessage) sensors.ISensor
}
