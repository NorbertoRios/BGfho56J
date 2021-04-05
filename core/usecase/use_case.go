package usecase

import (
	core "geometris-go/core/interfaces"
	"geometris-go/repository"
	"geometris-go/unitofwork"
)

//NewAPIRequestUseCase ...
func NewAPIRequestUseCase(_mysql, _rabbit repository.IRepository) core.IAPIUseCase {
	usecase := &APIRequestUseCase{}
	usecase.mysqlRepository = _mysql
	usecase.rabbitRepository = _rabbit
	return usecase
}

//APIRequestUseCase ...
type APIRequestUseCase struct {
	Base
}

//Launch ...
func (usecase *APIRequestUseCase) Launch(_request core.IRequest, _device core.IDevice, _process core.IProcess) {
	_uow := unitofwork.New(usecase.mysqlRepository, usecase.rabbitRepository)
	processResp := _process.NewRequest(_request, _device)
	usecase.flushProcessResults(processResp, _uow)
	_uow.Commit()
}
