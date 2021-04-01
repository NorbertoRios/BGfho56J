package states

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/states"
	"geometris-go/message/types"
)

//NewInProgressState ..
func NewInProgressState() interfaces.ITaskState {
	return &InProgress{}
}

//InProgress ...
type InProgress struct {
	states.Base
}

//NewMessageArrived ...
func (s *InProgress) NewMessageArrived(msg interface{}, _device interfaces.IDevice, _task interfaces.ITask) *list.List {
	locationMessage, f := msg.(*types.LocationMessage)
	cList := list.New()
	if !f || locationMessage == nil {
		return cList
	}
	_device.NewState(locationMessage.Sensors())
	cList.PushBack(commands.NewSendMessageCommand(locationMessage.Ack()))
	return cList
}
