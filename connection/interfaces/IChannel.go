package interfaces

import "time"

//IChannel ...
type IChannel interface {
	Send(string) error
	Type() string
	LastActivity() time.Time
}
