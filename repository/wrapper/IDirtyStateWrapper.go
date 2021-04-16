package wrapper

import (
	"geometris-go/dto"
)

//IDirtyStateWrapper ..
type IDirtyStateWrapper interface {
	Identity() string
	SyncParams() map[string]string
	RawData() []byte
	DTOMessage() dto.IMessage
	ValueByKey(string) interface{}
}
