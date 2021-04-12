package device

import (
	"container/list"
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/manager"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"sync"
	"time"
)

//NewDevice ...
func NewDevice(_identity, _syncParameter string, _sensors []sensors.ISensor, _channel connInterfaces.IChannel) interfaces.IDevice {
	return &Device{
		DeviceIdentity:  _identity,
		LastActivity:    time.Now().UTC(),
		CurrentState:    NewSensorBasedState(_sensors),
		UDPChannel:      _channel,
		Mutex:           &sync.Mutex{},
		DeviceProcesses: manager.New(_syncParameter),
	}
}

//Device struct
type Device struct {
	DeviceIdentity  string
	LastActivity    time.Time
	CurrentState    interfaces.IDeviceState
	UDPChannel      connInterfaces.IChannel
	Mutex           *sync.Mutex
	DeviceProcesses interfaces.IProcesses
}

//Send send command to device
func (device *Device) Send(message string) bool {
	err := device.UDPChannel.Send(message)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[Device "+device.Identity()+" | Send] Error:", err)
		return false
	}
	return true
}

//ProcessCommands process commands
func (device *Device) ProcessCommands(commands *list.List) {
	device.LastActivity = time.Now().UTC()
	device.Mutex.Lock()
	defer device.Mutex.Unlock()
	for commands.Len() > 0 {
		cmd := commands.Front()
		command, valid := cmd.Value.(interfaces.ICommand)
		if valid {
			nList := command.Execute(device)
			if nList != nil && nList.Len() > 0 {
				commands.PushFrontList(nList)
			}
			commands.Remove(cmd)
		}
	}
}

//NewState ...
func (device *Device) NewState(messageSensors []sensors.ISensor) {
	device.CurrentState = NewStateBasedState(device.CurrentState, messageSensors)
}

//Identity ...
func (device *Device) Identity() string {
	return device.DeviceIdentity
}

//NewChannel ...
func (device *Device) NewChannel(_channel connInterfaces.IChannel) {
	device.UDPChannel = _channel
}

//LastActivityTime ...
func (device *Device) LastActivityTime() time.Time {
	return device.LastActivity
}

//State returns device current state
func (device *Device) State() interfaces.IDeviceState {
	device.Mutex.Lock()
	defer device.Mutex.Unlock()
	return device.CurrentState
}

//Processes ...
func (device *Device) Processes() interfaces.IProcesses {
	return device.DeviceProcesses
}
