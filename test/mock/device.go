package mock

import (
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/device"
	"geometris-go/core/processes/manager"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"sync"
	"time"
)

//NewDevice ...
func NewDevice(_identity, _syncParam string, _sensors map[string]sensors.ISensor, _channel connInterfaces.IChannel) *Device {
	d := &Device{}
	d.DeviceIdentity = _identity
	d.LastActivity = time.Now().UTC()
	d.LastStateUpdateTime = time.Now().UTC()
	d.CurrentState = device.NewState(_sensors)
	d.UDPChannel = _channel
	d.Mutex = &sync.Mutex{}
	d.DeviceProcesses = manager.New(_syncParam)
	return d
}

//Device ..
type Device struct {
	device.Device
	lastSend string
}

//Send send command to device
func (device *Device) Send(message string) bool {
	logger.Logger().WriteToLog(logger.Info, "[Device "+device.Identity()+" | Send] Message:", message)
	device.lastSend = message
	return true
}

//LastSendCommand ...
func (device *Device) LastSendCommand() string {
	return device.lastSend
}
