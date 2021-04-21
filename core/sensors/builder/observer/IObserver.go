package observer

import "geometris-go/core/sensors"

//IObserver ...
type IObserver interface {
	Build(string, string, string) []sensors.ISensor
	Convert(string, interface{}, string) []sensors.ISensor
}
