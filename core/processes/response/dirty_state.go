package response

import "geometris-go/core/interfaces"

//NewDirtyState ...
func NewDirtyState(_identity, _syncParam string, _state interfaces.IDeviceState) interfaces.IDirtyState {
	return &DirtyState{
		identity:  _identity,
		state:     _state,
		syncParam: _syncParam,
	}
}

//DirtyState ...
type DirtyState struct {
	state     interfaces.IDeviceState
	identity  string
	syncParam string
}

//State ...
func (ds *DirtyState) State() interfaces.IDeviceState {
	return ds.state
}

//SyncParam ...
func (ds *DirtyState) Identity() string {
	return ds.syncParam
}

//SyncParam ...
func (ds *DirtyState) SyncParam() string {
	return ds.syncParam
}
