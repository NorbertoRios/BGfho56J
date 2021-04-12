package convert

import (
	"geometris-go/core/interfaces"
	"geometris-go/dto"
)

//NewStateToDTO ...
func NewStateToDTO(_state interfaces.IDeviceState) *StateToDTO {
	return &StateToDTO{
		state: _state,
	}
}

//StateToDTO ...
type StateToDTO struct {
	state interfaces.IDeviceState
}

//Convert ...
func (std *StateToDTO) Convert() dto.IMessage {
	dtoMessage := dto.NewMessage()
	for _, value := range std.state.State() {
		dtoMessage.SetValue(value.Symbol(), value.Value())
	}
	return dtoMessage
}
