package device

import (
	"container/list"
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/manager"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	iParser "geometris-go/parser/interfaces"
	"sync"
)

//NewDevice ...
func NewDevice(_identity string, _syncParameters map[string]string, _sid uint64, _sensors []sensors.ISensor, _channel connInterfaces.IChannel, _parser iParser.IParser) interfaces.IDevice {
	return &Device{
		sourseID:        _sid,
		DeviceIdentity:  _identity,
		CurrentState:    NewSensorBasedState(_sensors),
		UDPChannel:      _channel,
		Mutex:           &sync.Mutex{},
		DeviceProcesses: manager.New(_syncParameters),
		parser:          _parser,
	}
}

//Device struct
type Device struct {
	sourseID        uint64
	DeviceIdentity  string
	CurrentState    interfaces.IDeviceState
	UDPChannel      connInterfaces.IChannel
	Mutex           *sync.Mutex
	DeviceProcesses interfaces.IProcesses
	parser          iParser.IParser
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

//Parser ....
func (device *Device) Parser() iParser.IParser {
	return device.parser
}

//SourseID ...
func (device *Device) SourseID() uint64 {
	return device.sourseID
}

//NewSourseID ...
func (device *Device) NewSourseID(_id uint64) {
	device.sourseID = _id
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
	if len(messageSensors) == 0 {
		return
	}
	device.CurrentState = NewStateBasedState(device.CurrentState, messageSensors)
}

//Identity ...
func (device *Device) Identity() string {
	return device.DeviceIdentity
}

//Channel ...
func (device *Device) Channel() connInterfaces.IChannel {
	return device.UDPChannel
}

//NewChannel ...
func (device *Device) NewChannel(_channel connInterfaces.IChannel) {
	device.UDPChannel = _channel
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
