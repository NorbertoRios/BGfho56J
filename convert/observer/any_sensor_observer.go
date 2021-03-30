package observer

import (
	"geometris-go/core/sensors"
)

//NewAnyMessageObserver ..
func NewAnyMessageObserver() IObserver {
	return &AnyMessageObserver{
		symbols: []string{"TimeStamp", "OdometerKm", "OdometerMiles", "BusOdometerKM"},
	}
}

//AnyMessageObserver ...
type AnyMessageObserver struct {
	symbols []string
}

//Convert ...
func (amo *AnyMessageObserver) Convert(_sensor sensors.ISensor) map[string]interface{} {
	hash := make(map[string]interface{})
	for _, s := range amo.symbols {
		if s == _sensor.Symbol() {
			return hash
		}
	}
	hash[_sensor.Symbol()] = _sensor.Value()
	return hash
}
