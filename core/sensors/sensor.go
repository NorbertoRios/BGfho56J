package sensors

//NewSensor ..
func NewSensor(_id string, _symbol string, _value interface{}) ISensor {
	return &Sensor{
		symbol: _symbol,
		value:  _value,
		id:     _id,
	}
}

//Sensor ...
type Sensor struct {
	id     string
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

//ParamID ...
func (s *Sensor) ParamID() string {
	return s.id
}
