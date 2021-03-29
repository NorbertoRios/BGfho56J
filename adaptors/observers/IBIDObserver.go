package observers

import (
	"geometris-go/adaptors/dto"
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewIBIDObserver ...
func NewIBIDObserver() *IBIDObserver {
	return &IBIDObserver{
		Symbol: "IBID",
	}
}

//IBIDObserver ...
type IBIDObserver struct {
	Symbol string
}

//Notify ...
func (o *IBIDObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		float := types.NewFloat64(v.(float64))
		return sensors.BuildIButtonSensorFromString(float.String())
	}
	return nil
}
