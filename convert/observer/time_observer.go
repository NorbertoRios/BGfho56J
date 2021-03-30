package observer

import (
	"geometris-go/convert/dto"
	"geometris-go/core/sensors"
	"time"
)

//NewTimeObserver ...
func NewTimeObserver() IObserver {
	return &TimeObserver{
		symbols: []string{"TimeStamp"},
	}
}

//TimeObserver ...
type TimeObserver struct {
	symbols []string
}

//Convert ...
func (to *TimeObserver) Convert(_sensor sensors.ISensor) map[string]interface{} {
	hash := make(map[string]interface{})
	for _, s := range to.symbols {
		if _sensor.Symbol() == s {
			tm := &dto.Time{Time: time.Unix(_sensor.Value().(int64), 0)}
			hash[_sensor.Symbol()] = tm.String()
		}
	}
	return hash
}
