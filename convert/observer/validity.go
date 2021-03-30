package observer

import (
	"geometris-go/core/sensors"
)

//NewGPSValidityObserver ..
func NewGPSValidityObserver() IObserver {
	return &GPSValidityObserver{
		symbol: "Satellites",
	}
}

//GPSValidityObserver ...
type GPSValidityObserver struct {
	symbol string
}

//Convert ...
func (gvo *GPSValidityObserver) Convert(_sensor sensors.ISensor) map[string]interface{} {
	hash := make(map[string]interface{})
	if gvo.symbol == _sensor.Symbol() {
		value := _sensor.Value().(byte)
		if value > 8 {
			hash["GpsValidity"] = 1
		} else {
			hash["GpsValidity"] = 0
		}
	}
	return hash
}
