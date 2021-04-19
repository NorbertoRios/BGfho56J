package repository

import (
	"errors"
	"geometris-go/logger"
	"geometris-go/repository/models"
	"geometris-go/repository/wrapper"
	"geometris-go/storage"

	"github.com/go-sql-driver/mysql"
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
	d := &models.DeviceActivity{Software: &models.Software{SyncParams: make(map[string]string)}}
	err := repo.connection.Where("daiDeviceIdentity=?", _identity).Find(d).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Logger().WriteToLog(logger.Info, "[MySQL | Load] Device "+_identity+" not found")
	}
	return d
}

//Save ...
func (repo *MySQL) Save(values ...interface{}) error {
	for _, value := range values {
		states, s := value.([]wrapper.IDirtyStateWrapper)
		if s {
			repo.connection.Transaction(func(tx *gorm.DB) error {
				for _, state := range states {
					device := storage.Storage().Device(state.Identity())
					state.DTOMessage().SetValue("PrevSourceId", device.SourseID())
					id, err := repo.saveHistory(state)
					if err != nil {
						logger.Logger().WriteToLog(logger.Fatal, "[MySQL | Save] Error in transaction while save message history. Error", err)
						return err
					}
					state.DTOMessage().NewSID(id)
					err = repo.saveActivity(id, state)
					if err != nil {
						logger.Logger().WriteToLog(logger.Fatal, "[MySQL | Save] Error in transaction while save device activity. Error", err)
						return err
					}
					device.NewSourseID(id)
				}
				return nil
			})
		}
	}
	return nil
}

func (repo *MySQL) saveActivity(_sid uint64, _dirtyState wrapper.IDirtyStateWrapper) error {
	deviceActivity := models.NewDeviceActivity(_dirtyState, _sid)
	db := repo.connection.Model(&deviceActivity).Where("daiDeviceIdentity=?", deviceActivity.Identity).Updates(&deviceActivity)
	if db.RowsAffected == 0 {
		db = repo.connection.Create(&deviceActivity)
	}
	return db.Error
}

func (repo *MySQL) saveHistory(_dirtyState wrapper.IDirtyStateWrapper) (uint64, error) {
	messageHistory := models.NewMessageHistory(_dirtyState)
	err := repo.connection.Scopes(models.MessageHistoryTable(messageHistory)).Create(messageHistory).Error
	if err != nil {
		merr, ok := err.(*mysql.MySQLError)
		if ok && merr.Number == 1146 {
			err = repo.createMessageHistoryTable(messageHistory.DevID[len(messageHistory.DevID)-2:])
			if err == nil {
				err = repo.connection.Scopes(models.MessageHistoryTable(messageHistory)).Create(messageHistory).Error
			}
		}
	}
	return messageHistory.ID, err
}

//CreateMessageHistoryTable creates new table if table not exists detected
func (repo *MySQL) createMessageHistoryTable(tableName string) error {
	return repo.connection.Exec("CREATE TABLE IF NOT EXISTS  raw_data.`" + tableName + "` ( " +
		"`Id` bigint(20) NOT NULL AUTO_INCREMENT, " +
		"`DevId` varchar(100) NOT NULL, " +
		"`EntryData` blob, " +
		"`ParsedEntryData` blob, " +
		"`Time` datetime NOT NULL, " +
		"`RecievedTime` datetime NOT NULL, " +
		"`ReportClass` varchar(100) DEFAULT NULL, " +
		"`ReportType` int(11) DEFAULT NULL, " +
		"`Reason` varchar(5) DEFAULT NULL, " +
		"`Latitude` double DEFAULT NULL COMMENT 'degrees', " +
		"`Longitude` double DEFAULT NULL COMMENT 'degrees', " +
		"`Speed` double DEFAULT NULL, " +
		"`ValidFix` int(11) DEFAULT NULL, " +
		"`Altitude` double DEFAULT NULL, " +
		"`Heading` double DEFAULT NULL, " +
		"`IgnitionState` int(11) DEFAULT NULL, " +
		"`Odometer` int(10) DEFAULT NULL COMMENT 'm', " +
		"`Satellites` tinyint(3) unsigned DEFAULT NULL, " +
		"`Supply` int(10) DEFAULT NULL, " +
		"`GPIO` int(10) DEFAULT NULL COMMENT 'Input ports state', " +
		"`Relay` int(10) DEFAULT NULL COMMENT 'Output ports state', " +
		"`msg_id` binary(16) DEFAULT NULL, " +
		"`Extra` text, " +
		"`BatteryLow` double DEFAULT NULL, " +
		" PRIMARY KEY (`Id`,`Time`,`DevId`), " +
		"KEY `IX_RecievedTime` (`RecievedTime`,`DevId`) " +
		")" +
		"ENGINE = INNODB " +
		"AVG_ROW_LENGTH = 8192 " +
		"CHARACTER SET utf8 " +
		"COLLATE utf8_general_ci " +
		"PARTITION BY RANGE (to_days(Time)) " +
		"(" +
		"PARTITION p180201 VALUES LESS THAN (737091) ENGINE = InnoDB, " +
		"PARTITION p180301 VALUES LESS THAN (737119) ENGINE = InnoDB, " +
		"PARTITION p180401 VALUES LESS THAN (737150) ENGINE = InnoDB, " +
		"PARTITION p180501 VALUES LESS THAN (737180) ENGINE = InnoDB, " +
		"PARTITION p180601 VALUES LESS THAN (737211) ENGINE = InnoDB, " +
		"PARTITION p180701 VALUES LESS THAN (737241) ENGINE = InnoDB, " +
		"PARTITION p180801 VALUES LESS THAN (737272) ENGINE = InnoDB, " +
		"PARTITION p180901 VALUES LESS THAN (737303) ENGINE = InnoDB, " +
		"PARTITION p181001 VALUES LESS THAN (737333) ENGINE = InnoDB, " +
		"PARTITION p181101 VALUES LESS THAN (737364) ENGINE = InnoDB, " +
		"PARTITION p181201 VALUES LESS THAN (737394) ENGINE = InnoDB, " +
		"PARTITION p190101 VALUES LESS THAN (737425) ENGINE = InnoDB, " +
		"PARTITION p190201 VALUES LESS THAN (737456) ENGINE = InnoDB, " +
		"PARTITION p_cur VALUES LESS THAN MAXVALUE ENGINE = InnoDB " +
		");").Error
}
