package factory

import "geometris-go/core/sensors"

//IFactory ...
type IFactory interface {
	Build(string, string, string) []sensors.ISensor
	Convert(string, interface{}, string) []sensors.ISensor
}
