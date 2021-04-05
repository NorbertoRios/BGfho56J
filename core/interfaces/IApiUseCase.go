package interfaces

//IAPIUseCase ...
type IAPIUseCase interface {
	Launch(IRequest, IDevice, IProcess)
}
