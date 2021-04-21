package storage

import (
	"geometris-go/core/interfaces"
	"sync"
)

//NewMetric ...
func NewMetric() *Metric {
	return &Metric{
		devices: storage.(*DMStorage).Devices(),
		mutex:   &sync.Mutex{},
	}
}

//Metric ...
type Metric struct {
	devices map[string]interfaces.IDevice
	mutex   *sync.Mutex
}

//DevicesCount ...
func (m *Metric) DevicesCount() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return len(m.devices)
}

//ConnectionsCount ...
func (m *Metric) ConnectionsCount(_type string) int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	res := 0
	for _, device := range m.devices {
		switch device.Channel().Type() {
		case _type:
			res++
		}
	}
	return res
}
