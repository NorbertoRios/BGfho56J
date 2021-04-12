package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewTimes ...
func NewTimes() IObserver {
	return &Times{
		typeValue: "time",
	}
}

//Times ...
type Times struct {
	typeValue string
}

//Convert ...
func (f *Times) Convert(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	sensorsArr = append(sensorsArr, sensors.NewSensor(_key, _value))
	return sensorsArr
}

//Build ...
func (f *Times) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	strValue := types.NewString(_value)
	intTimeValue := strValue.Int(32)
	sensorValue := types.NewTime(int64(intTimeValue.(int32)))
	if _key == "TimeStamp" {
		sensorsArr = append(sensorsArr, sensors.NewSensor("Ack", intTimeValue))
	}
	sensorsArr = append(sensorsArr, sensors.NewSensor(_key, sensorValue.String()))
	return sensorsArr
}
