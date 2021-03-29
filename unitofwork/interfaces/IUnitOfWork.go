package interfaces

import "geometris-go/core/interfaces"

//IUnitOfWork ...
type IUnitOfWork interface {
	Device(string) interfaces.IDevice
	NewDevice(string) interfaces.IDevice
	AddNewTasks([]interfaces.ITask)
	AddDirtyTasks([]interfaces.ITask)
	AddDirtyStates([]interfaces.IDeviceState)
	Commit()
}
