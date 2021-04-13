package response

import "geometris-go/core/interfaces"

//NewDirtyState ...
func NewDirtyState(_identity, _syncParam string, _state interfaces.IDeviceState, _rawData []byte) interfaces.IDirtyState {
	return &DirtyState{
		identity:  _identity,
		state:     _state,
		syncParam: _syncParam,
		rawData:   _rawData,
	}
}

//DirtyState ...
type DirtyState struct {
	state     interfaces.IDeviceState
	identity  string
	syncParam string
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

//SyncParam ...
func (ds *DirtyState) SyncParam() string {
	return ds.syncParam
}
