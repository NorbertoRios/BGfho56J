package observers

import (
	"geometris-go/adaptors/dto"
	"geometris-go/core/sensors"
)

//NewFuelObserver ...
func NewFuelObserver() *FuelObserver {
	return &FuelObserver{
		Symbol: "Fuel",
	}
}

//FuelObserver ...
type FuelObserver struct {
	Symbol string
}

//Notify ...
func (o *FuelObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildFuelSensorFromString(v.(string))
	}
	return nil
}
