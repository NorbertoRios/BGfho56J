package factory

import (
	"fmt"
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewFloat32 ...
func NewFloat32() IFactory {
	return &Float64{
		typeValue: "float32",
	}
}

//Float64 ...
type Float64 struct {
	typeValue string
}

//Convert ...
func (f *Float64) Convert(_key string, _value interface{}, _type string) []sensors.ISensor {
	return f.Build(_key, fmt.Sprintf("%v", _value), _type)
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
