package factory

import "geometris-go/core/sensors"

//New ...
func New() IFactory {
	_factories := []IFactory{
		NewBytes(),
		NewFloat32(),
		NewString(),
		NewTimes(),
		NewInt32(),
		NewReason(),
		NewDTC(),
		NewOdometer(),
	}
	return &Factory{
		factories: _factories,
	}
}

//Factory ...
type Factory struct {
	factories []IFactory
}

//Convert ...
func (o *Factory) Convert(_key string, _value interface{}, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	for _, observer := range o.factories {
		sensorsArr = append(sensorsArr, observer.Convert(_key, _value, _type)...)
	}
	return sensorsArr
}

//Build ...
func (o *Factory) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	for _, observer := range o.factories {
		sensorsArr = append(sensorsArr, observer.Build(_key, _value, _type)...)
	}
	return sensorsArr
}
