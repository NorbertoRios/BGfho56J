package types

import (
	"fmt"
	"geometris-go/logger"
	"time"
)

//NewTimeFromString ...
func NewTimeFromString(_json string) *JSONTime {
	b := []byte(_json)
	sd := string(b[1 : len(b)-1])
	datetime, terr := time.Parse("2006-01-02T15:04:05Z", sd)
	if terr != nil {
		logger.Logger().WriteToLog(logger.Error, "[NewTimeFromString] Error while cteating new time value from string. ", _json)
		return nil
	}
	return &JSONTime{Time: datetime}
}

//NewTime ...
func NewTime(_unixTimeStamp int64) *JSONTime {
	return &JSONTime{Time: time.Unix(_unixTimeStamp, 0)}
}

//JSONTime ...
type JSONTime struct {
	time.Time
}

//MarshalJSON serialize timestamp
func (t *JSONTime) String() string {
	return fmt.Sprintf("%s", t.Format("2006-01-02T15:04:05Z"))
}
