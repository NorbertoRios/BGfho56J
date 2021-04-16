package response

import "geometris-go/core/interfaces"

//NewDirtyState ...
func NewDirtyState(_identity string, _syncParams map[string]string, _state interfaces.IDeviceState, _rawData []byte) interfaces.IDirtyState {
	return &DirtyState{
		identity:  _identity,
		state:     _state,
		syncParam: _syncParams,
		rawData:   _rawData,
	}
}

//DirtyState ...
type DirtyState struct {
	state     interfaces.IDeviceState
	identity  string
	syncParam map[string]string
	rawData   []byte
}

//RawData ...
func (ds *DirtyState) RawData() []byte {
	return ds.rawData
}

//State ...
func (ds *DirtyState) State() interfaces.IDeviceState {
	return ds.state
}

//Identity ...
func (ds *DirtyState) Identity() string {
	return ds.identity
}

//SyncParams ...
func (ds *DirtyState) SyncParams() map[string]string {
	return ds.syncParam
}
