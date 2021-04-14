package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewBytes ...
func NewBytes() IObserver {
	return &Bytes{
		typeValue: "byte",
	}
}

//Bytes ...
type Bytes struct {
	typeValue string
}

//Convert ...
func (f *Bytes) Convert(_key, _value, _type string) []sensors.ISensor {
	return f.Build(_key, _value, _type)
}

//Build ...
func (f *Bytes) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	strValue := types.NewString(_value)
	value := strValue.Byte()
	sensorsArr = append(sensorsArr, sensors.NewSensor(_key, value))
	return sensorsArr
}
