package sensors

import "time"

//NewSensor ..
func NewSensor(_symbol string, _value interface{}) ISensor {
	return &Sensor{
		symbol:    _symbol,
		value:     _value,
		createdAt: time.Now().UTC(),
	}
}

//Sensor ...
type Sensor struct {
	symbol    string
	value     interface{}
	createdAt time.Time
}

//Symbol ...
func (s *Sensor) Symbol() string {
	return s.symbol
}

//Value ...
func (s *Sensor) Value() interface{} {
	return s.value
}

//CreatedAt ...
func (s *Sensor) CreatedAt() time.Time {
	return s.createdAt
}
