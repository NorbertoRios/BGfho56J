package usecase

import (
	core "geometris-go/core/interfaces"
	"geometris-go/repository"
	"geometris-go/unitofwork/interfaces"
)

//Base ...
type Base struct {
	mysqlRepository  repository.IRepository
	rabbitRepository repository.IRepository
}

func (b *Base) flushProcessResults(_response core.IProcessResponse, _uow interfaces.IUnitOfWork) {
	_uow.AddDirtyStates(_response.States()...)
	_uow.AddDirtyTasks(_response.DirtyTasks()...)
	_uow.AddNewTasks(_response.NewTasks()...)
}
