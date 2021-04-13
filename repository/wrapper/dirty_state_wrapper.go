package wrapper

import (
	"encoding/json"
	"geometris-go/convert"
	"geometris-go/core/interfaces"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"geometris-go/types"
	"sync"
	"time"
)

//NewDirtyStateWrapper ...
func NewDirtyStateWrapper(_state interfaces.IDirtyState) IDirtyStateWrapper {
	return &DirtyStateWrapper{
		state:   _state,
		sensors: _state.State().StateMap(),
		mutex:   &sync.Mutex{},
	}
}

//DirtyStateWrapper ..
type DirtyStateWrapper struct {
	state   interfaces.IDirtyState
	sensors map[string]sensors.ISensor
	mutex   *sync.Mutex
}

func (dsw *DirtyStateWrapper) valueByKey(_key string) interface{} {
	dsw.mutex.Lock()
	defer dsw.mutex.Unlock()
	if value, found := dsw.sensors[_key]; found {
		return value.Value()
	}
	logger.Logger().WriteToLog(logger.Error, "[DirtyStateWrapper | DirtyStateWrapper] Cant find value by key ", _key)
	return nil
}

//RawData ...
func (dsw *DirtyStateWrapper) RawData() []byte {
	return dsw.state.RawData()
}

//Identity ...
func (dsw *DirtyStateWrapper) Identity() string {
	return dsw.state.Identity()
}

//StringMessage ...
func (dsw *DirtyStateWrapper) StringMessage() string {
	dtoMessage := convert.NewStateToDTO(dsw.state.State().State()).Convert()
	jMess, jErr := json.Marshal(dtoMessage)
	if jErr != nil {
		logger.Logger().WriteToLog(logger.Error, "[DirtyStateWrapper] Error while marshaling dto. ", jErr)
		jMess = []byte{}
	}
	return string(jMess)
}

//Firmware ...
func (dsw *DirtyStateWrapper) Firmware() string {
	value := dsw.valueByKey("Firmware")
	if value == nil {
		return ""
	}
	return value.(string)
}

//TimeStamp ...
func (dsw *DirtyStateWrapper) TimeStamp() time.Time {
	return dsw.valueByKey("TimeStamp").(*types.JSONTime).Time
}
