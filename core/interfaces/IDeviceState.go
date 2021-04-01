package interfaces

import "geometris-go/core/sensors"

//IDeviceState ..
type IDeviceState interface {
	Sensors() []sensors.ISensor
	State() map[string]sensors.ISensor
}
