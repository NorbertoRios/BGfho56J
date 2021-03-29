package observers

import (
	"geometris-go/adaptors/dto"
	"geometris-go/core/sensors"
)

//NewDTCCodeObserver ...
func NewDTCCodeObserver() *DTCCodeObserver {
	return &DTCCodeObserver{
		Symbol: "DTCCode",
	}
}

//DTCCodeObserver ...
type DTCCodeObserver struct {
	Symbol string
}

//Notify ...
func (o *DTCCodeObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildDTCCodesSensorFromString(v.(string))
	}
	return nil
}
