package models

import (
	"geometris-go/core/interfaces"
	"geometris-go/repository/wrapper"
	"time"

	"gorm.io/gorm"
)

//NewMessageHistory ...
func NewMessageHistory(_deviceState interfaces.IDirtyState) *MessageHistory {
	deviceStateWrapper := wrapper.NewDirtyStateWrapper(_deviceState)
	return &MessageHistory{
		DevID:           deviceStateWrapper.Identity(),
		EntryData:       deviceStateWrapper.RawData(),
		ParsedEntryData: []byte(deviceStateWrapper.StringMessage()),
		Time:            deviceStateWrapper.TimeStamp(),
		RecievedTime:    time.Now().UTC(),
		//ReportClass: ,
		//ReportType
		//Reason: ,
		//Latitude: ,
		//Longitude: ,
		//Speed: ,
		//ValidFix: ,
		//Altitude: ,
		//Heading: ,
		//IgnitionState: ,
		//Odometer: ,
		//Satellites: ,
		//Supply: ,
		//GPIO: ,
		//Relay: ,
	}
	// Тут нужно дополнить Wrapper над IDirtyState чтобы достать нужные поля.
}

//MessageHistory struct
type MessageHistory struct {
	ID              uint64    `gorm:"column:ID;primary_key"`
	DevID           string    `gorm:"column:DevId"`
	EntryData       []byte    `gorm:"column:EntryData"`
	ParsedEntryData []byte    `gorm:"column:ParsedEntryData"`
	Time            time.Time `gorm:"column:Time"`
	RecievedTime    time.Time `gorm:"column:RecievedTime"`
	ReportClass     string    `gorm:"column:ReportClass"`
	ReportType      int32     `gorm:"column:ReportType"`
	Reason          int32     `gorm:"column:Reason"`
	Latitude        float32   `gorm:"column:Latitude"`
	Longitude       float32   `gorm:"column:Longitude"`
	Speed           float32   `gorm:"column:Speed"`
	ValidFix        byte      `gorm:"column:ValidFix"`
	Altitude        float32   `gorm:"column:Altitude"`
	Heading         float32   `gorm:"column:Heading"`
	IgnitionState   byte      `gorm:"column:IgnitionState"`
	Odometer        int32     `gorm:"column:Odometer"`
	Satellites      byte      `gorm:"column:Satellites"`
	Supply          int32     `gorm:"column:Supply"`
	GPIO            byte      `gorm:"column:GPIO"`
	Relay           byte      `gorm:"column:Relay"`
}

//MessageHistoryTable ...
func MessageHistoryTable(h *MessageHistory) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tableName := "raw_data." + h.DevID[len(h.DevID)-2:]
		return tx.Table(tableName)
	}
}

//CreateMessageHistoryTable creates new table if table not exists detected
// func CreateMessageHistoryTable(tableName string) error {
// 	return rawdb.Exec("CREATE TABLE IF NOT EXISTS  raw_data.`" + tableName + "` ( " +
// 		"`Id` bigint(20) NOT NULL AUTO_INCREMENT, " +
// 		"`DevId` varchar(100) NOT NULL, " +
// 		"`EntryData` blob, " +
// 		"`ParsedEntryData` blob, " +
// 		"`Time` datetime NOT NULL, " +
// 		"`RecievedTime` datetime NOT NULL, " +
// 		"`ReportClass` varchar(100) DEFAULT NULL, " +
// 		"`ReportType` int(11) DEFAULT NULL, " +
// 		"`Reason` varchar(5) DEFAULT NULL, " +
// 		"`Latitude` double DEFAULT NULL COMMENT 'degrees', " +
// 		"`Longitude` double DEFAULT NULL COMMENT 'degrees', " +
// 		"`Speed` double DEFAULT NULL, " +
// 		"`ValidFix` int(11) DEFAULT NULL, " +
// 		"`Altitude` double DEFAULT NULL, " +
// 		"`Heading` double DEFAULT NULL, " +
// 		"`IgnitionState` int(11) DEFAULT NULL, " +
// 		"`Odometer` int(10) DEFAULT NULL COMMENT 'm', " +
// 		"`Satellites` tinyint(3) unsigned DEFAULT NULL, " +
// 		"`Supply` int(10) DEFAULT NULL, " +
// 		"`GPIO` int(10) DEFAULT NULL COMMENT 'Input ports state', " +
// 		"`Relay` int(10) DEFAULT NULL COMMENT 'Output ports state', " +
// 		"`msg_id` binary(16) DEFAULT NULL, " +
// 		"`Extra` text, " +
// 		"`BatteryLow` double DEFAULT NULL, " +
// 		" PRIMARY KEY (`Id`,`Time`,`DevId`), " +
// 		"KEY `IX_RecievedTime` (`RecievedTime`,`DevId`) " +
// 		")" +
// 		"ENGINE = INNODB " +
// 		"AVG_ROW_LENGTH = 8192 " +
// 		"CHARACTER SET utf8 " +
// 		"COLLATE utf8_general_ci " +
// 		"PARTITION BY RANGE (to_days(Time)) " +
// 		"(" +
// 		"PARTITION p180201 VALUES LESS THAN (737091) ENGINE = InnoDB, " +
// 		"PARTITION p180301 VALUES LESS THAN (737119) ENGINE = InnoDB, " +
// 		"PARTITION p180401 VALUES LESS THAN (737150) ENGINE = InnoDB, " +
// 		"PARTITION p180501 VALUES LESS THAN (737180) ENGINE = InnoDB, " +
// 		"PARTITION p180601 VALUES LESS THAN (737211) ENGINE = InnoDB, " +
// 		"PARTITION p180701 VALUES LESS THAN (737241) ENGINE = InnoDB, " +
// 		"PARTITION p180801 VALUES LESS THAN (737272) ENGINE = InnoDB, " +
// 		"PARTITION p180901 VALUES LESS THAN (737303) ENGINE = InnoDB, " +
// 		"PARTITION p181001 VALUES LESS THAN (737333) ENGINE = InnoDB, " +
// 		"PARTITION p181101 VALUES LESS THAN (737364) ENGINE = InnoDB, " +
// 		"PARTITION p181201 VALUES LESS THAN (737394) ENGINE = InnoDB, " +
// 		"PARTITION p190101 VALUES LESS THAN (737425) ENGINE = InnoDB, " +
// 		"PARTITION p190201 VALUES LESS THAN (737456) ENGINE = InnoDB, " +
// 		"PARTITION p_cur VALUES LESS THAN MAXVALUE ENGINE = InnoDB " +
// 		");").Error
// }
