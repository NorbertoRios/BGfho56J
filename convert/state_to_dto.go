package convert

import (
	"geometris-go/core/sensors"
	"geometris-go/dto"
)

//NewStateToDTO ...
func NewStateToDTO(_state []sensors.ISensor) *StateToDTO {
	return &StateToDTO{
		state: _state,
	}
}

//StateToDTO ...
type StateToDTO struct {
	state []sensors.ISensor
}

//Convert ...
func (std *StateToDTO) Convert() dto.IMessage {
	dtoMessage := dto.NewMessage()
	for _, value := range std.state {
		dtoMessage.SetValue(value.Symbol(), value.Value())
	}
	return dtoMessage
}
