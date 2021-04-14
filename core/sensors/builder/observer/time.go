package observer

import (
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"geometris-go/types"
	"time"
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
	time, err := time.Parse("2006-01-02T15:04:05Z", _value)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[Times | Convert] Error while unmarshal TimeStamp. ", err.Error())
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, &types.JSONTime{}))
	} else {
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, &types.JSONTime{Time: time}))
	}
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
	sensorsArr = append(sensorsArr, sensors.NewSensor(_key, sensorValue))
	return sensorsArr
}
