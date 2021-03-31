package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/types"
	"time"
)

//NewTimes ...
func NewTimes() IObserver {
	return &Times{
		typeValue: "Time",
	}
}

//Times ...
type Times struct {
	typeValue string
}

//Build ...
func (f *Times) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	strValue := types.NewString(_value)
	sensorsArr = append(sensorsArr, sensors.NewSensor(_key, time.Unix(strValue.Int(64).(int64), 0)))
	sensorsArr = append(sensorsArr, sensors.NewSensor("ReceivedTime", time.Now().Unix()))
	return sensorsArr
}
