package usecase

import (
	core "geometris-go/core/interfaces"
	"geometris-go/unitofwork/interfaces"
)

//NewAPIRequestUseCase ...
func NewAPIRequestUseCase(_uow interfaces.IUnitOfWork, _device core.IDevice) *APIRequestUseCase {
	usecase := &APIRequestUseCase{
		device: _device,
	}
	usecase.unitOfWork = _uow

	return usecase
}

//APIRequestUseCase ...
type APIRequestUseCase struct {
	Base
	device core.IDevice
}

//Launch ...
func (usecase *APIRequestUseCase) Launch(_request core.IRequest, extractor func(core.IDevice, core.IRequest) core.IExtractor) {
	process := extractor(usecase.device, _request).Extract()
	processResp := process.NewRequest(_request, usecase.device)
	usecase.flushProcessResults(processResp)
	usecase.unitOfWork.Commit()
}
