package storage

import "geometris-go/core/interfaces"

//IDmStorage ...
type IDmStorage interface {
	AddDevice(interfaces.IDevice)
	Device(string) interfaces.IDevice
	DeviceExist(string) bool
	Devices() map[string]interfaces.IDevice
	RemoveDevice(string)
}
