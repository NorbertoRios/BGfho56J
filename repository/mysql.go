package repository

import (
	"errors"
	"geometris-go/core/interfaces"
	"geometris-go/logger"
	"geometris-go/repository/models"

	sqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

//NewMySQL ...
func NewMySQL(_connectionString string) IRepository {
	conn, err := gorm.Open(sqlDriver.Open(_connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, "[MySQLRepository] Error while connecting to database. Error: ", err)
	}
	return &MySQL{
		connection: conn,
	}
}

//MySQL ...
type MySQL struct {
	connection *gorm.DB
}

//Load ...
func (repo *MySQL) Load(_identity string) *models.DeviceActivity {
	d := &models.DeviceActivity{}
	err := repo.connection.Where("daiDeviceIdentity=?", _identity).Find(d).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Logger().WriteToLog(logger.Info, "[MySQL | Load] Device "+_identity+" not found")
	}
	return d
}

//Save ...
func (repo *MySQL) Save(values ...interface{}) error {
	// for _, value := range values {
	// 	states, s := value.([]interfaces.IDirtyState)
	// 	if s {
	// 		for _, state := range states {
	// 			convertor := convert.NewStateToDTO(state.(interfaces.IDirtyState).State())
	// 			dtoMessage := convertor.Convert()
	// 			jMessage, _ := json.Marshal(dtoMessage)
	// 			activity := models.DeviceActivity
	// 		}
	// 	}
	// }
	// return nil
	return nil
}

func (repo *MySQL) saveActivity(_dirtyStates []interfaces.IDirtyState) {

}

func (repo *MySQL) saveHistory() uint64 {
	return 0
}
