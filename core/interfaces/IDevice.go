package interfaces

import (
	"container/list"
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/sensors"
	"geometris-go/parser/interfaces"
)

//IDevice ...
type IDevice interface {
	NewSourseID(uint64)
	Send(string) bool
	State() IDeviceState
	NewChannel(connInterfaces.IChannel)
	Channel() connInterfaces.IChannel
	NewState([]sensors.ISensor)
	ProcessCommands(*list.List)
	Identity() string
	SourseID() uint64
	Processes() IProcesses
	Parser() interfaces.IParser
}
