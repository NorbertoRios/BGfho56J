package factory

import (
	"fmt"
	"geometris-go/core/sensors"
	"geometris-go/types"
)

//NewReason ...
func NewReason() IFactory {
	return &Reason{
		typeValue: "reason",
	}
}

//Reason ...
type Reason struct {
	typeValue string
}

//Convert ...
func (f *Reason) Convert(_key string, _value interface{}, _type string) []sensors.ISensor {
	return f.Build(_key, fmt.Sprintf("%v", _value), _type)
}

//Build ...
func (f *Reason) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	strValue := types.NewString(_value)
	sensorsArr = append(sensorsArr, sensors.NewSensor(_key, f.mapReasons(strValue.IntValue().(int))))
	sensorsArr = append(sensorsArr, sensors.NewSensor("ReportType", strValue.Int(32)))
	return sensorsArr
}

func (f *Reason) mapReasons(_deviceReason int) int32 {
	switch _deviceReason {
	case 0: //Power UP
		return 0x00
	case 1: //Ignition off
		return 0x02
	case 2: //Ignition on
		return 0x03
	case 4, 11: //PowerOff
		return 0x31
	case 9: //PowerOffBatt
		return 0x1F
	case 10: //IdleTimer
		return 0x10
	case 26: //BeginMove
		return 0x1D
	case 15: //SpeedingStart
		return 0x12
	case 28: //Speeding stop
		return 0x13
	case 29: //Deceleration
		return 62
	case 30: //Harsh turn
		return 100
	case 31: //Acceleration
		return 61
	case 48: //Jamming
		return 108
	default:
		return 0x06
	}
}
