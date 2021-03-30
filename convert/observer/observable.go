package observer

import (
	"geometris-go/core/sensors"
)

//NewObservable ..
func NewObservable() *Observable {
	return &Observable{
		observers: map[IObserver]bool{
			NewAnyMessageObserver():  true,
			NewGPSValidityObserver(): true,
			NewOdometerObserver():    true,
			NewPowerState():          true,
			NewTimeObserver():        true,
		},
	}
}

//Observable ...
type Observable struct {
	observers map[IObserver]bool
}

//Notify ...
func (o *Observable) Notify(_sensor sensors.ISensor) map[string]interface{} {
	hash := make(map[string]interface{})
	for observer, _ := range o.observers {
		res := observer.Convert(_sensor)
		for k, v := range res {
			hash[k] = v
		}
	}
	return hash
}
