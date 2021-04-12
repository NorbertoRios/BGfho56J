package unitofwork

import (
	"geometris-go/core/interfaces"
	"geometris-go/repository"
	unitofwork "geometris-go/unitofwork/interfaces"
	"sync"
)

//New ...
func New(_mysql, _rabbit repository.IRepository) unitofwork.IUnitOfWork {
	return &UnitOfWork{
		mysqlRepository:  _mysql,
		rabbitRepository: _rabbit,
		newTasks:         []interfaces.ITask{},
		dirtyTasks:       []interfaces.ITask{},
		dirtyStates:      []interfaces.IDirtyState{},
		mutex:            &sync.Mutex{},
	}
}

//UnitOfWork ...
type UnitOfWork struct {
	mysqlRepository  repository.IRepository
	rabbitRepository repository.IRepository
	newTasks         []interfaces.ITask
	dirtyTasks       []interfaces.ITask
	dirtyStates      []interfaces.IDirtyState
	mutex            *sync.Mutex
}

//AddNewTasks ...
func (uow *UnitOfWork) AddNewTasks(task ...interfaces.ITask) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.newTasks = append(uow.newTasks, task...)
}

//AddDirtyTasks ...
func (uow *UnitOfWork) AddDirtyTasks(task ...interfaces.ITask) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.dirtyTasks = append(uow.dirtyTasks, task...)
}

//AddDirtyStates ...
func (uow *UnitOfWork) AddDirtyStates(states ...interfaces.IDirtyState) {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.dirtyStates = append(uow.dirtyStates, states...)
}

//Commit ...
func (uow *UnitOfWork) Commit() {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()
	uow.rabbitRepository.Save(uow.dirtyStates)
	uow.rabbitRepository.Save(uow.dirtyTasks)
	uow.rabbitRepository.Save(uow.newTasks)
	uow.mysqlRepository.Save(uow.dirtyStates)
	uow.mysqlRepository.Save(uow.dirtyTasks)
	uow.mysqlRepository.Save(uow.newTasks)
	uow.dirtyStates = []interfaces.IDirtyState{}
	uow.dirtyTasks = []interfaces.ITask{}
	uow.newTasks = []interfaces.ITask{}
}
