package unitofwork

import (
	"geometris-go/core/interfaces"
	"geometris-go/repository"
	"geometris-go/repository/wrapper"
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
	var states []wrapper.IDirtyStateWrapper
	for _, ds := range uow.dirtyStates {
		states = append(states, wrapper.NewDirtyStateWrapper(ds))
	}
	uow.mysqlRepository.Save(states)
	uow.mysqlRepository.Save(uow.dirtyTasks)
	uow.rabbitRepository.Save(states)
	uow.rabbitRepository.Save(uow.dirtyTasks)
	uow.dirtyStates = []interfaces.IDirtyState{}
	uow.dirtyTasks = []interfaces.ITask{}
	uow.newTasks = []interfaces.ITask{}
}
