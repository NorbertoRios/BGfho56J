package interfaces

import (
	"container/list"
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/sensors"
	"time"
)

//IDevice ...
type IDevice interface {
	NewSourseID(uint64)
	Send(string) bool
	State() IDeviceState
	LastActivityTime() time.Time
	NewChannel(connInterfaces.IChannel)
	Channel() connInterfaces.IChannel
	NewState([]sensors.ISensor)
	ProcessCommands(*list.List)
	Identity() string
	SourseID() uint64
	Processes() IProcesses
}
