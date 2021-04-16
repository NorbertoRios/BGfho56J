package interfaces

//IDirtyState ...
type IDirtyState interface {
	State() IDeviceState
	SyncParams() map[string]string
	Identity() string
	RawData() []byte
}
