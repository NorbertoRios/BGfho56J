package rabbitlogger

import (
	"geometris-go/dto"
	"geometris-go/repository"
	"geometris-go/types"
	"time"
)

var rLogger *rabbitLogger

//BuildRabbitLogger ...
func BuildRabbitLogger(_rabbitRepository repository.IRepository) {
	rLogger = &rabbitLogger{
		rabbitRepo: _rabbitRepository,
	}
}

//Logger ...
func Logger() *rabbitLogger {
	return rLogger
}

type rabbitLogger struct {
	rabbitRepo repository.IRepository
}

//WriteToLog write content to log
func (l *rabbitLogger) Log(message, identity string) {
	dtoMessage := dto.NewMessage()
	dtoMessage.SetValue("Message", message)
	dtoMessage.SetValue("DevId", identity)
	dtoMessage.SetValue("TimeStamp", &types.JSONTime{Time: time.Now().UTC()})
	dtoMessage.SetValue("LocationMessage", false)
	l.rabbitRepo.Save(dtoMessage.Marshal())
}
