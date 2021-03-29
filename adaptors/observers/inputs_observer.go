package observers

import (
	"geometris-go/adaptors/dto"
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewInputsObserver ...
func NewInputsObserver() *InputsObserver {
	return &InputsObserver{
		Symbol: "GPIO",
	}
}

//InputsObserver ...
type InputsObserver struct {
	Symbol string
}

//Notify ...
func (o *InputsObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		float := types.NewFloat64(v.(float64))
		return sensors.BuildInputsFromString(float.String())
	}
	return nil
}
