package observer

import "geometris-go/core/sensors"

//NewObservable ..
func NewObservable() *Observable {
	observers := []IObserver{
		NewBytes(),
		NewFloat32(),
		NewString(),
		NewTimes(),
		NewInt32(),
		NewReason(),
		NewDTC(),
	}
	return &Observable{
		observers: observers,
	}
}

//Observable ...
type Observable struct {
	observers []IObserver
}

//Convert ...
func (o *Observable) Convert(_key string, _value interface{}, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	for _, observer := range o.observers {
		sensorsArr = append(sensorsArr, observer.Convert(_key, _value, _type)...)
	}
	return sensorsArr
}

//Build ...
func (o *Observable) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	for _, observer := range o.observers {
		sensorsArr = append(sensorsArr, observer.Build(_key, _value, _type)...)
	}
	return sensorsArr
}
