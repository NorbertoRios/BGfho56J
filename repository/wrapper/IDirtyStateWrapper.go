package wrapper

import "time"

//IDirtyStateWrapper ..
type IDirtyStateWrapper interface {
	Identity() string
	StringMessage() string
	Firmware() string
	TimeStamp() time.Time
	RawData() []byte
}
