package models

//NewSoftware ...
func NewSoftware(_syncParam, _firmware string) *Software {
	return &Software{
		SyncParam: _syncParam,
		Firmware:  _firmware,
	}
}

//Software ...
type Software struct {
	SyncParam string `json:"syncParam"`
	Firmware  string `json:"firmware"`
}
