package task

import (
	"container/list"
	"geometris-go/core/interfaces"
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
	command := ""
	for cmd := i.curretn; cmd != nil; cmd = cmd.Next() {
		strCommand, _ := cmd.Value.(string)
		if len(command)+len(strCommand) <= 450 {
			command = command + strCommand
		} else {
			i.curretn = cmd.Next()
			break
		}
	}
	return command
}

//NextExist ...
func (i *CommandsManager) NextExist() bool {
	return !(i.curretn == nil)
}
