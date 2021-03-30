package usecase

import (
	core "geometris-go/core/interfaces"
	"geometris-go/repository"
	"geometris-go/unitofwork"
)

//NewAPIRequestUseCase ...
func NewAPIRequestUseCase(_mysql, _rabbit repository.IRepository, _device core.IDevice) *APIRequestUseCase {
	usecase := &APIRequestUseCase{
		Base: Base{
			mysqlRepository:  _mysql,
			rabbitRepository: _rabbit,
		},
		device: _device,
	}
	return usecase
}

//APIRequestUseCase ...
type APIRequestUseCase struct {
	Base
	device core.IDevice
}

//Launch ...
func (usecase *APIRequestUseCase) Launch(_request core.IRequest, extractor func(core.IDevice, core.IRequest) core.IExtractor) {
	_uow := unitofwork.New(usecase.mysqlRepository, usecase.rabbitRepository)
	process := extractor(usecase.device, _request).Extract()
	processResp := process.NewRequest(_request, usecase.device)
	usecase.flushProcessResults(processResp, _uow)
	_uow.Commit()
}
