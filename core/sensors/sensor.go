package sensors

//NewSensor ..
func NewSensor(_symbol string, _value interface{}) ISensor {
	return &Sensor{
		symbol: _symbol,
		value:  _value,
	}
}

//Sensor ...
type Sensor struct {
	symbol string
	value  interface{}
}

//Symbol ...
func (s *Sensor) Symbol() string {
	return s.symbol
}

//Value ...
func (s *Sensor) Value() interface{} {
	return s.value
}
