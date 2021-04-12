package models

import (
	"geometris-go/configuration"
	"geometris-go/convert"
	"geometris-go/core/sensors"
	"geometris-go/types"
	"time"
)

//NewDeviceActivity ...
func NewDeviceActivity() *DeviceActivity {
	return nil
}

//DeviceActivity device activity model
type DeviceActivity struct {
	Identity           string    `gorm:"column:daiDeviceIdentity"`
	MessageTime        time.Time `gorm:"column:daiLastMessageTime"`
	LastUpdateTime     time.Time `gorm:"column:daiLastUpdateTime"`
	LastMessageID      uint64    `gorm:"column:daiLastMessageId"`
	LastMessage        string    `gorm:"column:daiLastMessage"`
	Serializedsoftware string    `gorm:"column:daiSoftware"`
}

//TableName for DeviceActivity model
func (DeviceActivity) TableName() string {
	return "ats.tblDeviceActivityInfo"
}

//Sate ...
func (d *DeviceActivity) Sate() []sensors.ISensor {
	convertor := convert.NewDTOToState(d.LastMessage, configuration.ReportConfig(types.NewFile("/config/initialize/ReportConfiguration.xml")))
	return convertor.Convert()
}
