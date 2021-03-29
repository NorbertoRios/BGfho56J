package observers

import (
	"geometris-go/adaptors/dto"
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewOutputsObserver ...
func NewOutputsObserver() *OutputsObserver {
	return &OutputsObserver{
		Symbol: "Relay",
	}
}

//OutputsObserver ...
type OutputsObserver struct {
	Symbol string
}

//Notify ...
func (o *OutputsObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		float := types.NewFloat64(v.(float64))
		return sensors.BuildOutputsFromString(float.String())
	}
	return nil
}
