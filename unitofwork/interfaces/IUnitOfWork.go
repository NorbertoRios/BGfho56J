package interfaces

import (
	"geometris-go/core/interfaces"
)

//IUnitOfWork ...
type IUnitOfWork interface {
	AddNewTasks(...interfaces.ITask)
	AddDirtyTasks(...interfaces.ITask)
	AddDirtyStates(...interfaces.IDeviceState)
	Commit()
}
