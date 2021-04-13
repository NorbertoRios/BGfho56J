package models

import (
	"encoding/json"
	"geometris-go/configuration"
	"geometris-go/convert"
	"geometris-go/core/interfaces"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"geometris-go/repository/wrapper"
	"geometris-go/types"
	"time"

	"gorm.io/gorm"
)

//NewDeviceActivity ...
func NewDeviceActivity(_dirtyState interfaces.IDirtyState, _sourseID uint64) *DeviceActivity {
	deviceStateWrapper := wrapper.NewDirtyStateWrapper(_dirtyState)
	software := NewSoftware(_dirtyState.SyncParam(), deviceStateWrapper.Firmware())
	return &DeviceActivity{
		Identity:           deviceStateWrapper.Identity(),
		MessageTime:        deviceStateWrapper.TimeStamp(),
		LastUpdateTime:     time.Now().UTC(),
		LastMessage:        deviceStateWrapper.StringMessage(),
		LastMessageID:      _sourseID,
		SerializedSoftware: software.Marshal(),
		software:           software,
	}
}

//DeviceActivity device activity model
type DeviceActivity struct {
	Identity           string    `gorm:"column:daiDeviceIdentity"`
	MessageTime        time.Time `gorm:"column:daiLastMessageTime"`
	LastUpdateTime     time.Time `gorm:"column:daiLastUpdateTime"`
	LastMessageID      uint64    `gorm:"column:daiLastMessageId"`
	LastMessage        string    `gorm:"column:daiLastMessage"`
	SerializedSoftware string    `gorm:"column:daiSoftware"`
	software           *Software `gorm:"-" sql:"-"`
}

//TableName for DeviceActivity model
func (DeviceActivity) TableName() string {
	return "ats.tblDeviceActivityInfo"
}

//BeforeUpdate unmarshal string to struct
func (activity *DeviceActivity) BeforeUpdate(tx *gorm.DB) error {
	activity.SerializedSoftware = activity.software.Marshal()
	return nil
}

//AfterFind unmarshal string to struct
func (activity *DeviceActivity) AfterFind(tx *gorm.DB) error {
	software := &Software{}
	err := json.Unmarshal([]byte(activity.SerializedSoftware), software)
	activity.software = software
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[DeviceActivity | AfterFind] Error while unmarshaling software. ", err)
		return err
	}
	return nil
}

//State ...
func (activity *DeviceActivity) State() []sensors.ISensor {
	convertor := convert.NewDTOToState(activity.LastMessage, configuration.ReportConfig(types.NewFile("/config/initialize/ReportConfiguration.xml")))
	return convertor.Convert()
}
