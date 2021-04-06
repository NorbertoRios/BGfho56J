package task

import (
	"container/list"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
)

//NewCommandsManager ...
func NewCommandsManager(_commands []string) interfaces.ICommandsManager {
	commands := list.New()
	for _, c := range _commands {
		commands.PushBack(c)
	}
	return &CommandsManager{
		commands: commands,
		curretn:  commands.Front(),
	}
}

//CommandsManager ...
type CommandsManager struct {
	curretn  *list.Element
	commands *list.List
}

//Command ...
func (i *CommandsManager) Command() string {
	if i.curretn != nil {
		cmd := i.curretn.Value.(string)
		i.curretn = i.curretn.Next()
		return cmd
	}
	logger.Logger().WriteToLog(logger.Info, "[CommandsManager | Command] Commands is over")
	return ""
}

//NextExist ...
func (i *CommandsManager) NextExist() bool {
	return !(i.curretn.Next() == nil)
}
