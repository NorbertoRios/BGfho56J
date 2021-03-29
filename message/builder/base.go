package builder

import (
	"geometris-go/logger"
	"geometris-go/types"
	"regexp"
)

//Base ...
type Base struct {
	pattern *regexp.Regexp
}

func (b *Base) extractKeys(_key, _message string) string {
	names := b.pattern.SubexpNames()
	strArr := &types.StringArray{Data: names}
	if id, found := strArr.IndexOf(_key); found {
		return b.pattern.FindAllStringSubmatch(_message, -1)[0][id]
	}
	logger.Logger().WriteToLog(logger.Error, "[BaseBuilder] Cant extract parameter : ", _key, " from message : ", _message)
	return ""
}

//IsParsable ...
func (b *Base) IsParsable(_packet []byte) bool {
	return b.pattern.Match(_packet)
}
