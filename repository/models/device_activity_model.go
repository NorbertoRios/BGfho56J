package models

import (
	"encoding/json"
	"geometris-go/configuration"
	"geometris-go/convert"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"geometris-go/repository/wrapper"
	"geometris-go/types"
	"time"

	"gorm.io/gorm"
)

//NewDeviceActivity ...
func NewDeviceActivity(_dirtyState wrapper.IDirtyStateWrapper, _sourseID uint64) *DeviceActivity {
	firmware := ""
	if v := _dirtyState.ValueByKey("Firmware"); v != nil {
		firmware = v.(string)
	}
	software := NewSoftware(_dirtyState.SyncParams(), firmware)
	return &DeviceActivity{
		Identity:           _dirtyState.Identity(),
		MessageTime:        _dirtyState.ValueByKey("TimeStamp").(*types.JSONTime).Time,
		LastUpdateTime:     time.Now().UTC(),
		LastMessage:        _dirtyState.DTOMessage().Marshal(),
		LastMessageID:      _sourseID,
		SerializedSoftware: software.Marshal(),
		Software:           software,
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
	Software           *Software `gorm:"-" sql:"-"`
}

//TableName for DeviceActivity model
func (DeviceActivity) TableName() string {
	return "ats.tblDeviceActivityInfo"
}

//BeforeUpdate unmarshal string to struct
func (activity *DeviceActivity) BeforeUpdate(tx *gorm.DB) error {
	activity.SerializedSoftware = activity.Software.Marshal()
	return nil
}

//AfterFind unmarshal string to struct
func (activity *DeviceActivity) AfterFind(tx *gorm.DB) error {
	software := &Software{}
	err := json.Unmarshal([]byte(activity.SerializedSoftware), software)
	activity.Software = software
	if err != nil {
		logger.Logger().WriteToLog(logger.Error, "[DeviceActivity | AfterFind] Error while unmarshaling software. ", err)
		return err
	}
	return nil
}

//State ...
func (activity *DeviceActivity) State(_config *configuration.ReportConfiguration) []sensors.ISensor {
	convertor := convert.NewDTOToState(activity.LastMessage, _config)
	return convertor.Convert()
}
