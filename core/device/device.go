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
func NewDevice(_identity string, _syncParam string, _sensors map[string]sensors.ISensor, _channel connInterfaces.IChannel) interfaces.IDevice {
	return &Device{
		identity:            _identity,
		LastActivity:        time.Now().UTC(),
		LastStateUpdateTime: time.Now().UTC(),
		CurrentState:        NewState(_sensors),
		UDPChannel:          _channel,
		Mutex:               &sync.Mutex{},
		processes:           manager.BuildProcesses(_syncParam),
	}
}

//Device struct
type Device struct {
	identity            string
	LastStateUpdateTime time.Time
	LastActivity        time.Time
	CurrentState        *State
	UDPChannel          connInterfaces.IChannel
	Mutex               *sync.Mutex
	processes           interfaces.IProcesses
}

//Send send command to device
func (device *Device) Send(message string) bool {
	err := device.UDPChannel.Send(message)
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[Device"+device.identity+" | Send] Error:", err)
		return false
	}
	return true
}

//ProcessCommands process commands
func (device *Device) ProcessCommands(commands *list.List) {
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
	device.LastStateUpdateTime = time.Now().UTC()
	device.CurrentState = NewSensorState(device.CurrentState, messageSensors)
}

//Identity ...
func (device *Device) Identity() string {
	return device.identity
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
	return device.processes
}

//New24Param ...
func (device *Device) New24Param(_param24 string) {

}

//SyncParam ...
func (device *Device) SyncParam() string {
	return ""
}
