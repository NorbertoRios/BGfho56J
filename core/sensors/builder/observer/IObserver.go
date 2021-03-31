package observer

import "geometris-go/core/sensors"

//IObserver ...
type IObserver interface {
	Build(string, string, string) []sensors.ISensor
}
