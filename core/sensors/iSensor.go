package sensors

//ISensor sensor's intergace
type ISensor interface {
	Symbol() string
	Value() interface{}
}
