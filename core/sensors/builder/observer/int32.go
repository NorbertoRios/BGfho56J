package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewInt32 ...
func NewInt32() IObserver {
	return &Int32{
		typeValue: "int32",
	}
}

//Int32 ...
type Int32 struct {
	typeValue string
}

//Convert ...
func (f *Int32) Convert(_key, _value, _type string) []sensors.ISensor {
	return f.Build(_key, _value, _type)
}

//Build ...
func (f *Int32) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	strValue := types.NewString(_value)
	sensorsArr = append(sensorsArr, sensors.NewSensor(_key, strValue.Int(32)))
	if _key == "LocationAge" {
		if strValue.Int(32).(int32) > 0 {
			sensorsArr = append(sensorsArr, sensors.NewSensor("GpsValidity", byte(0)))
		} else {
			sensorsArr = append(sensorsArr, sensors.NewSensor("GpsValidity", byte(1)))
		}
	}
	return sensorsArr
}
