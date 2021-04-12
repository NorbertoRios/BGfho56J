package models

import (
	"encoding/json"
	"geometris-go/configuration"
	"geometris-go/convert"
	"geometris-go/core/interfaces"
	"geometris-go/core/sensors"
	"geometris-go/logger"
	"geometris-go/types"
	"time"

	"gorm.io/gorm"
)

//NewDeviceActivity ...
func NewDeviceActivity(dirtyState interfaces.IDirtyState, _sourseId uint64) *DeviceActivity {
	return &DeviceActivity{
		Identity:       dirtyState.Identity(),
		LastUpdateTime: time.Now().UTC(),
		LastMessageID:  _sourseId,
		syncParam:      dirtyState.SyncParam(),
	}
}

//DeviceActivity device activity model
type DeviceActivity struct {
	Identity           string            `gorm:"column:daiDeviceIdentity"`
	MessageTime        time.Time         `gorm:"column:daiLastMessageTime"`
	LastUpdateTime     time.Time         `gorm:"column:daiLastUpdateTime"`
	LastMessageID      uint64            `gorm:"column:daiLastMessageId"`
	LastMessage        string            `gorm:"column:daiLastMessage"`
	SerializedSoftware string            `gorm:"column:daiSoftware"`
	deviceFirmware     string            `gorm:"-" sql:"-"`
	software           *Software         `gorm:"-" sql:"-"`
	syncParam          string            `gorm:"-" sql:"-"`
	state              []sensors.ISensor `gorm:"-" sql:"-"`
}

//TableName for DeviceActivity model
func (DeviceActivity) TableName() string {
	return "ats.tblDeviceActivityInfo"
}

//AfterFind unmarshal string to struct
func (activity *DeviceActivity) BeforeUpdate(tx *gorm.DB) error {
	activity.lastMessage()
	activity.timeStamp()
	activity.firmware()
	activity.software = NewSoftware(activity.syncParam, activity.deviceFirmware)
	activity.SerializedSoftware = activity.software.Marshal()
	return nil
}

func (activity *DeviceActivity) lastMessage() {
	dtoMessage := convert.NewStateToDTO(activity.state).Convert()
	jMess, jErr := json.Marshal(dtoMessage)
	if jErr != nil {
		logger.Logger().WriteToLog(logger.Error, "[NewDeviceActivity] Error while marshaling dto. ", jErr)
		jMess = []byte{}
	}
	activity.LastMessage = string(jMess)
}

func (activity *DeviceActivity) firmware() {
	for _, sensor := range activity.state {
		if sensor.Symbol() == "Firmware" {
			activity.deviceFirmware = sensor.Value().(string)
		}
	}
	logger.Logger().WriteToLog(logger.Error, "[DeviceActivity | timeStamp] Cant find time stamp is sensors.")
	activity.deviceFirmware = ""
}

func (activity *DeviceActivity) timeStamp() {
	for _, sensor := range activity.state {
		if sensor.Symbol() == "TimeStamp" {
			activity.MessageTime = sensor.Value().(*types.JSONTime).Time
		}
	}
	logger.Logger().WriteToLog(logger.Error, "[DeviceActivity | timeStamp] Cant find time stamp is sensors.")
	activity.MessageTime = time.Time{}
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

//Sate ...
func (d *DeviceActivity) State() []sensors.ISensor {
	convertor := convert.NewDTOToState(d.LastMessage, configuration.ReportConfig(types.NewFile("/config/initialize/ReportConfiguration.xml")))
	return convertor.Convert()
}
