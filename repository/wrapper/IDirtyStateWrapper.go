package wrapper

import (
	"geometris-go/dto"
)

//IDirtyStateWrapper ..
type IDirtyStateWrapper interface {
	Identity() string
	SyncParam() string
	RawData() []byte
	DTOMessage() dto.IMessage
	ValueByKey(string) interface{}
}
