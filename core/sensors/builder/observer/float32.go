package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewFloat32 ...
func NewFloat32() IObserver {
	return &Float64{
		typeValue: "float32",
	}
}

//Float64 ...
type Float64 struct {
	typeValue string
}

//Convert ...
func (f *Float64) Convert(_key, _value, _type string) []sensors.ISensor {
	return f.Build(_key, _value, _type)
}

//Build ...
func (f *Float64) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type == f.typeValue {
		strValue := types.NewString(_value)
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, strValue.Float(32)))
	}
	return sensorsArr
}
