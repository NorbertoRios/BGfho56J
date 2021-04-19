package response

import "geometris-go/stats"

//ServiceStatistic ...
type ServiceStatistic struct {
	TotalDeviceCount             int
	UDPConnectionsCount          int
	TCPConnectionsCount          int
	UnregisteredConnectionsCount int
	TotalCountByWorkers          int
	ProcessInfo                  *stats.Process
}
