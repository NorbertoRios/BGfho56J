package interfaces

import "geometris-go/core/sensors"

//IDeviceState ..
type IDeviceState interface {
	State() []sensors.ISensor
	StateMap() map[string]sensors.ISensor
}
