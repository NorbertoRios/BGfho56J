package factory

import (
	"fmt"
	"geometris-go/core/sensors"
)

//NewString ...
func NewString() IFactory {
	return &String{
		typeValue: "string",
	}
}

//String ...
type String struct {
	typeValue string
}

//Convert ...
func (f *String) Convert(_key string, _value interface{}, _type string) []sensors.ISensor {
	return f.Build(_key, fmt.Sprintf("%v", _value), _type)
}

//Build ...
func (f *String) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type == f.typeValue {
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, _value))
	}
	return sensorsArr
}
