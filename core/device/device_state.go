package device

import (
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"sync"
	"time"
)

//NewSensorState ...
func NewSensorState(_lastState *State, deviceSensors []sensors.ISensor) *State {
	hash := _lastState.State()
	for _, sensor := range deviceSensors {
		if sensor.Symbol() == "" {
			logger.Logger().WriteToLog(logger.Error, "[NewSensorState] Cant find symbol for sensor. Sensor: ", sensor)
			continue
		}
		hash[sensor.Symbol()] = sensor
	}
	return NewState(hash)
}

//NewState ...
func NewState(deviceSensors map[string]sensors.ISensor) *State {
	return &State{
		mutex:         &sync.Mutex{},
		deviceSensors: deviceSensors,
		updateTime:    time.Now().UTC(),
	}
}

//State ...
type State struct {
	mutex         *sync.Mutex
	deviceSensors map[string]sensors.ISensor
	updateTime    time.Time
}

//State ...
func (s *State) State() map[string]sensors.ISensor {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.deviceSensors
}

//Sensors ...
func (s *State) Sensors() []sensors.ISensor {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deviceSensors := []sensors.ISensor{}
	for _, v := range s.deviceSensors {
		deviceSensors = append(deviceSensors, v)
	}
	return deviceSensors
}
