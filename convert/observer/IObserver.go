package observer

import "geometris-go/core/sensors"

//IObserver ...
type IObserver interface {
	Convert(sensors.ISensor) map[string]interface{}
}
