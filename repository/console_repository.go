package repository

import (
	"encoding/json"
	"geometris-go/convert"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
	"geometris-go/repository/models"
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
				convertor := convert.NewStateToDTO(state.(interfaces.IDirtyState).State().State())
				dtoMessage := convertor.Convert()
				jMessage, _ := json.Marshal(dtoMessage)
				logger.Logger().WriteToLog(logger.Info, "[ConsoleRepository_"+repo.repoType+" | Save] ", string(jMessage))
			}
		}
	}
	return nil
}

//Load ...
func (repo *ConsoleRepository) Load(key string) *models.DeviceActivity {
	logger.Logger().WriteToLog(logger.Info, "[ConsoleRepository_"+repo.repoType+" | Load] For ", key)
	return nil
}
