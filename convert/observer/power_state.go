package observer

import (
	"geometris-go/core/sensors"
)

//NewPowerState ...
func NewPowerState() IObserver {
	return &TimeObserver{
		symbols: []string{"Supply"},
	}
}

//PowerState ...
type PowerState struct {
	symbols []string
}

//Convert ...
func (ps *PowerState) Convert(_sensor sensors.ISensor) map[string]interface{} {
	hash := make(map[string]interface{})
	for _, s := range ps.symbols {
		if _sensor.Symbol() == s {
			supply := _sensor.Value().(int)
			var state string
			if supply > 0 {
				state = "Powered"
			} else {
				state = "Backup battery"
			}
			hash["PowerState"] = state
		}
	}
	return hash
}
