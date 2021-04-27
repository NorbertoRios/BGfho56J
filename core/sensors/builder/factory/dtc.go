package factory

import (
	"geometris-go/core/sensors"
	"strings"
)

//NewDTC ...
func NewDTC() IFactory {
	return &DTC{
		typeValue: "dtc",
	}
}

//DTC ...
type DTC struct {
	typeValue string
}

//Convert ...
func (f *DTC) Convert(_key string, _value interface{}, _type string) []sensors.ISensor {
	sensorsArray := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArray
	}
	values := _value.([]interface{})
	dtc := []string{}
	for _, code := range values {
		dtc = append(dtc, code.(string))
	}
	sensorsArray = append(sensorsArray, sensors.NewSensor(_key, dtc))
	return sensorsArray
}

//Build ...
func (f *DTC) Build(_key, _value, _type string) []sensors.ISensor {
	sensorsArr := []sensors.ISensor{}
	if _type != f.typeValue {
		return sensorsArr
	}
	if _value == "0:0" {
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, []string{}))
		return sensorsArr
	}
	if !strings.Contains(_value, ":") {
		sensorsArr = append(sensorsArr, sensors.NewSensor(_key, []string{_value}))
	}
	values := strings.Split(_value, ":")
	codes := strings.Split(values[2], " ")
	sensorsArr = append(sensorsArr, sensors.NewSensor(_key, codes))
	return sensorsArr
}
