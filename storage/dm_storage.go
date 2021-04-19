package storage

import (
	"geometris-go/core/interfaces"
	"geometris-go/logger"
	"sync"
)

var storage IDmStorage

//Storage ...
func Storage() IDmStorage {
	if storage == nil {
		storage = &DMStorage{make(map[string]interfaces.IDevice), &sync.Mutex{}}
	}
	return storage
}

//DMStorage ...
type DMStorage struct {
	devices map[string]interfaces.IDevice
	mutex   *sync.Mutex
}

//Devices ..
func (storage *DMStorage) Devices() map[string]interfaces.IDevice {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	return storage.devices
}

//DeviceExist ..
func (storage *DMStorage) DeviceExist(_identity string) bool {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	_, f := storage.devices[_identity]
	return f
}

//Device ..
func (storage *DMStorage) Device(_identity string) interfaces.IDevice {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	v, f := storage.devices[_identity]
	if !f {
		logger.Logger().WriteToLog(logger.Info, "[DMStorage | Device] Device with identity "+_identity+" not found")
	}
	return v
}

//AddDevice ...
func (storage *DMStorage) AddDevice(_device interfaces.IDevice) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	storage.devices[_device.Identity()] = _device
	logger.Logger().WriteToLog(logger.Info, "[DMStorage | Device] Device with identity "+_device.Identity()+" added")
}
