package usecase

import (
	core "geometris-go/core/interfaces"
	"geometris-go/unitofwork/interfaces"
)

//Base ...
type Base struct {
	unitOfWork interfaces.IUnitOfWork
}

func (b *Base) flushProcessResults(_response core.IProcessResponse) {
	b.unitOfWork.AddDirtyStates(_response.States())
	b.unitOfWork.AddDirtyTasks(_response.DirtyTasks())
	b.unitOfWork.AddNewTasks(_response.NewTasks())
}
