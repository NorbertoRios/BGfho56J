package interfaces

import (
	"container/list"
	connInterfaces "geometris-go/connection/interfaces"
	"geometris-go/core/sensors"
	"time"
)

//IDevice ...
type IDevice interface {
	Send(string) bool
	State() IDeviceState
	LastActivityTime() time.Time
	NewChannel(connInterfaces.IChannel)
	NewState([]sensors.ISensor)
	ProcessCommands(*list.List)
	New24Param(string)
	Identity() string
	Processes() IProcesses
	BuildProcesses(string) []IProcess
}
