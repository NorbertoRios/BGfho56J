package convert

import (
	"geometris-go/convert/observer"
	"geometris-go/dto"
	"geometris-go/message/interfaces"
	"geometris-go/message/types"
)

//NewStateToDTO ...
func NewStateToDTO() *StateToDTO {
	return &StateToDTO{
		observable: observer.NewObservable(),
	}
}

//StateToDTO ...
type StateToDTO struct {
	observable *observer.Observable
}

//Convert ...
func (std *StateToDTO) Convert(_message interfaces.IMessage) dto.IMessage {
	dtoMessage := dto.NewMessage()
	locationMessage, s := _message.(*types.LocationMessage)
	if !s {
		return dtoMessage
	}
	dtoMessage.SetValue("DevID", _message.Identity())
	for _, sensor := range locationMessage.Sensors() {
		hash := std.observable.Notify(sensor)
		dtoMessage.AppendRange(hash)
	}
	return dtoMessage
}
