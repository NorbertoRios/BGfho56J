package dto

import (
	"fmt"
	"time"
)

//Time ...
type Time struct {
	time.Time
}

func (t *Time) String() string {
	return fmt.Sprintf("\"%s\"", t.Format("2006-01-02T15:04:05Z"))
}
