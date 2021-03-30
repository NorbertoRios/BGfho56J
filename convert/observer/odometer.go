package observer

import (
	"geometris-go/convert/dto"
	"geometris-go/core/sensors"
	"time"
)

//NewOdometerObserver ...
func NewOdometerObserver() IObserver {
	return &TimeObserver{
		symbols: []string{"OdometerKm", "OdometerMiles", "BusOdometerKM"},
	}
}

//OdometerObserver ...
type OdometerObserver struct {
	symbols []string
}

//Convert ...
func (to *OdometerObserver) Convert(_sensor sensors.ISensor) map[string]interface{} {
	hash := make(map[string]interface{})
	for _, s := range to.symbols {
		if _sensor.Symbol() == s {
			tm := &dto.Time{Time: time.Unix(_sensor.Value().(int64), 0)}
			hash[_sensor.Symbol()] = tm.String()
		}
	}
	return hash
}
