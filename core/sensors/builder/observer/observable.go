package observer

import "geometris-go/core/sensors"

//NewObservable ..
func NewObservable() *Observable {
	observers := []IObserver{
		NewBytes(),
		NewFloat64(),
		NewInts(),
		NewString(),
		NewTimes(),
		NewUInt32(),
	}
	return &Observable{
		observers: observers,
	}
}

//Observable ...
type Observable struct {
	observers []IObserver
}

//Build ...
func (o *Observable) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	for _, observer := range o.observers {
		sensorsArr = append(sensorsArr, observer.Build(_key, _value, _type)...)
	}
	return sensorsArr
}
