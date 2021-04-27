package builder

import (
	"geometris-go/core/sensors"
	"geometris-go/core/sensors/builder/factory"
)

//NewSensorBuilder ...
func NewSensorBuilder() *SensorBuilder {
	return &SensorBuilder{
		observable: factory.New(),
	}
}

//SensorBuilder ...
type SensorBuilder struct {
	observable factory.IFactory
}

//Build ...
func (sb *SensorBuilder) Build(_key, _value, _type string) []sensors.ISensor {
	return sb.observable.Build(_key, _value, _type)
}

//Convert ...
func (sb *SensorBuilder) Convert(_key string, _value interface{}, _type string) []sensors.ISensor {
	return sb.observable.Convert(_key, _value, _type)
}
