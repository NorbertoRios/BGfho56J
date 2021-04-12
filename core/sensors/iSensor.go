package sensors

import "time"

//ISensor sensor's intergace
type ISensor interface {
	Symbol() string
	Value() interface{}
	CreatedAt() time.Time
}
