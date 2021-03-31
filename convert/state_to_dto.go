package convert

import (
	"geometris-go/dto"
	"geometris-go/message/interfaces"
)

//StateToDTO ...
type StateToDTO struct {
}

//Convert ...
func (std *StateToDTO) Convert(_message interfaces.IMessage) dto.IMessage {
	return nil
}
