package repository

import (
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
)

//NewConsoleRepository ....
func NewConsoleRepository(_repoType string) IRepository {
	return &ConsoleRepository{
		repoType: _repoType,
	}
}

//ConsoleRepository ...
type ConsoleRepository struct {
	repoType string
}

//Save ...
func (repo *ConsoleRepository) Save(values ...interface{}) error {
	for _, value := range values {
		states, s := value.([]interfaces.IDeviceState)
		if s {
			for _, state := range states {
				content := ""
				for _, value := range state.State() {
					content = fmt.Sprintf("%v%v;", content, value)
				}
				logger.Logger().WriteToLog(logger.Info, "[ConsoleRepository_"+repo.repoType+" | Save] ", content)
			}
			return nil
		}
	}
	return nil
}

//Load ...
func (repo *ConsoleRepository) Load(key string) interface{} {
	logger.Logger().WriteToLog(logger.Info, "[ConsoleRepository_"+repo.repoType+" | Load] For ", key)
	return nil
}
