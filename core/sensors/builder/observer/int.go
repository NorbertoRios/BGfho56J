package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewInts ...
func NewInts() IObserver {
	return &Ints{
		typeValue: "int",
	}
}

//Ints ...
type Ints struct {
	typeValue string
}

//Build ...
func (f *Ints) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type == f.typeValue {
		strValue := types.NewString(_value)
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, strValue.IntValue()))
	}
	return sensorsArr
}
