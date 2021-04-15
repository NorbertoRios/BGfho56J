package wrapper

import (
	"geometris-go/convert"
	"geometris-go/core/interfaces"
	"geometris-go/core/sensors"
	"geometris-go/dto"
	"geometris-go/logger"
	"sync"
)

//NewDirtyStateWrapper ...
func NewDirtyStateWrapper(_state interfaces.IDirtyState) IDirtyStateWrapper {
	dto := convert.NewStateToDTO(_state.State().State()).Convert()
	dto.SetValue("DevId", _state.Identity())
	return &DirtyStateWrapper{
		state:      _state,
		sensors:    _state.State().StateMap(),
		mutex:      &sync.Mutex{},
		dtoMessage: dto,
	}
}

//DirtyStateWrapper ..
type DirtyStateWrapper struct {
	state      interfaces.IDirtyState
	sensors    map[string]sensors.ISensor
	dtoMessage dto.IMessage
	mutex      *sync.Mutex
}

//ValueByKey ...
func (dsw *DirtyStateWrapper) ValueByKey(_key string) interface{} {
	dsw.mutex.Lock()
	defer dsw.mutex.Unlock()
	if value, found := dsw.sensors[_key]; found {
		return value.Value()
	}
	logger.Logger().WriteToLog(logger.Error, "[DirtyStateWrapper | DirtyStateWrapper] Cant find value by key ", _key)
	return nil
}

//SyncParam ...
func (dsw *DirtyStateWrapper) SyncParam() string {
	return dsw.state.SyncParam()
}

//RawData ...
func (dsw *DirtyStateWrapper) RawData() []byte {
	return dsw.state.RawData()
}

//Identity ...
func (dsw *DirtyStateWrapper) Identity() string {
	return dsw.state.Identity()
}

//DTOMessage ...
func (dsw *DirtyStateWrapper) DTOMessage() dto.IMessage {
	return dsw.dtoMessage
}
