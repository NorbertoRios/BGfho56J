package rabbitlogger

import (
	"geometris-go/repository"
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
func (l *rabbitLogger) Log(message string) {
	l.rabbitRepo.Save(message)
}
