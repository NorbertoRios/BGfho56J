package builder

import (
	"geometris-go/core/sensors"
	"geometris-go/core/sensors/builder/observer"
)

//NewSensorBuilder ...
func NewSensorBuilder() *SensorBuilder {
	return &SensorBuilder{
		observable: observer.NewObservable(),
	}
}

//SensorBuilder ...
type SensorBuilder struct {
	observable *observer.Observable
}

//Build ...
func (sb *SensorBuilder) Build(_key, _value, _type string) []sensors.ISensor {
	return sb.observable.Build(_key, _value, _type)
}
