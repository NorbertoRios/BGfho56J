package commands

import (
	"container/list"
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
	"geometris-go/rabbitlogger"
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
		time.Sleep(time.Millisecond * 500)
		if !_device.Send(c.command) {
			logger.Logger().WriteToLog(logger.Error, "[SendMessageCommand | Execute] Something went wrong while sending message "+c.command)
		}
		rabbitlogger.Logger().Log(fmt.Sprintf("%v sent to device through receipt service.", c.command), _device.Identity())
		logger.Logger().WriteToLog(logger.Info, "[SendMessageCommand | Execute] Message "+c.command+" sent.")
	}()
	return list.New()
}
