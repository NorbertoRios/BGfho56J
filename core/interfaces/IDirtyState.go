package interfaces

//IDirtyState ...
type IDirtyState interface {
	State() IDeviceState
	SyncParam() string
	Identity() string
	RawData() []byte
}
