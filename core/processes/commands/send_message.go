package commands

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
	"time"
)

//NewSendMessageCommand ...
func NewSendMessageCommand(_command string) interfaces.ICommand {
	return &SendMessageCommand{
		command: _command,
	}
}

//SendMessageCommand ...
type SendMessageCommand struct {
	command string
}

//Execute ...
func (c *SendMessageCommand) Execute(_device interfaces.IDevice) *list.List {
	go func() {
		time.Sleep(time.Millisecond * 700)
		if !_device.Send(c.command) {
			logger.Logger().WriteToLog(logger.Error, "[SendMessageCommand | Execute] Something went wrong while sending message "+c.command)
		}
		logger.Logger().WriteToLog(logger.Info, "[SendMessageCommand | Execute] Message "+c.command+" sent.")
	}()
	return list.New()
}
