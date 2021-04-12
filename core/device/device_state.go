package device

import (
	"geometris-go/core/interfaces"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"sync"
	"time"
)

//NewStateBasedState ...
func NewStateBasedState(_lastState interfaces.IDeviceState, deviceSensors []sensors.ISensor) interfaces.IDeviceState {
	hash := _lastState.StateMap()
	for _, sensor := range deviceSensors {
		if sensor.Symbol() == "" {
			logger.Logger().WriteToLog(logger.Error, "[NewStateBasedState] Cant find symbol for sensor. Sensor: ", sensor)
			continue
		}
		hash[sensor.Symbol()] = sensor
	}
	return newState(hash)
}

//NewSensorBasedState ...
func NewSensorBasedState(_sensors []sensors.ISensor) interfaces.IDeviceState {
	deviceSensors := make(map[string]sensors.ISensor)
	for _, sensor := range _sensors {
		deviceSensors[sensor.Symbol()] = sensor
	}
	return newState(deviceSensors)
}

func newState(_sensorMap map[string]sensors.ISensor) interfaces.IDeviceState {
	return &State{
		mutex:         &sync.Mutex{},
		deviceSensors: _sensorMap,
		updateTime:    time.Now().UTC(),
	}
}

//State ...
type State struct {
	mutex         *sync.Mutex
	deviceSensors map[string]sensors.ISensor
	updateTime    time.Time
}

//StateMap ...
func (s *State) StateMap() map[string]sensors.ISensor {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.deviceSensors
}

//State ...
func (s *State) State() []sensors.ISensor {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deviceSensors := []sensors.ISensor{}
	for _, v := range s.deviceSensors {
		deviceSensors = append(deviceSensors, v)
	}
	return deviceSensors
}
