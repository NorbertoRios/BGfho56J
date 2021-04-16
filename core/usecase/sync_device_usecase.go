package usecase

import "geometris-go/core/interfaces"

//NewSyncDeviceUseCase ...
func NewSyncDeviceUseCase(_crc string) *SyncDeviceUseCase {
	return &SyncDeviceUseCase{
		crc: _crc,
	}
}

//SyncDeviceUseCase ...
type SyncDeviceUseCase struct {
	crc string
}

//Launch ...
func (sduc *SyncDeviceUseCase) Launch(_device interfaces.IDevice) {
	_device.Processes().Synchronization().NewRequest(sduc.crc, _device)
}
