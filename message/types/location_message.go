package types

import (
	"fmt"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"geometris-go/message/interfaces"
)

//NewLocationMessage ...
func NewLocationMessage(_identity string, _sensors []sensors.ISensor) interfaces.IMessage {
	return &LocationMessage{
		Base: Base{
			identity: _identity,
		},
		sensors: _sensors,
	}
}

//LocationMessage ...
type LocationMessage struct {
	Base
	sensors []sensors.ISensor
}

//Ack ...
func (lm *LocationMessage) Ack() string {
	for _, sensor := range lm.sensors {
		if sensor.Symbol() == "Ack" {
			return fmt.Sprintf("ACK %v", sensor.Value())
		}
	}
	logger.Logger().WriteToLog(logger.Error, "[LocationMessage | ACK] Cant find 36 parameter to ack message. Ack will be empty for device "+lm.identity)
	return ""
}

//Sensors ...
func (lm *LocationMessage) Sensors() []sensors.ISensor {
	return lm.sensors
}
