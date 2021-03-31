package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewFloat64 ...
func NewFloat64() IObserver {
	return &Float64{
		typeValue: "float64",
	}
}

//Float64 ...
type Float64 struct {
	typeValue string
}

//Build ...
func (f *Float64) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type == f.typeValue {
		strValue := types.NewString(_value)
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, strValue.Float(64)))
	}
	return sensorsArr
}
