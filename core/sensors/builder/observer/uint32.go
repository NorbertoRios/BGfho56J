package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewUInt32 ...
func NewUInt32() IObserver {
	return &UInt32{
		typeValue: "uint32",
	}
}

//UInt32 ...
type UInt32 struct {
	typeValue string
}

//Convert ...
func (f *UInt32) Convert(_key, _value, _type string) []sensors.ISensor {
	return f.Build(_key, _value, _type)
}

//Build ...
func (f *UInt32) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type == f.typeValue {
		strValue := types.NewString(_value)
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, strValue.UInt(32)))
	}
	return sensorsArr
}
