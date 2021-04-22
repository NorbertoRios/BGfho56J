package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
	"strings"
)

//NewOdometer ...
func NewOdometer() IObserver {
	return &Odometer{
		typeValue: "odometer",
	}
}

//Odometer ...
type Odometer struct {
	typeValue string
}

//Convert ...
func (f *Odometer) Convert(_key string, _value interface{}, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	floatValue := _value.(float64)
	if strings.Contains(_key, "Mil") {
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, int32(floatValue/1609.34)))
	} else {
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, int32(floatValue/1000)))
	}
	return sensorsArr
}

//Build ...
func (f *Odometer) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	floatValue := types.NewString(_value).Float(32).(float32)
	if strings.Contains(_key, "Mil") {
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, int32(floatValue*1609.34)))
	} else {
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, int32(floatValue*1000)))
	}
	return sensorsArr
}
