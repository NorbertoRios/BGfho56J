package states

import (
	"container/list"
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/commands"
	"geometris-go/core/processes/response"
	"geometris-go/core/processes/states"
	"geometris-go/logger"
	"geometris-go/message/types"
	"geometris-go/parser"
	"sync"
)

//NewInProgressState ..
func NewInProgressState(_synchParams map[string]string) interfaces.ILocationInProgressState {
	return &InProgress{
		synchParams: _synchParams,
		mutex:       &sync.Mutex{},
	}
}

//InProgress ...
type InProgress struct {
	mutex       *sync.Mutex
	synchParams map[string]string
	states.InProgress
}

//Pause ...
func (s *InProgress) Pause(_task interfaces.ITask) {
	_task.ChangeState(states.NewPauseState(s))
}

//NewSynchParameter ...
func (s *InProgress) NewSynchParameter(crc, syncParam string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.synchParams[crc] = syncParam
}

//NewLocationMessageArrived ...
func (s *InProgress) NewLocationMessageArrived(msg *types.RawLocationMessage, resp interfaces.IProcessResponse, _device interfaces.IDevice) *list.List {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	cList := list.New()
	synchParameter, f := s.synchParams[msg.CRC()]
	if !f {
		logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[NewLocationMessageArrived] Cant find sync parameter for message: %v", msg.RawData()))
		return cList
	}
	messageParser := parser.New()
	locationMessage := messageParser.Parse(msg, synchParameter).(*types.LocationMessage)
	_device.NewState(locationMessage.Sensors())
	resp.AppendState(response.NewDirtyState(msg.Identity(), s.synchParams, _device.State(), msg.RawByteData()))
	cList.PushBack(commands.NewSendMessageCommand(locationMessage.Ack()))
	return cList
}

//Run ...
func (s *InProgress) Run() {
}
